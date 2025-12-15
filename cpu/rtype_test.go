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
