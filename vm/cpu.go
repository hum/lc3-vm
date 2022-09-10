package vm

import (
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

instr_loop:
	for {
		var instruction uint16
		if instruction, err := cpu.RAM.MemRead(cpu.reg[R_PC]); err != nil {
			log.Printf("WARNING: MEM_READ out of range for value %x", instruction)
			break
		}
		cpu.reg[R_PC] = cpu.reg[R_PC] + 1

		var op uint16 = instruction >> 12
		log.Printf("Instruction: %x, OP CODE: %x", instruction, op)
		switch op {
		case OP_BR:
			//cpu.branch(op)
		case OP_ADD:
			cpu.add(op)
		case OP_LD:
			cpu.load_indirect(op)
		case OP_ST:
			cpu.store(op)
		case OP_JSR:
			//cpu.jump(op)
		case OP_AND:
			//cpu.bitwiseAnd(op)
		case OP_LDR:
			//cpu.loadRegister(op)
		case OP_STR:
			//cpu.storeRegister(op)
		case OP_NOT:
			//cpu.bitwiseNot(op)
		case OP_LDI:
			//cpu.loadIndirect(op)
		case OP_STI:
			//cpu.storeIndirect(op)
		case OP_JMP:
			//cpu.jump(op)
		case OP_LEA:
			//cpu.loadEffectiveAddr(op)
		case OP_TRAP:
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
	var sign uint16 = FL_POS // DEFAULT POSITIVE
	if cpu.reg[r] == 0 {     // FLAG ZERO
		sign = FL_ZRO
	} else if (cpu.reg[r] >> 15) != 0 { // FLAG NEGATIVE
		sign = FL_NEG
	}
	cpu.reg[R_COND] = sign
}

func (cpu *CPU) add(instr uint16) {
	/* in register mode, the second value to add is in a register
	* in immediate mode, the second value is embedded
	* in the right-most 5 bits of the instruction */
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

func (cpu *CPU) load_indirect(instr uint16) {
	// loads value from mem into the cpu register
	var r0 uint16 = (instr >> 9) & 0x7
	var offset uint16 = cpu.signExtend((instr & 0x1FF), 9)

	// add offset to PC, peek memory to get the final pointer address
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

	// stores value into mem from the first register
	cpu.RAM.MemWrite(cpu.reg[R_PC]+pcOffset, cpu.reg[r0])
}
