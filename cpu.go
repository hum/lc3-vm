package vm

// ALL TYPES
type (
	REG_VAL  uint16
	OP_VAL   uint16
	COND_VAL uint16
)

// REGISTERS
const (
	R_R0 REG_VAL = iota
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
	OP_BR   OP_VAL = iota // branch
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
	FL_POS COND_VAL = 1 << 0 // positive
	FL_ZRO COND_VAL = 1 << 1 // zero
	FL_NEG COND_VAL = 1 << 2 // negative
)

// REGISTER STORAGE
var reg [R_COUNT]REG_VAL
