package cpu

import (
	"mips_emulator/defs"
)

type JInstr func(uint32) error

func (cpu *CPU) jInstr(addr uint32) error {
	addr = addr << 2

	// concat the 4 high bits of the PC
	addr = addr | (cpu.PC & 0xF0000000)
	cpu.PC = addr
	return nil
}

func (cpu *CPU) jalInstr(addr uint32) error {
	addr = addr << 2
	addr = addr | (cpu.PC & 0xF0000000)
	cpu.Registers[31] = defs.Word(cpu.PC)
	cpu.PC = addr

	return nil
}
