from registers import MemoryMappedRegisters as keyboard_reg
from registers import Registers

def check_key() -> bool:
  # ???
  return True

def mem_write(address, value):
  memory[address] = value

def mem_read(address) -> int:
  if address == keyboard_reg.MR_KBSR and check_key():
    memory[keyboard_reg.MR_KBSR] = (1 << 15)
    memory[keyboard_reg.MR_KBDR] = getchar()
  else:
    memory[keyboard_reg.MR_KBSR] = 0

  return memory[address]
 
if __name__ == '__main__':
  PC_START = 0x3000

  running = True

  while running:
    instr = mem_read(reg[Registers.R_PC] + 1) 
    op = instr >> 12


