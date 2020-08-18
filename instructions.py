from enum import Enum
from enum import auto

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
