package cpu

import (
	"mips_emulator/defs"
	"errors"
)

type IInstr func(rs, rt uint8, imm int16) error

func (cpu CPU) lwInstr(rs, rt uint8, imm int16) error {
	baseAddr := cpu.Registers[rs]
	fullAddr := uint32(baseAddr + defs.Word(imm))
	memoryVal, err := cpu.MainMemory.LoadWord(fullAddr)
	if err != nil {
		return err
	}

	cpu.Registers[rt] = memoryVal
	return nil
}

func (cpu *CPU) swInstr(rs, rt uint8, imm int16) error {
	baseAddr := cpu.Registers[rs]
	fullAddr := uint32(baseAddr + defs.Word(imm))
	err := cpu.MainMemory.StoreWord(fullAddr, cpu.Registers[rt])

	if err != nil {
		return err
	}
	return nil
}

func (cpu *CPU) addiInstr(rs, rt uint8, imm int16) error {
	op1 := cpu.Registers[rs]
	check := op1 + defs.Word(imm)
	if isOverflow(op1, defs.Word(imm), check) {
		return errors.New("signed overflow exception")
	}

	cpu.Registers[rt] = check
	return nil
}
