package vm

import (
  "fmt"
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

// REGISTER STORAGE
var reg [R_COUNT]uint16

type CPU struct {
	StartPosition uint16
}

func GetCPU() *CPU {
	return &CPU{
		StartPosition: PC_START,
	}
}

func (cpu *CPU) Run() {
  reg[R_PC] = cpu.StartPosition

  running := 0
  for running < 1 {
    var instruction uint16 = reg[R_PC]
    var op uint16 = cpu.memRead(instruction) >> 12
    reg[R_PC] = reg[R_PC] + 1
    fmt.Println(instruction, op)
  }
}

// TODO: implement, and move to RAM struct
func (cpu *CPU) memRead(addr uint16) (value uint16) {
  value = addr
  return value
}

// TODO: implement, and move to RAM struct
func (cpu *CPU) memWrite(addr uint16, value uint16) bool {
  return false
}
