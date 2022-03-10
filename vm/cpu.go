package vm

import (
	"log"
)

// REGISTERS
const (
	R_R0 uint16 = iota
	R_R1
	R_R2
	R_R3
	R_R4
	R_R5
	R_R6
	R_R7
	R_PC    // program counter
	R_COND  // condition flag
	R_COUNT = 10
)

// OP CODES
const (
	OP_BR   uint16 = iota // branch
	OP_ADD                // add
	OP_LD                 // load
	OP_ST                 // store
	OP_JSR                // jump register
	OP_AND                // bitwise and
	OP_LDR                // load register
	OP_STR                // store register
	OP_RTI                // unused
	OP_NOT                // bitwise not
	OP_LDI                // load indirect
	OP_STI                // store indirect
	OP_JMP                // jump
	OP_RES                // reserved (unused)
	OP_LEA                // load effective address
	OP_TRAP               // execute trap
)

// CONDITION FLAGS
const (
	FL_POS uint16 = 1 << 0 // positive
	FL_ZRO uint16 = 1 << 1 // zero
	FL_NEG uint16 = 1 << 2 // negative
)

// PROGRAM COUNTER STARTING REGISTER
const (
	PC_START uint16 = 0x3000
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
			cpu.branch(op)
		case OP_ADD:
			cpu.add(op)
		case OP_LD:
			cpu.load(op)
		case OP_ST:
			cpu.store(op)
		case OP_JSR:
			cpu.jump(op)
		case OP_AND:
			cpu.bitwiseAnd(op)
		case OP_LDR:
			cpu.loadRegister(op)
		case OP_STR:
			cpu.storeRegister(op)
		case OP_NOT:
			cpu.bitwiseNot(op)
		case OP_LDI:
			cpu.loadIndirect(op)
		case OP_STI:
			cpu.storeIndirect(op)
		case OP_JMP:
			cpu.jump(op)
		case OP_LEA:
			cpu.loadEffectiveAddr(op)
		case OP_TRAP:
		case OP_RES:
		case OP_RTI:
		default:
			log.Printf("ERROR: Invalid OP code %x. Quitting.", op)
			break
		}
	}
}

func (cpu *CPU) signExtend(x uint16, bitCount int) uint16 {
	/* extends bits to size 16
	fill positive with 0
	fill negative with 1
	*/
	if (x >> (bitCount - 1)) & 1 {
		x |= (0xFFFF << bitCount)
	}
	return x
}

func (cpu *CPU) updateFlags(r uint16) {
	var sign uint16 = FL_POS // DEFAULT POSITIVE
	if cpu.reg[r] == 0 {     // FLAG ZERO
		sign = FL_ZERO
	} else if cpu.reg[r] >> 15 { // FLAG NEGATIVE
		sign = FL_NEG
	}
	reg[R_COND] = sign
}

func (cpu *CPU) add(op uint16) {
	var r0 uint16 = (op >> 9) & 0x7
	var r1 uint16 = (op >> 6) & 0x7
	var immediate_mode uint16 = (op >> 5) & 0x1

	if immediate_mode {
		var imm5 uint16 = cpu.signExtend(op&0x1F, 5)
		cpu.reg[r0] = cpu.reg[r1] + imm5
	} else {
		var r2 uint16 = op & 0x7
		cpu.reg[r0] = cpu.reg[r1] + cpu.reg[r2]
	}
	cpu.updateFlags(r0)
}
