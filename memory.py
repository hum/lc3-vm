import array

class Memory(object):
  def __init__(self):
	  self.memory_size = 2**16
	  self.memory = array.array("H", range(self.memory_size))
