package cpu

import (
	"testing"
)

func TestAdd(t *testing.T) {
	cpu := InitCPU()
	cpu.Registers[8] = 5 // set register $t0
	cpu.Registers[9] = 10 // set register $t1

	// add $t2, $t1, $t0
	cpu.Instruction = 0x01285020
	cpu.DecodeInstr()
	
	if cpu.Registers[10] != 15 {
		t.Errorf("destination register wrong. expected=%d, got=%d", 15, cpu.Registers[10])
	}
}

func TestSubtract(t *testing.T) {
	cpu := InitCPU()

	cpu.Registers[8] = 5 // set register $t0
	cpu.Registers[9] = 10 // set register $t1

	// sub $t2, $t1, $t0
	cpu.Instruction = 0x01285022
	cpu.DecodeInstr()
	
	if cpu.Registers[10] != 5 {
		t.Errorf("destination register wrong. expected=%d, got=%d", 5, cpu.Registers[10])
	}

}

func TestAddU(t *testing.T) {
	cpu := InitCPU()
	cpu.Registers[8] = 0x7FFFFFFF
	cpu.Registers[9] = 0x7FFFFFFF

	// addu $t2, $t1, $t0
	cpu.Instruction = 0x01095021
	cpu.DecodeInstr()

	if cpu.Registers[10] != -2 {
		t.Errorf("destination register wrong. expected=%d, got=%d", -2, cpu.Registers[10])
	}
}

func TestSubtractU(t *testing.T) {
	cpu := InitCPU()
	cpu.Registers[8] = -0x7FFFFFFF
	cpu.Registers[9] = 0x7FFFFFFF

	// subu $t2, $t1, $t0
	cpu.Instruction = 0x01285023
	cpu.DecodeInstr()
	if cpu.Registers[10] != -2 {
		t.Errorf("destination register wrong. expected=%d, got=%d", -2, cpu.Registers[10])
	}
}

func TestMultiply(t *testing.T) {
	cpu := InitCPU()
	cpu.Registers[8] = 10
	cpu.Registers[9] = 5

	// mult $t1, $t0

	cpu.Instruction = 0x01280018
	cpu.DecodeInstr()
	if cpu.HiLow.lo != 50 {
		t.Errorf("lo register wrong. expected=%d, got=%d", 50, cpu.HiLow.lo)
	}

	if cpu.HiLow.hi != 0 {
		t.Errorf("hi register wrong. expected=%d, got=%d", 0, cpu.HiLow.hi)
	}
}

func TestDivide(t *testing.T) {
	cpu := InitCPU()
	cpu.Registers[8] = 5
	cpu.Registers[9] = 10

	// div $t1, $t0
	cpu.Instruction = 0x0128001A
	cpu.DecodeInstr()
	if cpu.HiLow.lo != 2 {
		t.Errorf("lo register wrong. expected=%d, got=%d", 2, cpu.HiLow.lo)
	}

	if cpu.HiLow.hi != 0 {
		t.Errorf("hi register wrong. expected=%d, got=%d", 0, cpu.HiLow.hi)
	}

	cpu.Registers[8] = 0
	err := cpu.DecodeInstr()
	if err != nil {
		t.Fatalf("expected divide by zero exception")
	}
}
