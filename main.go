package main

import (
	"log"
	"os"

	lc3 "github.com/hum/lc3-vm/vm"
)

func main() {
	args := os.Args[1:]

	// Initialize memory buffer
	var memBuffer [MEM_SIZE]uint16 = [MEM_SIZE]uint16{}

	// Load image file
	if len(args) == 0 {
		log.Println("No file input specified. Quitting.")
		os.Exit(0)
	}

	var filename string = args[0]
	log.Printf("loading \"%s\" image file", filename)

	// Load image into the buffer
	memBuffer, err := readFile(filename)
	if err != nil {
		panic(err)
	}

	// Init
	var ram *lc3.RAM = lc3.GetRAM(memBuffer)
	var cpu *lc3.CPU = lc3.GetCPU(ram)

	// Execute
	cpu.Run()
}
