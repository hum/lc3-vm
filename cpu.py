from registers import MemoryMappedRegisters as _keyboard_reg
from registers import Registers as _registers
from registers import Instructions as _instructions

def check_key() -> bool:
  # ???
  return True

def mem_write(address, value):
  memory[address] = value

def mem_read(address) -> int:
  if address == _keyboard_reg.MR_KBSR and check_key():
    memory[_keyboard_reg.MR_KBSR] = (1 << 15)
    memory[_keyboard_reg.MR_KBDR] = getchar()
  else:
    memory[_keyboard_reg.MR_KBSR] = 0

  return memory[address]
 
if __name__ == '__main__':
  PC_START = 0x3000

  running = True

  while running:
    instr = mem_read(reg[_registers.R_PC] + 1) 
    op = instr >> 12

		if op == _instructions.OP_ADD:
			break
		elif op == _instructions.OP_AND:
			break
		elif op == _instructions.OP_NOT:
			break
		elif op == _instructions.OP_BR:
			break
		elif op == _instructions.OP_JMP:
			break
		elif op == _instructions.OP_JSR:
			break
		elif op == _instructions.OP_LD:
			break
		elif op == _instructions.OP_LDI:
			break
		elif op == _instructions.OP_LDR
			break
		elif op == _instructions.OP_LEA:
			break
		elif op == _instructions.OP_ST:
			break
		elif op == _instructions.OP_STI:
			break
		elif op == _instructions.OP_STR:
			break
		elif op == _instructions.OP_TRAP:
			break
		elif op == _instructions.OP_RES:
			break
		elif op == _instructions.OP_RTI:
			break
		else:
			break
