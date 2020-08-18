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
