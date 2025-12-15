package cpu

import (
	"errors"
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
	cpu.HiLow.lo = int32(op1 / op2)
	cpu.HiLow.hi = int32(op1 % op2)

	return nil
}

func (cpu * CPU) multUInstr(rs, rt, rd, shift uint8) error {
	op1 := uint32(cpu.Registers[rs])
	op2 := uint32(cpu.Registers[rt])
	product := int(op1 * op2)
	cpu.HiLow.hi = int32(uint(product) & uint(0xFFFFFFFF00000000) >> 32)
	cpu.HiLow.lo = int32(uint(product) & uint(0x00000000FFFFFFFF))

	return nil
}

func (cpu *CPU) multInstr(rs, rt, rd, shift uint8) error {
	op1 := cpu.Registers[rs]
	op2 := cpu.Registers[rt]
	product := int(op1 * op2)
	cpu.HiLow.hi = int32(uint(product) & uint(0xFFFFFFFF00000000) >> 32)
	cpu.HiLow.lo = int32(uint(product) & uint(0x00000000FFFFFFFF))

	return nil
}

func (cpu *CPU) subUInstr(rs, rt, rd, shift uint8) error {
	op1 := uint32(cpu.Registers[rs])
	op2 := uint32(cpu.Registers[rt])
	cpu.Registers[rd] = int32(op1 - op2)

	return nil
}

func (cpu *CPU) subInstr(rs, rt, rd, shift uint8) error {
	op1 := cpu.Registers[rs]
	op2 := cpu.Registers[rt]
	check := int(op1 - op2)
	if check > MAX32 {
		return errors.New("signed overflow exception")
	}
	cpu.Registers[rd] = int32(check)

	return nil
}

func (cpu *CPU) addUInstr(rs, rt, rd, shift uint8) error {
		op1 := uint32(cpu.Registers[rs])
		op2 := uint32(cpu.Registers[rt])
		cpu.Registers[rd] = int32(op1 + op2)
		return nil
}

func (cpu *CPU) addInstr(rs, rt, rd, shift uint8) error {
	op1 := cpu.Registers[rs]
	op2 := cpu.Registers[rt]
	check := int(op1 + op2)
	if check > MAX32 {
		return errors.New("signed overflow exception")
	}
	cpu.Registers[rd] = int32(check)

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

