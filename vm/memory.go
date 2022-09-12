package vm

import "fmt"

const MemSize uint32 = 65536

type RAM struct {
	mem [MemSize]uint16
}

func GetRAM(memoryBuffer [MemSize]uint16) *RAM {
	return &RAM{
		mem: memoryBuffer,
	}
}

func (r *RAM) MemRead(addr uint16) uint16 {
	return r.mem[addr]
}

func (r *RAM) MemWrite(addr uint16, value uint16) (ok bool) {
	r.mem[addr] = value
	return true
}

func (r *RAM) Dump() {
	for i, v := range r.mem {
		if v == 0 {
			continue
		}
		fmt.Printf("index: %d, value: %d", i, v)
	}
}
