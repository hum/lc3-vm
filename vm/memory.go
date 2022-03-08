package vm

import (
	"fmt"
)

const MemSize uint16 = 65535

type RAM struct {
	mem []uint16
}

func GetRAM() *RAM {
	return &RAM{
		mem: make([]uint16, MemSize),
	}
}

func (r *RAM) MemRead(addr uint16) (uint16, error) {
	if addr >= MemSize {
		return 0, fmt.Errorf("Addr %x is not within range %x.", addr, MemSize)
	}
	return r.mem[addr], nil
}

func (r *RAM) MemWrite(addr uint16, value uint16) (ok bool) {
	if addr >= MemSize {
		return false
	}
	r.mem[addr] = value
	return true
}
