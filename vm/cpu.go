package vm

import (
	"fmt"
)

const (
	R_R0 uint16 = iota
	R_R1
	R_R2
	R_R3
	R_R4
	R_R5
	R_R6
	R_R7
	R_PC         // Program counter
	R_COND       // Condition flag
	R_COUNT = 10 // Count of the registers
)

type CPU struct {
	// Initial address the program loads into
	ProgramStart uint16
	// Array of the available cpu.Registeristers
	Register [R_COUNT]uint16
}

func (cpu *CPU) dump() {
	fmt.Println("\ndumping CPU registers")
	for i, v := range cpu.Register {
		fmt.Printf("R_R%d, value: 0x%04x(%d)\n", i, v, v)
	}
}

func (cpu *CPU) Execute() error {
	cpu.Register[R_COND] = FL_ZRO

instr_loop:
	for {
		// Fetch the instruction
		var instr uint16 = ram.Read(cpu.Register[R_PC])
		// Get the instruction's OP code
		var op uint16 = instr >> 12

		// Increment the counter
		cpu.Register[R_PC]++

		switch op {
		case OP_BR:
			cpu.branch(instr)
		case OP_ADD:
			cpu.add(instr)
		case OP_LD:
			cpu.load(instr)
		case OP_JSR:
			cpu.jumpRegister(instr)
		case OP_AND:
			cpu.bitwiseAnd(instr)
		case OP_NOT:
			cpu.bitwiseNot(instr)
		case OP_LDR:
			cpu.loadRegister(instr)
		case OP_LEA:
			cpu.loadEffectiveAddress(instr)
		case OP_LDI:
			cpu.loadIndirect(instr)
		case OP_ST:
			cpu.store(instr)
		case OP_STR:
			cpu.storeRegister(instr)
		case OP_STI:
			cpu.storeIndirect(instr)
		case OP_JMP:
			cpu.jump(instr)
		case OP_RES:
		case OP_RTI:
		case OP_TRAP:
			cpu.Register[R_R7] = cpu.Register[R_PC]

			switch instr & 0xFF {
			case TRAP_GETC:
				cpu.trapGetC()
			case TRAP_OUT:
				cpu.trapOut()
			case TRAP_PUTS:
				cpu.trapPuts()
			case TRAP_IN:
				cpu.trapIn()
			case TRAP_PUTSP:
				cpu.trapPutsP()
			case TRAP_HALT:
				Dump()
				break instr_loop
			}
		default:
			Dump()
			return fmt.Errorf("invalid OP code %x", op)
		}
	}
	return nil
}

func (cpu *CPU) signExtend(x uint16, bitCount int) uint16 {
	// Extends bits to the size 16
	// Fills positive with 0
	// Fills negative with 1
	if (x >> (bitCount - 1) & 1) == 1 {
		x |= 0xFFFF << bitCount
	}
	return x
}

func (cpu *CPU) updateFlags(r uint16) {
	// Default sign is positive
	var sign uint16 = FL_POS
	if cpu.Register[r] == 0 {
		// If value in the register is zero
		sign = FL_ZRO
	} else if (cpu.Register[r] >> 15) != 0 {
		// One (1) in the left-most bit indicates negative number
		sign = FL_NEG
	}
	cpu.Register[R_COND] = sign
}

func (cpu *CPU) add(instr uint16) {
	// In register mode, the second value to add is in a register
	// In immediate mode, the second value is embedded
	// In the right-most 5 bits of the instruction
	var r0 uint16 = (instr >> 9) & 0x7             // destination register
	var r1 uint16 = (instr >> 6) & 0x7             // first operand
	var immediate_mode uint16 = (instr >> 5) & 0x1 // intermediate mode? -- 5th bit 0 or 1?

	if immediate_mode != 0 {
		var imm5 uint16 = cpu.signExtend((instr & 0x1F), 5)
		cpu.Register[r0] = cpu.Register[r1] + imm5
	} else {
		var r2 uint16 = instr & 0x7
		cpu.Register[r0] = cpu.Register[r1] + cpu.Register[r2]
	}
	cpu.updateFlags(r0)
}

func (cpu *CPU) loadIndirect(instr uint16) {
	// Loads value from mem into the cpu.Registerister
	var r0 uint16 = (instr >> 9) & 0x7
	var offset uint16 = cpu.signExtend((instr & 0x1FF), 9)

	// Add offset to PC, peek memory to get the final pointer address

	cpu.Register[r0] = ram.Read(ram.Read(cpu.Register[R_PC] + offset))
	cpu.updateFlags(r0)
}

func (cpu *CPU) store(instr uint16) {
	var r0 uint16 = (instr >> 9) & 0x7
	var pcOffset uint16 = cpu.signExtend(instr&0x1FF, 9)

	// Stores value into mem from the first register
	ram.Write(cpu.Register[R_PC]+pcOffset, cpu.Register[r0])
}

func (cpu *CPU) jump(instr uint16) {
	// This also handles RET (return from subroutine)
	// whenver R1 is 111 in binary (decimal 7)
	var r1 uint16 = (instr >> 6) & 0x7

	// Move program counter to the value of second register
	cpu.Register[R_PC] = cpu.Register[r1]
}

func (cpu *CPU) bitwiseAnd(instr uint16) {
	var r0 uint16 = (instr >> 9) & 0x7      // Destination register
	var r1 uint16 = (instr >> 6) & 0x7      // Source register for value
	var immFlag uint16 = (instr >> 5) & 0x1 // Immediate value

	if immFlag == 1 {
		// Sign extend to 16 bits
		var imm5 uint16 = cpu.signExtend(instr&0x1F, 5)
		// AND the values
		cpu.Register[r0] = cpu.Register[r1] & imm5
	} else {
		// Get value from SR2
		var r2 uint16 = instr & 0x7
		// AND the values
		cpu.Register[r0] = cpu.Register[r1] & cpu.Register[r2]
	}
	cpu.updateFlags(r0)
}

func (cpu *CPU) branch(instr uint16) {
	var pcOffset uint16 = cpu.signExtend(instr&0x1FF, 9)
	var condFlag uint16 = (instr >> 9) & 0x7

	if condFlag&cpu.Register[R_COND] != 0 {
		cpu.Register[R_PC] += pcOffset
	}
}

func (cpu *CPU) bitwiseNot(instr uint16) {
	var r0 uint16 = (instr >> 9) & 0x7 // Destination register
	var r1 uint16 = (instr >> 6) & 0x7

	// Bitwise complement
	cpu.Register[r0] = ^cpu.Register[r1] // Using XOR instead
	cpu.updateFlags(r0)
}

func (cpu *CPU) load(instr uint16) {
	var r0 uint16 = (instr >> 9) & 0x7
	var offset uint16 = cpu.signExtend(instr&0x1FF, 9)

	cpu.Register[r0] = ram.Read(cpu.Register[R_PC] + offset)
	cpu.updateFlags(r0)
}

func (cpu *CPU) loadRegister(instr uint16) {
	var r0 uint16 = (instr >> 9) & 0x7
	var r1 uint16 = (instr >> 6) & 0x7
	var memOffset uint16 = cpu.signExtend(instr&0x3F, 6)

	cpu.Register[r0] = ram.Read(cpu.Register[r1] + memOffset)
	cpu.updateFlags(r0)
}

func (cpu *CPU) loadEffectiveAddress(instr uint16) {
	var r0 uint16 = (instr >> 9) & 0x7
	var pcOffset uint16 = cpu.signExtend(instr&0x1FF, 9)

	cpu.Register[r0] = cpu.Register[R_PC] + pcOffset
	cpu.updateFlags(r0)
}

func (cpu *CPU) storeIndirect(instr uint16) {
	var r0 uint16 = (instr >> 9) & 0x7
	var pcOffset uint16 = cpu.signExtend(instr&0x1FF, 9)

	v := ram.Read(cpu.Register[R_PC] + pcOffset)
	ram.Write(v, cpu.Register[r0])
}

func (cpu *CPU) storeRegister(instr uint16) {
	var r0 uint16 = (instr >> 9) & 0x7
	var r1 uint16 = (instr >> 6) & 0x7
	var memOffset uint16 = cpu.signExtend(instr&0x3F, 6)

	ram.Write(cpu.Register[r1]+memOffset, cpu.Register[r0])
}

func (cpu *CPU) jumpRegister(instr uint16) {
	var longFlag uint16 = (instr >> 11) & 1
	cpu.Register[R_R7] = cpu.Register[R_PC]

	if longFlag == 1 {
		var longPCOffset uint16 = cpu.signExtend(instr&0x7FF, 11)
		cpu.Register[R_PC] += longPCOffset
	} else {
		var r1 uint16 = (instr >> 6) & 0x7
		cpu.Register[R_PC] = cpu.Register[r1]
	}
}

func (cpu *CPU) trapPuts() {
	var origin uint16 = cpu.Register[R_R0]
	var ch uint16 = ram.Read(origin)

	for ch != 0 {
		WriteChar(ch)

		origin++
		ch = ram.Read(origin)
	}
}

func (cpu *CPU) trapPutsP() {
	var origin uint16 = cpu.Register[R_R0]
	var ch uint16 = ram.Read(origin)

	for ch != 0 {
		var c1 uint16 = ch & 0xFF
		WriteChar(c1)

		var c2 uint16 = ch >> 8
		if c2 == 1 {
			WriteChar(c2)
		}
		origin++
		ch = ram.Read(origin)
	}
}

func (cpu *CPU) trapGetC() {
	// Syscall to get char from STDIN
	ch, err := GetChar()
	if err != nil {
		panic(err)
	}
	cpu.Register[R_R0] = uint16(ch)
	cpu.updateFlags(R_R0)
}

func (cpu *CPU) trapOut() {
	WriteChar(cpu.Register[R_R0])
}

func (cpu *CPU) trapIn() {
	fmt.Println("Enter a character: ")
	c, err := GetChar()
	if err != nil {
		panic(err)
	}
	WriteChar(c)
	cpu.Register[R_R0] = uint16(c)
	cpu.updateFlags(R_R0)
}
