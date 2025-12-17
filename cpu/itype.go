package cpu

import (
	"mips_emulator/defs"
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
