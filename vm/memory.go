package vm

import "log"

type RAM struct {
	// Hard limit for the amount of memory
	memoryLimit uint16
	// Array for the actual data being stored in the memory
	buffer [MEM_LIMIT]uint16
}

func (r *RAM) dump() {
	log.Println("dumping RAM buffer")
	for i, v := range r.buffer {
		if v == 0 {
			continue
		}
		log.Printf("index: 0x%04x(%d), value: 0x%04x(%d)\n", i, i, v, v)
	}
}

func (r *RAM) Write(addr, value uint16) uint16 {
	if addr >= r.memoryLimit {
		panic("program wrote outside of its memory stack")
	}
	r.buffer[addr] = value
	return value
}

func (r *RAM) Read(addr uint16) uint16 {
	if addr >= r.memoryLimit {
		panic("program wrote outside of its memory stack")
	}
	return r.buffer[addr]
}
