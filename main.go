package main

import (
	"fmt"
	"vm/cpu"
)

func main() {

	fmt.Printf("This is my Virtual Machine\n")
	cpu := cpu.InitCPU()
	cpu.Instruction = 0x00000020
	cpu.DecodeInstr()

	for _, val := range *cpu.Registers {
		fmt.Printf("%d ", val)
	}
	fmt.Printf("\n")
}
