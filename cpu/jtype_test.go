package cpu

import (
	"testing"
	"mips_emulator/memory"
)

func TestJump(t *testing.T) {
	cpu := InitCPU(memory.InitMemory(), 0)
	cpu.PC = 0x04

	// j done
	cpu.Instruction = 0x080003e8
	cpu.DecodeInstr()

	if cpu.PC != 4000 {
		t.Errorf("PC wrong. expected=%d, got=%d", 4000, cpu.PC)
	}
}
