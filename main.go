package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hum/lc3-vm/vm"
)

func main() {
	// TODO:
	// Add all of the logs into a debug mode instead
	log.Println("starting LC-3 VM")

	// Check input
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("No file input specified. Quitting.")
		os.Exit(0)
	}

	var filename string = args[0]
	log.Printf("loading \"%s\" file into the memory", filename)

	// Load image into the buffer
	data, n, err := ReadObjFile(filename)
	if err != nil || n <= 0 {
		if err != nil {
			panic(err)
		}
		// TODO:
		// Better handling of the input in general
		panic("read file has 0 size")
	}

	// Load the byte slice into the VM
	vm.LoadByteSliceBuffer(data, n)
	log.Printf("loaded %d bytes buffer into the memory", n)

	if err := vm.Execute(); err != nil {
		// Gracefully shutdown the vm
		fmt.Errorf("vm error: %s", err)
		os.Exit(0)
	}
}
