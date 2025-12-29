package cpu

import (
	"testing"
)

func TestBeq(t *testing.T) {
	cpu := InitCPU()
	cpu.PC = 0x04
	cpu.Registers[16] = 10 // setting $s0
	cpu.Registers[8] = 10 // setting $t0

	// beq $s0, $t0, done
	cpu.Instruction = 0x1208000e
	cpu.DecodeInstr()
	var targetPC uint32 = (0xe << 2) + 0x04

	if cpu.PC != targetPC {
		t.Errorf("PC wrong. expected=%d, got=%d", targetPC, cpu.PC)
	}
}

func TestAddi(t *testing.T) {
	cpu := InitCPU()
	cpu.Registers[8] = 5; // writing to $t0

	// addi $s0, $t0, 4095
	cpu.Instruction = 0x21100FFF
	cpu.DecodeInstr()

	if cpu.Registers[16] != 4095+5 {
		t.Errorf("destination register wrong. expected=%d, got=%d", 4095+5, cpu.Registers[16])
	}
}

func TestLW(t *testing.T) {
	cpu := InitCPU()
	cpu.Registers[16] = 0xFF // writing to $s0

	// writing to memory at the specified address
	fullAddr := uint32(4 + cpu.Registers[16])
	err := cpu.MainMemory.StoreWord(fullAddr, 1024)
	if err != nil {
		t.Fatalf("memory write failed")
	}
	// lw $t0, 4($s0)
	cpu.Instruction = 0x8e080004
	cpu.DecodeInstr()

	if cpu.Registers[8] != 1024 {
		t.Errorf("destination register wrong. expected=%d, got=%d", 1024, cpu.Registers[8])
	}
}

func TestSW(t *testing.T) {
	cpu := InitCPU()
	cpu.Registers[16] = 0xFF // writing to $s0
	cpu.Registers[8] = 1024

	// sw $t0, 4($s0)
	cpu.Instruction = 0xae080004
	cpu.DecodeInstr()

	actual, err := cpu.MainMemory.LoadWord(uint32(4 + cpu.Registers[16]))
	if err != nil {
		t.Fatalf("memory read failed")
	}

	if actual != 1024 {
		t.Errorf("value at memory address wrong. expected=%d, got=%d", 1024, actual)
	}
}
	
