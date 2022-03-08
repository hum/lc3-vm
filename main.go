package main

import (
	"fmt"
	lc3 "github.com/hum/lc3-vm/vm"
	"os"
)

func main() {
	args := os.Args[1:]
	fmt.Println(args)

	for arg := range args {
		fmt.Println("File: %s", arg)
		// TODO: check if exists and read the assembly from the file
	}

	var ram *lc3.RAM = lc3.GetRAM()
	var cpu *lc3.CPU = lc3.GetCPU(ram)
	cpu.Run()
}
