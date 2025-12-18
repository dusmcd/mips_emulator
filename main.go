package main

import (
	"fmt"
	"mips_emulator/cpu"
)

func main() {

	fmt.Printf("This is a MIPS Emulator\n")
	cpu := cpu.InitCPU()
	initialAddr := uint32(0x01)
	numInstructions := 2
	cpu.MainMemory.LoadInstruction(initialAddr, 0x01285020)
	cpu.MainMemory.LoadInstruction(initialAddr + 4, 0x01285022)
	cpu.Run(numInstructions, initialAddr)

	fmt.Println("Registers")
	fmt.Println("=========")
	for i, val := range *cpu.Registers {
		fmt.Printf("%d:\t%d\n", i, val)	
	}

}
