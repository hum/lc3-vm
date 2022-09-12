package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

const (
	HEADER_SIZE = 2     // Header size of the image file
	MEM_SIZE    = 65536 // Total memory limit
)

func readFile(fn string) ([MEM_SIZE]uint16, error) {
	// Final memory buffer
	var memoryBuffer [MEM_SIZE]uint16

	f, err := os.Open(fn)
	if err != nil {
		return memoryBuffer, err
	}
	defer f.Close()

	// TODO:
	// Check format and other metadata of the file
	stats, err := f.Stat()
	if err != nil {
		// File probably doesn't exist
		return memoryBuffer, err
	}

	// The header is 16 bits long
	hBytes := make([]byte, HEADER_SIZE)
	s, err := f.Read(hBytes)
	if s != HEADER_SIZE {
		return memoryBuffer, fmt.Errorf("reader returned more bytes than requested. wanted %d bytes, got %d", HEADER_SIZE, s)
	}

	// Starting address of the program
	var origin uint16

	// Read the address from the header buffer
	hBuffer := bytes.NewBuffer(hBytes)
	err = binary.Read(hBuffer, binary.BigEndian, &origin)
	if err != nil {
		return memoryBuffer, err
	}

	log.Printf("starting address: 0x%x", origin)

	// Read file into the memory slice
	size := stats.Size()
	log.Printf("creating file buffer size: %d bytes", size)

	fileByteSlice := make([]byte, size)
	_, err = f.Read(fileByteSlice)
	if err != nil {
		return memoryBuffer, err
	}

	// Buffered reader for the rest of the file
	b := bytes.NewBuffer(fileByteSlice)

	for i := origin; i < MEM_SIZE-1; i++ {
		var v uint16

		// LC-3 is a big-endian computer
		binary.Read(b, binary.BigEndian, &v)

		// Save each value into the memory buffer
		memoryBuffer[i] = v
	}
	return memoryBuffer, nil
}
