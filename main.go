package main

import (
	"log"
	"os"

	lc3 "github.com/hum/lc3-vm/vm"
)

func main() {
	// Check input
	args := os.Args[1:]
	if len(args) == 0 {
		log.Println("No file input specified. Quitting.")
		os.Exit(0)
	}

	// Initialize memory buffer
	// TODO:
	// Refactor the ReadObj function to only pass back a slice of bytes
	// And then fill in the memory buffer as part of the inicialization process
	// of the VM
	var memBuffer [MEM_SIZE]uint16 = [MEM_SIZE]uint16{}

	var filename string = args[0]
	log.Printf("loading \"%s\" image file", filename)

	// Load image into the buffer
	memBuffer, err := ReadObj(filename)
	if err != nil {
		panic(err)
	}

	// Init
	var ram *lc3.RAM = lc3.GetRAM(memBuffer)
	var cpu *lc3.CPU = lc3.GetCPU(ram)

	// Execute
	cpu.Run()
}
