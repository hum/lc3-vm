package main

import (
	"fmt"
	lc3 "github.com/hum/lc3-vm/vm"
)

func main() {
	vm := lc3.GetCPU()
	fmt.Println(vm.StartPosition)
	vm.Run()
}
