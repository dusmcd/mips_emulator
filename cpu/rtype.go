package cpu

import (
	"errors"
	"mips_emulator/defs"
)

type RFunc func(rs, rt, rd, shift uint8) error

var funcMap map[uint8]RFunc 

func (cpu *CPU) divInstr(rs, rt, rd, shift uint8) error {
	op1 := cpu.Registers[rs]
	op2 := cpu.Registers[rt]
	if op2 == 0 {
		return errors.New("divide by zero exception")
	}

	cpu.HiLow.lo = op1 / op2
	cpu.HiLow.hi = op1 % op2

	return nil
}

func (cpu *CPU) divUInstr(rs, rt, rd, shift uint8) error {
	op1 := uint32(cpu.Registers[rs])
	op2 := uint32(cpu.Registers[rt])

	if op2 == 0 {
		return errors.New("divide by zero exception")
	}
	cpu.HiLow.lo = defs.Word(op1 / op2)
	cpu.HiLow.hi = defs.Word(op1 % op2)

	return nil
}

func (cpu * CPU) multUInstr(rs, rt, rd, shift uint8) error {
	op1 := uint32(cpu.Registers[rs])
	op2 := uint32(cpu.Registers[rt])
	product := op1 * op2
	cpu.HiLow.hi = defs.Word(uint(product) & uint(0xFFFFFFFF00000000) >> 32)
	cpu.HiLow.lo = defs.Word(uint(product) & uint(0x00000000FFFFFFFF))

	return nil
}

func (cpu *CPU) multInstr(rs, rt, rd, shift uint8) error {
	op1 := cpu.Registers[rs]
	op2 := cpu.Registers[rt]
	product := op1 * op2
	cpu.HiLow.hi = defs.Word(uint(product) & uint(0xFFFFFFFF00000000) >> 32)
	cpu.HiLow.lo = defs.Word(uint(product) & uint(0x00000000FFFFFFFF))

	return nil
}

func (cpu *CPU) subUInstr(rs, rt, rd, shift uint8) error {
	op1 := uint32(cpu.Registers[rs])
	op2 := uint32(cpu.Registers[rt])
	cpu.Registers[rd] = defs.Word(op1 - op2)

	return nil
}

func isOverflow(op1, op2, result defs.Word) bool {
	if (op1 < 0 && op2 < 0 && result > 0) ||
		(op1 > 0 && op2 > 0 && result < 0) {
			return true
		}
	return false
}

func (cpu *CPU) subInstr(rs, rt, rd, shift uint8) error {
	op1 := cpu.Registers[rs]
	op2 := cpu.Registers[rt]
	check := op1 - op2
	if (op2 < 0 && op1 > 0 && check < 0) {
		return errors.New("signed overflow exception")
	}
	cpu.Registers[rd] = defs.Word(check)

	return nil
}

func (cpu *CPU) addUInstr(rs, rt, rd, shift uint8) error {
		op1 := uint32(cpu.Registers[rs])
		op2 := uint32(cpu.Registers[rt])
		cpu.Registers[rd] = defs.Word(op1 + op2)
		return nil
}

func (cpu *CPU) addInstr(rs, rt, rd, shift uint8) error {
	op1 := cpu.Registers[rs]
	op2 := cpu.Registers[rt]
	check := op1 + op2

	if isOverflow(op1, op2, check) {
		return errors.New("signed overflow exception")
	}

	cpu.Registers[rd] = defs.Word(check)

	return nil
}

func (cpu *CPU) andInstr(rs, rt, rd, shift uint8) error {
	op1 := cpu.Registers[rs]
	op2 := cpu.Registers[rt]
	cpu.Registers[rd] = op1 & op2

	return nil
}

func (cpu *CPU) orInstr(rs, rt, rd, shift uint8) error {
	op1 := cpu.Registers[rs]
	op2 := cpu.Registers[rt]
	cpu.Registers[rd] = op1 | op2

	return nil
}

func (cpu *CPU) xorInstr(rs, rt, rd, shift uint8) error {
	op1 := cpu.Registers[rs]
	op2 := cpu.Registers[rt]
	cpu.Registers[rd] = op1 ^ op2

	return nil
}

func (cpu *CPU) norInstr(rs, rt, rd, shift uint8) error {
	op1 := cpu.Registers[rs]
	op2 := cpu.Registers[rt]
	cpu.Registers[rd] = ^(op1 | op2)

	return nil
}

