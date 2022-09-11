package vm

import (
	"fmt"
	"log"
)

const (
	PC_START uint16 = 0x3000 // program counter starting register
)

type CPU struct {
	RAM           *RAM
	reg           [R_COUNT]uint16
	startPosition uint16
}

func GetCPU(ram *RAM) *CPU {
	return &CPU{
		startPosition: PC_START,
		RAM:           ram,
	}
}

func (cpu *CPU) Run() {
	cpu.reg[R_PC] = cpu.startPosition
	cpu

instr_loop:
	for {
		var instr uint16
		if instr, err := cpu.RAM.MemRead(cpu.reg[R_PC]); err != nil {
			log.Printf("WARNING: MEM_READ out of range for value %x", instr)
			break
		}
		cpu.reg[R_PC] = cpu.reg[R_PC] + 1
		var op uint16 = instr >> 12
		log.Printf("instr: %x, OP CODE: %x", instr, op)
		switch op {
		case OP_BR:
			cpu.branch(op)
		case OP_ADD:
			cpu.add(op)
		case OP_LD:
			cpu.loadIndirect(op)
		case OP_JSR:
			cpu.jump(op)
		case OP_AND:
			cpu.bitwiseAnd(op)
		case OP_NOT:
			cpu.bitwiseNot(op)
		case OP_LDR:
			cpu.loadRegister(op)
		case OP_LEA:
			cpu.loadEffectiveAddress(op)
		case OP_LDI:
			cpu.loadIndirect(op)
		case OP_ST:
			cpu.store(op)
		case OP_STR:
			cpu.storeRegister(op)
		case OP_STI:
			cpu.storeIndirect(op)
		case OP_JMP:
			cpu.jump(op)
		case OP_TRAP:
			switch instr & 0xFF {
			case TRAP_GETC:
				cpu.trapGetC(op)
				break
			case TRAP_OUT:
				cpu.trapOut(op)
			case TRAP_PUTS:
				cpu.trapPuts(op)
			case TRAP_IN:
				break
			case TRAP_PUTSP:
				break
			case TRAP_HALT:
				break
			}
		case OP_RES:
		case OP_RTI:
		default:
			log.Printf("ERROR: Invalid OP code %x. Quitting.", op)
			break instr_loop
		}
	}
}

func (cpu *CPU) signExtend(x uint16, bitCount int) uint16 {
	/* extends bits to size 16
	fill positive with 0
	fill negative with 1
	*/
	if (x >> (bitCount - 1) & 1) > 0 {
		x |= (0xFFFF << bitCount)
	}
	return x
}

func (cpu *CPU) updateFlags(r uint16) {
	// Default sign is positive
	var sign uint16 = FL_POS
	if cpu.reg[r] == 0 {
		// If value in the register is zero
		sign = FL_ZRO
	} else if (cpu.reg[r] >> 15) == 1 {
		// One (1) in the left-most bit indicates negative number
		sign = FL_NEG
	}
	cpu.reg[R_COND] = sign
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
		cpu.reg[r0] = cpu.reg[r1] + imm5
	} else {
		var r2 uint16 = instr & 0x7
		cpu.reg[r0] = cpu.reg[r1] + cpu.reg[r2]
	}
	cpu.updateFlags(r0)
}

func (cpu *CPU) loadIndirect(instr uint16) {
	// Loads value from mem into the cpu register
	var r0 uint16 = (instr >> 9) & 0x7
	var offset uint16 = cpu.signExtend((instr & 0x1FF), 9)

	// Add offset to PC, peek memory to get the final pointer address
	finalMem, err := cpu.RAM.MemRead(cpu.reg[R_PC] + offset)
	if err != nil {
		panic(err)
	}
	v, err := cpu.RAM.MemRead(finalMem)
	if err != nil {
		panic(err)
	}
	cpu.reg[r0] = v
}

func (cpu *CPU) store(instr uint16) {
	var r0 uint16 = (instr >> 9) & 0x7
	var pcOffset uint16 = cpu.signExtend(instr&0x1FF, 9)

	// Stores value into mem from the first register
	cpu.RAM.MemWrite(cpu.reg[R_PC]+pcOffset, cpu.reg[r0])
}

func (cpu *CPU) jump(instr uint16) {
	// This also handles RET (return from subroutine)
	// whenver R1 is 111 in binary (decimal 7)
	var r1 uint16 = (instr >> 6) & 0x7

	// Move program counter to the value of second register
	cpu.reg[R_PC] = cpu.reg[r1]
}

func (cpu *CPU) bitwiseAnd(instr uint16) {
	var r0 uint16 = (instr >> 9) & 0x7      // Destination register
	var r1 uint16 = (instr >> 6) & 0x7      // Source register for value
	var immFlag uint16 = (instr >> 5) & 0x1 // Immediate value

	if immFlag == 1 {
		// Sign extend to 16 bits
		var imm5 uint16 = cpu.signExtend(instr&0x1F, 5)
		// AND the values
		cpu.reg[r0] = cpu.reg[r1] & imm5
	} else {
		// Get value from SR2
		var r2 uint16 = instr & 0x7
		// AND the values
		cpu.reg[r0] = cpu.reg[r1] & cpu.reg[r2]
	}
	cpu.updateFlags(r0)
}

func (cpu *CPU) branch(instr uint16) {
	var pcOffset uint16 = cpu.signExtend(instr&0x1FF, 9)
	var condFlag uint16 = (instr >> 9) & 0x7

	if condFlag&cpu.reg[R_COND] == 1 {
		cpu.reg[R_PC] += pcOffset
	}
}

func (cpu *CPU) bitwiseNot(instr uint16) {
	var r0 uint16 = (instr >> 9) & 0x7 // Destination register
	var r1 uint16 = (instr >> 6) & 0x7

	// Bitwise complement
	cpu.reg[r0] = ^cpu.reg[r1] // Using XOR instead
	cpu.updateFlags(r0)
}

func (cpu *CPU) loadRegister(instr uint16) {
	var r0 uint16 = (instr >> 9) & 0x7
	var r1 uint16 = (instr >> 6) & 0x7
	var memOffset uint16 = cpu.signExtend(instr&0x3F, 6)

	v, err := cpu.RAM.MemRead(cpu.reg[r1] + memOffset)
	if err != nil {
		panic(err)
	}
	cpu.reg[r0] = v
	cpu.updateFlags(r0)
}

func (cpu *CPU) loadEffectiveAddress(instr uint16) {
	var r0 uint16 = (instr >> 9) & 0x7
	var pcOffset uint16 = cpu.signExtend(instr&0x1FF, 9)

	cpu.reg[r0] = cpu.reg[R_PC] + pcOffset
	cpu.updateFlags(r0)
}

func (cpu *CPU) storeIndirect(instr uint16) {
	var r0 uint16 = (instr >> 9) & 0x7
	var pcOffset uint16 = cpu.signExtend(instr&0x1FF, 9)
	v, err := cpu.RAM.MemRead(cpu.reg[R_PC] + pcOffset)
	if err != nil {
		panic(err)
	}
	cpu.RAM.MemWrite(v, cpu.reg[r0])
}

func (cpu *CPU) storeRegister(instr uint16) {
	var r0 uint16 = (instr >> 9) & 0x7
	var r1 uint16 = (instr >> 6) & 0x7
	var memOffset uint16 = cpu.signExtend(instr&0x3F, 6)

	cpu.RAM.MemWrite(cpu.reg[r1]+memOffset, cpu.reg[r0])
}

func (cpu *CPU) trapOut(instr uint16) {
	// Output a single character from the first register
	fmt.Printf("%c\n", rune(cpu.reg[R_R0]))
}

func (cpu *CPU) trapPuts(instr uint16) {
	var addr uint16 = cpu.reg[R_R0] // beginning of the memory

	for {
		// Read from memory
		ch, err := cpu.RAM.MemRead(addr)
		if err != nil {
			panic(err)
		}

		// Stop print if exit bits
		if ch == 0x00 {
			break
		}

		// Convert to rune to display the character in terminal
		fmt.Printf("%c", rune(ch))
	}
	fmt.Printf("\n")
}

func (cpu *CPU) trapGetC(instr uint16) {
	// TODO:
	// Implement similar func to getchar in C
	return
}
