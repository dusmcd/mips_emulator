package main

import (
	"fmt"
	"mips_emulator/cpu"
	"mips_emulator/defs"
	"mips_emulator/loader"
	"encoding/binary"
	"os"
	"log"
)

func ReadInstructions(filePath string, initialAddr uint32, cpu *cpu.CPU) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	instructions := []uint32{}
	fmt.Printf("size of data %dB\n", len(data))
	for i := 0; i < len(data); i += 4{
		var instr uint32 = binary.LittleEndian.Uint32(data[i:i+4])
		instructions = append(instructions, instr)
	}

	for i, instr := range instructions {
		cpu.MainMemory.LoadInstruction(initialAddr + uint32(i * 4), defs.Word(instr))
	}

	return nil
}

func main() {

	fmt.Printf("This is a MIPS Emulator\n")
	err := loader.ParseFile("c_files/main")
	if err != nil {
		log.Fatalf("%s", err)
	}
	
	if (len(os.Args) < 2) {
		log.Fatalf("example usage: mips_em <binary file>")
	}
	cpu := cpu.InitCPU()
	initialAddr := uint32(0x04)
	err = ReadInstructions(os.Args[1], initialAddr, cpu)
	if err != nil {
		log.Fatalf("error reading binary file: %s", err.Error())
	}

	cpu.Run(initialAddr)

	fmt.Println("Registers")
	fmt.Println("=========")
	for i, val := range cpu.Registers {
		fmt.Printf("%d:\t%d\n", i, val)	
	}
	

}
