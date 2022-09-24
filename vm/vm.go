package vm

import (
	"bytes"
	"encoding/binary"
	"log"
)

const (
	MEM_LIMIT  = 65535  // 16-bit
	PROG_START = 0x3000 // Memory mapped starting point
)

var (
	cpu *CPU
	ram *RAM
)

func Init() {
	cpu = &CPU{
		ProgramStart: PROG_START,
		Register:     [R_COUNT]uint16{R_PC: PROG_START},
	}
	ram = &RAM{memoryLimit: MEM_LIMIT}
}

func Dump() {
	cpu.dump()
	ram.dump()
}

func LoadByteSliceBuffer(b []byte, size int) {
	if cpu == nil || ram == nil {
		// Initialize the components if not yet done
		Init()
	}

	// Header which includes the starting memory address of the program
	var origin uint16 = binary.BigEndian.Uint16(b[:2])
	log.Printf("begin memory address 0x%04x(%d)", origin, origin)

	// Skip the first two bytes because that's the header
	// which is loaded separately above
	buf := bytes.NewBuffer(b[2:])

	// Similar as for the buffer
	// the limit is set to "size-2" because the header was already loaded
	for i := 0; i < size-2; i += 2 {
		var v uint16
		// Since we are reading into a uint16 value, this will read 2 bytes
		err := binary.Read(buf, binary.BigEndian, &v)
		if err != nil {
			break
		}

		//fmt.Printf("origin: 0x%04x(%d), value: 0x%04x(%d)\n", origin, origin, v, v)

		ram.Write(origin, v)
		origin++
	}
}

func Execute() error {
	return cpu.Execute()
}
