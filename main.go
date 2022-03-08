package main

import (
	lc3 "github.com/hum/lc3-vm/vm"
)

func main() {
	var ram *lc3.RAM = lc3.GetRAM()
	var cpu *lc3.CPU = lc3.GetCPU(ram)
	cpu.Run()
}
