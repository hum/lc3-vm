package vm

import (
	"fmt"
	"log"
)

const (
	MR_KBSR = 0xFE00 // Keyboard status
	MR_KBDR = 0xFE02 // Keyboard data
)

type RAM struct {
	// Array for the actual data being stored in the memory
	buffer [MEM_LIMIT]uint16
}

func (r *RAM) dump() {
	fmt.Println("\ndumping RAM buffer")
	for i, v := range r.buffer {
		if v == 0 {
			continue
		}
		fmt.Printf("index: 0x%04x(%d), value: 0x%04x(%d)\n", i, i, v, v)
	}
}

func (r *RAM) Write(addr uint16, value uint16) uint16 {
	if int(addr) >= len(r.buffer) {
		panic("program wrote outside of its memory stack")
	}
	r.buffer[addr] = value
	return value
}

func (r *RAM) Read(addr uint16) uint16 {
	if int(addr) >= len(r.buffer) {
		panic("program wrote outside of its memory stack")
	}

	if addr == MR_KBSR {
		log.Println("waiting for keyboard input")
		if v := readRuneFromInput(); v != 0 {
			r.Write(MR_KBSR, 1<<15)
			r.Write(MR_KBDR, v)
		} else {
			r.Write(MR_KBSR, 0)
		}
	}
	return r.buffer[addr]
}
