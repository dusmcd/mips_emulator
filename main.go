package main

import (
	"fmt"
	"mips_emulator/cpu"
	"mips_emulator/defs"
	"os"
	"log"
)

func ReadInstructions(filePath string, initialAddr uint32, cpu *cpu.CPU) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	instructions := []uint32{}
	for i := 0; i < len(data); i += 4{
		var instr uint32 = 0
		instr = instr | uint32(data[i]) // least significant byte
		instr = instr | uint32(data[i + 1]) << 8
		instr = instr | uint32(data[i + 2]) << 16
		instr = instr | uint32(data[i + 3]) << 24 // most siginificant byte
		instructions = append(instructions, instr)
	}

	for i, instr := range instructions {
		cpu.MainMemory.LoadInstruction(initialAddr + uint32(i * 4), defs.Word(instr))
	}

	return nil
}

func main() {

	fmt.Printf("This is a MIPS Emulator\n")
	if (len(os.Args) < 2) {
		log.Fatalf("example usage: mips_em <binary file>")
	}
	cpu := cpu.InitCPU()
	initialAddr := uint32(0x01)
	err := ReadInstructions(os.Args[1], initialAddr, &cpu)
	if err != nil {
		log.Fatalf("error reading binary file")
	}

	cpu.Run(initialAddr)

	fmt.Println("Registers")
	fmt.Println("=========")
	for i, val := range *cpu.Registers {
		fmt.Printf("%d:\t%d\n", i, val)	
	}

}
