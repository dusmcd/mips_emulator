package cpu

import (
	"testing"
	"mips_emulator/defs"
	"mips_emulator/memory"
)

func TestSlt(t *testing.T) {
	cpu := InitCPU(memory.InitMemory(), 0)
	cpu.Registers[8] = 10 // setting $t0
	cpu.Registers[16] = 5 // setting $s0

	// slt $s1, $s0, $t0
	cpu.Instruction = 0x0208882a
	cpu.DecodeInstr()

	if cpu.Registers[17] != 1 {
		t.Errorf("destination register wrong. expected=%d, got=%d", 1, cpu.Registers[17])
	}

	cpu.Registers[16] = 10
	cpu.DecodeInstr()

	if cpu.Registers[17] != 0 {
		t.Errorf("destination register wrong. expected=%d, got=%d", 0, cpu.Registers[17])
	}
}

func TestSltu(t *testing.T) {
	cpu := InitCPU(memory.InitMemory(), 0)
	cpu.Registers[8] = 10 // setting $t0
	cpu.Registers[16] = 5 // setting $s0

	// sltu $s1, $s0, $t0
	cpu.Instruction = 0x0208882b
	cpu.DecodeInstr()

	if cpu.Registers[17] != 1 {
		t.Errorf("destination register wrong. expected=%d, got=%d", 1, cpu.Registers[17])
	}

	cpu.Registers[16] = 10
	cpu.DecodeInstr()

	if cpu.Registers[17] != 0 {
		t.Errorf("destination register wrong. expected=%d, got=%d", 0, cpu.Registers[17])
	}
}


func TestSll(t *testing.T) {
	cpu := InitCPU(memory.InitMemory(), 0)
	cpu.Registers[8] = 0x00000FFF

	// sll $t2, $t0, 16
	cpu.Instruction = 0x00085400
	cpu.DecodeInstr()

	if cpu.Registers[10] != 0x00000FFF << 16 {
		t.Errorf("destination register wrong. exected=%d, got=%d", 0x00000FFF << 16, cpu.Registers[10])
	}
}

func TestSra(t *testing.T) {
	cpu := InitCPU(memory.InitMemory(), 0)
	cpu.Registers[8] = -1024

	// sra $t2, $t0, 1
	cpu.Instruction = 0x00085043
	cpu.DecodeInstr()

	if cpu.Registers[10] != -512 {
		t.Errorf("destination register wrong. expected=%d, got=%d", -512, cpu.Registers[10])
	}
}

func TestSrl(t *testing.T) {
	cpu := InitCPU(memory.InitMemory(), 0)
	cpu.Registers[8] = -1024

	// srl $t2, $t0, 1
	cpu.Instruction = 0x00085042
	cpu.DecodeInstr()

	if cpu.Registers[10] != 0x7FFFFE00 {
		t.Errorf("destination register wrong. expected=%d, got=%d", 0x7FFFFE00, cpu.Registers[10])
	}
}

func TestAnd(t *testing.T) {
	cpu := InitCPU(memory.InitMemory(), 0)
	cpu.Registers[8] = 5
	cpu.Registers[9] = 10

	// and $t2, $t1, $t0
	cpu.Instruction = 0x01285024
	cpu.DecodeInstr()

	if cpu.Registers[10] != 5 & 10 {
		t.Errorf("destination register wrong. expected=%d, got=%d", 5 & 10, cpu.Registers[10])
	}

}

func TestOr(t *testing.T) {
	cpu := InitCPU(memory.InitMemory(), 0)
	cpu.Registers[8] = 5
	cpu.Registers[9] = 10

	// or $t2, $t1, $t0
	cpu.Instruction = 0x01285025
	cpu.DecodeInstr()

	if cpu.Registers[10] != 5 | 10 {
		t.Errorf("destination register wrong. expected=%d, got=%d", 5 | 10, cpu.Registers[10])
	}

}

func TestXor(t *testing.T) {
	cpu := InitCPU(memory.InitMemory(), 0)
	cpu.Registers[8] = 5
	cpu.Registers[9] = 10

	// xor $t2, $t1, $t0
	cpu.Instruction = 0x01285026
	cpu.DecodeInstr()

	if cpu.Registers[10] != 5 ^ 10 {
		t.Errorf("destination register wrong. expected=%d, got=%d", 5 ^ 10, cpu.Registers[10])
	}

}

func TestNor(t *testing.T) {
	cpu := InitCPU(memory.InitMemory(), 0)
	cpu.Registers[8] = 5
	cpu.Registers[9] = 10

	// or $t2, $t1, $t0
	cpu.Instruction = 0x01285027
	cpu.DecodeInstr()

	if cpu.Registers[10] != ^(5 | 10) {
		t.Errorf("destination register wrong. expected=%d, got=%d", ^(5 | 10), cpu.Registers[10])
	}

}

func TestMultiplyU(t *testing.T) {
	cpu := InitCPU(memory.InitMemory(), 0)
	cpu.Registers[8] = -1
	cpu.Registers[9] = 2

	
	cpu.Instruction = 0x01285019
	cpu.DecodeInstr()
	var signedInt defs.Word = -2
	expected := uint32(signedInt)

	hi := defs.Word(uint(expected) & 0xFFFFFFFF00000000 >> 32)
	lo := defs.Word(uint(expected) & 0x00000000FFFFFFFF)

	if cpu.HiLow.lo != lo {
		t.Errorf("lo register wrong. expected=%d, got=%d", lo, cpu.HiLow.lo)
	}

	if cpu.HiLow.hi != hi {
		t.Errorf("hi register wrong. expected=%d, got=%d", hi, cpu.HiLow.hi)
	}

}

func TestDivideU(t *testing.T) {
	cpu := InitCPU(memory.InitMemory(), 0)
	cpu.Registers[8] = 3
	cpu.Registers[9] = 10

	cpu.Instruction = 0x0128501B
	cpu.DecodeInstr()

	if cpu.HiLow.hi != 1 {
		t.Errorf("hi register wrong. expected=%d, got=%d", 1, cpu.HiLow.hi)
	}
	if cpu.HiLow.lo != 3 {
		t.Errorf("lo register wrong. expected=%d, got=%d", 3, cpu.HiLow.lo)
	}
}


func TestAdd(t *testing.T) {
	cpu := InitCPU(memory.InitMemory(), 0)
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
	cpu := InitCPU(memory.InitMemory(), 0)

	cpu.Registers[8] = 300 // set register $t0
	cpu.Registers[9] = 100 // set register $t1

	// sub $t2, $t1, $t0
	cpu.Instruction = 0x01285022
	cpu.DecodeInstr()
	
	if cpu.Registers[10] != -200 {
		t.Errorf("destination register wrong. expected=%d, got=%d", -200, cpu.Registers[10])
	}

	cpu.Registers[8] = 100 // set register $t0
	cpu.Registers[9] = 300 // set register $t1
	cpu.DecodeInstr()
	if cpu.Registers[10] != 200 {
		t.Errorf("destination register wrong. expected=%d, got=%d", 200, cpu.Registers[10])
	}

}

func TestAddU(t *testing.T) {
	cpu := InitCPU(memory.InitMemory(), 0)
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
	cpu := InitCPU(memory.InitMemory(), 0)
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
	cpu := InitCPU(memory.InitMemory(), 0)
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
	cpu := InitCPU(memory.InitMemory(), 0)
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
	if err == nil {
		t.Fatalf("expected divide by zero exception")
	}
}

