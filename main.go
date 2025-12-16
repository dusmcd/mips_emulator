package main

import (
	"fmt"
	"mips_emulator/cpu"
	"mips_emulator/memory"
)

func main() {

	fmt.Printf("This is a MIPS Emulator\n")
	memory := memory.InitMemory()
	fmt.Println("Memory initialized")
	memory.LoadWord(0)
	cpu := cpu.InitCPU()
	cpu.Instruction = 0x00000020
	cpu.DecodeInstr()

	for _, val := range *cpu.Registers {
		fmt.Printf("%d ", val)
	}
	fmt.Printf("\n")
}
