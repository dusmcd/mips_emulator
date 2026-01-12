package cpu

import (
	"testing"
	"mips_emulator/defs"
)

func TestAddiu(t *testing.T) {
	cpu := InitCPU(true)
	cpu.Registers[16] = 0x000FFFFF // setting $s0

	// addiu $t0, $s0, 1024
	cpu.Instruction = 0x26080400
	cpu.DecodeInstr()

	expected := defs.Word(0x000FFFFF + 1024)
	if cpu.Registers[8] != expected {
		t.Errorf("destination register wrong. expected=%d, got=%d",expected, cpu.Registers[8])
	}
}

func TestAndi(t *testing.T) {
	cpu := InitCPU(true)
	cpu.Registers[16] = 1024 // setting $s0

	// andi $t0, $s0, 512
	cpu.Instruction = 0x32080200
	cpu.DecodeInstr()

	expected := defs.Word(1024 & 512)
	if cpu.Registers[8] != expected {
		t.Errorf("destination register wrong. expected=%d, got=%d", expected, cpu.Registers[8])
	}
}

func TestXori(t *testing.T) {
	cpu := InitCPU(true)
	cpu.Registers[16] = 1024 // setting $s0

	// xori $t0, $s0, 16000
	cpu.Instruction = 0x3a083e80
	cpu.DecodeInstr()

	expected := defs.Word(1024 ^ 16000)
	if cpu.Registers[8] != expected {
		t.Errorf("destination register wrong. expected=%d, got=%d", expected, cpu.Registers[8])
	}
}

func TestOri(t *testing.T) {
	cpu := InitCPU(true)
	cpu.Registers[16] = 1000000 // setting $s0

	// ori $t0, $s0, 10
	cpu.Instruction = 0x3608000a 
	cpu.DecodeInstr()

	expected := defs.Word(1000000 | 10)
	if cpu.Registers[8] != expected {
		t.Errorf("destination register wrong. expected=%d, got=%d", expected, cpu.Registers[8])
	}
}

func TestBne(t *testing.T) {
	cpu := InitCPU(true)
	cpu.PC = 0x04
	cpu.Registers[16] = 20 // setting $s0
	cpu.Registers[8] = 10 // setting $t0

	// bne $s0, $t0, done
	cpu.Instruction = 0x1608000e
	cpu.DecodeInstr()
	var targetPC uint32 = (0xe << 2) + 0x04

	if cpu.PC != targetPC {
		t.Errorf("PC wrong. expected=%d, got=%d", targetPC, cpu.PC)
	}

	cpu.Registers[16] = 10 // setting $s0
	cpu.Registers[8] = 10 // setting $t0

	cpu.DecodeInstr()
	if cpu.PC != targetPC {
		t.Errorf("PC wrong. expected=%d, got=%d", targetPC, cpu.PC)
	}
}

func TestBgtz(t *testing.T) {
	cpu := InitCPU(true)
	cpu.PC = 0x04
	cpu.Registers[16] = 10 // setting $s0

	// bgtz $s0, done
	cpu.Instruction = 0x1e00000e
	cpu.DecodeInstr()

	var targetPC uint32 = (0xe << 2) + 0x04
	if cpu.PC != targetPC {
		t.Errorf("PC wrong. expected=%d, got=%d", targetPC, cpu.PC)
	}

	cpu.Registers[16] = 0 // setting $s0
	cpu.DecodeInstr()
	if cpu.PC != targetPC {
		t.Errorf("PC wrong. expected=%d, got=%d", targetPC, cpu.PC)
	}

}

func TestBlez(t *testing.T) {
	cpu := InitCPU(true)
	cpu.PC = 0x04
	cpu.Registers[16] = 0 // setting $s0

	// blez $s0, done
	cpu.Instruction = 0x1a00000e
	cpu.DecodeInstr()

	var targetPC uint32 = (0xe << 2) + 0x04
	if cpu.PC != targetPC {
		t.Errorf("PC wrong. expected=%d, got=%d", targetPC, cpu.PC)
	}

	cpu.Registers[16] = -1024 // setting $s0
	targetPC += (0xe << 2)
	cpu.DecodeInstr()
	if cpu.PC != targetPC {
		t.Errorf("PC wrong. expected=%d, got=%d", targetPC, cpu.PC)
	}

	cpu.Registers[16] = 1024 // setting $s0
	cpu.DecodeInstr()
	if cpu.PC != targetPC {
		t.Errorf("PC wrong. expected=%d, got=%d", targetPC, cpu.PC)
	}

}

func TestBeq(t *testing.T) {
	cpu := InitCPU(true)
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

	cpu.Registers[16] = 15 // setting $s0
	cpu.Registers[8] = 10 // setting $t0
	cpu.DecodeInstr()
	if cpu.PC != targetPC {
		t.Errorf("PC wrong. expected=%d, got=%d", targetPC, cpu.PC)
	}


}

func TestAddi(t *testing.T) {
	cpu := InitCPU(true)
	cpu.Registers[8] = 5; // writing to $t0

	// addi $s0, $t0, 4095
	cpu.Instruction = 0x21100FFF
	cpu.DecodeInstr()

	if cpu.Registers[16] != 4095+5 {
		t.Errorf("destination register wrong. expected=%d, got=%d", 4095+5, cpu.Registers[16])
	}
}

func TestLW(t *testing.T) {
	cpu := InitCPU(true)
	cpu.Registers[16] = 0xFF0 // writing to $s0

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
	cpu := InitCPU(true)
	cpu.Registers[16] = 0xFF0 // writing to $s0
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
	
