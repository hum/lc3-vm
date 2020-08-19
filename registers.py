from enum import Enum
from enum import auto

class Registers(Enum):
  R_R0 = 0
  R_R1 = auto()
  R_R2 = auto()
  R_R4 = auto()
  R_R5 = auto()
  R_R6 = auto()
  R_R7 = auto()
  R_PC = auto() # program counter
  R_COND = auto()
  R_COUNT = auto()

class Instructions(Enum):
	OP_BR = 0,
	OP_ADD = auto()
	OP_LD = auto()
	OP_ST = auto()
	OP_JSR = auto()
	OP_AND = auto()
	OP_LDR = auto()
	OP_STR = auto()
	OP_RTI = auto()
	OP_NOT = auto()
	OP_LDI = auto()
	OP_STI = auto()
	OP_JMP = auto()
	OP_RES = auto()
	OP_LEA = auto()
	OP_TRAP	= auto()

class Conditions(Enum):
  FL_POS = 1 << 0
  FL_ZRO = 1 << 1
  FL_NEG = 1 << 2

class MemoryMappedRegisters(Enum):
  MR_KBSR = 0xFE00
  MR_KBDR = 0xFE02
