package cpu

import (
	"errors"
	"mips_emulator/defs"
)

const (
	SYS_EXIT = 4001
)

type RFunc func(rs, rt, rd, shift uint8) error

func (cpu *CPU) mthiInstr(rs, rt, rd, shift uint8) error {
	val := cpu.Registers[rs]
	cpu.HiLow.hi = val

	return nil
}

func (cpu *CPU) mtloInstr(rs, rt, rd, shift uint8) error {
	val := cpu.Registers[rs]
	cpu.HiLow.lo = val

	return nil
}

func (cpu *CPU) mulInstr(rs, rt, rd, shift uint8) error {
	op1 := signExtend(cpu.Registers[rs])
	op2 := signExtend(cpu.Registers[rt])
	product := op1 * op2
	hiBits := defs.Word(uint(product) & uint(0xFFFFFFFF00000000) >> 32)
	loBits := defs.Word(uint(product) & uint(0x00000000FFFFFFFF))

	cpu.Registers[rd] = loBits
	cpu.HiLow.lo = loBits
	cpu.HiLow.hi = hiBits

	return nil
}

func (cpu *CPU) msubuInstr(rs, rt, rd, shift uint8) error {
	op1 := uint64(uint(cpu.Registers[rs]) & 0x00000000FFFFFFFF)
	op2 := uint64(uint(cpu.Registers[rt]) & 0x00000000FFFFFFFF)
	product := op1 * op2
	hiBits := defs.Word(uint(product) & uint(0xFFFFFFFF00000000) >> 32)
	loBits := defs.Word(uint(product) & uint(0x00000000FFFFFFFF))

	cpu.HiLow.hi -= hiBits
	cpu.HiLow.lo -= loBits

	return nil
}

func (cpu *CPU) msubInstr(rs, rt, rd, shift uint8) error {
	op1 := signExtend(cpu.Registers[rs])
	op2 := signExtend(cpu.Registers[rt])
	product := op1 * op2
	hiBits := defs.Word(uint(product) & uint(0xFFFFFFFF00000000) >> 32)
	loBits := defs.Word(uint(product) & uint(0x00000000FFFFFFFF))

	cpu.HiLow.hi -= hiBits
	cpu.HiLow.lo -= loBits

	return nil
}

func signExtend(num defs.Word) int64 {
	if num < 0 {
		return int64(num)
	}
	return int64(uint(num) & 0x00000000FFFFFFFF)
}

func (cpu *CPU) maddInstr(rs, rt, rd, shift uint8) error {
	op1 := signExtend(cpu.Registers[rs])
	op2 := signExtend(cpu.Registers[rt])

	product := op1 * op2	
	hiBits := defs.Word(uint(product) & uint(0xFFFFFFFF00000000) >> 32)
	loBits := defs.Word(uint(product) & uint(0x00000000FFFFFFFF))

	cpu.HiLow.hi += hiBits
	cpu.HiLow.lo += loBits

	return nil
}

func (cpu *CPU) madduInstr(rs, rt, rd, shift uint8) error {
	op1 := uint64(cpu.Registers[rs]) & 0x00000000FFFFFFFF
	op2 := uint64(cpu.Registers[rt]) & 0x00000000FFFFFFFF
	product := op1 * op2
	hiBits := defs.Word(uint(product) & uint(0xFFFFFFFF00000000) >> 32)
	loBits := defs.Word(uint(product) & uint(0x00000000FFFFFFFF))

	cpu.HiLow.hi += hiBits
	cpu.HiLow.lo += loBits

	return nil
}



func (cpu *CPU) jrInstr(rs, rt, rd, shift uint8) error {
	addr := cpu.Registers[rs]
	cpu.PC = uint32(addr)

	return nil
}

func (cpu *CPU) jalrInstr(rs, rt, rd, shift uint8) error {
	addr := cpu.Registers[rs]
	cpu.Registers[31] = defs.Word(cpu.PC) // setting return address
	cpu.PC = uint32(addr)

	return nil
}

func (cpu *CPU) clzInstr(rs, rt, rd, shift uint8) error {
	op1 := cpu.Registers[rs]

	if rd == 0 {
		return errors.New("cannot write to $zero register")
	}

	// only positive numbers will have leading zeros in signed integers
	if op1 < 0 {
		cpu.Registers[rd] = 0
		return nil
	}

	var bitMask uint32 = 0xF0000000
	shiftAmt := 28
	current := (uint32(op1) & bitMask) >> shiftAmt
	count := 0

	for current <= 0x7 {
		if current == 0x0 {
			count += 4
		} else if current == 0x1 {
			count += 3
		} else if current <= 0x3 {
			count += 2
		} else {
			count++
		}

		if current > 0x0 {
			break
		}

		shiftAmt -= 4
		if shiftAmt < 0 {
			break
		}

		bitMask = bitMask >> 4
		current = (uint32(op1) & bitMask) >> shiftAmt
	}

	cpu.Registers[rd] = defs.Word(count)
	return nil
}

func (cpu *CPU) cloInstr(rs, rt, rd, shift uint8) error {
	op1 := cpu.Registers[rs]

	if rd == 0 {
		return errors.New("cannot write to $zero register")
	}

	// only negative numbers will have a leading one in signed integers
	if op1 >= 0 {
		cpu.Registers[rd] = 0
		return nil
	}
	var bitMask uint32 = 0xF0000000
	shiftAmt := 28
	current := (uint32(op1) & bitMask) >> shiftAmt
	count := 0

	for current >= 0x8 {
		if current == 0xF {
			count += 4
		} else if current == 0xE {
			count += 3
		} else if current >= 0xC {
			count += 2
		} else {
			count++
		}

		if current < 0xF {
			break
		}

		shiftAmt -= 4
		if shiftAmt < 0 {
			break
		}

		bitMask = bitMask >> 4
		current = (uint32(op1) & bitMask) >> shiftAmt
	}

	cpu.Registers[rd] = defs.Word(count)
	
	return nil
}

func (cpu *CPU) sltuInstr(rs, rt, rd, shift uint8) error {
	op1 := uint32(cpu.Registers[rs])
	op2 := uint32(cpu.Registers[rt])

	if rd == 0 {
		return errors.New("cannot write to $zero register")
	}

	if op1 < op2 {
		cpu.Registers[rd] = 1
	} else {
		cpu.Registers[rd] = 0
	}

	return nil
}

func (cpu *CPU) sltInstr(rs, rt, rd, shift uint8) error {
	op1 := cpu.Registers[rs]
	op2 := cpu.Registers[rt]

	if rd == 0 {
		return errors.New("cannot write to $zero register")
	}

	if op1 < op2 {
		cpu.Registers[rd] = 1
	} else {
		cpu.Registers[rd] = 0
	}
	return nil
}

func (cpu *CPU) syscall(rs, rt, rd, shift uint8) error {
	// switch on value in $v0 register
	switch(cpu.Registers[2]) {
	case SYS_EXIT:
		cpu.Exit = true
	}
	return nil
}

func (cpu *CPU) sllInstr(rs, rt, rd, shift uint8) error {
	op1 := cpu.Registers[rt]
	if rd == 0 {
		return errors.New("cannot write to $zero register")
	}
	cpu.Registers[rd] = op1 << shift
	return nil
}

func (cpu *CPU) srlInstr(rs, rt, rd, shift uint8) error {
	op1 := cpu.Registers[rt]

	// casting to unsigned so that shift does *not* sign-extend
	result := uint32(op1) >> shift
	if rd == 0 {
		return errors.New("cannot write to $zero register")
	}

	cpu.Registers[rd] = defs.Word(result)

	return nil
}

func (cpu *CPU) sraInstr(rs, rt, rd, shift uint8) error {
	op1 := cpu.Registers[rt]
	if rd == 0 {
		return errors.New("cannot write to $zero register")
	}

	// this shift will sign-extend because op1 is a signed integer
	cpu.Registers[rd] = op1 >> shift
	return nil
}

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
	op1 := uint64(uint(cpu.Registers[rs]) & 0x00000000FFFFFFFF)
	op2 := uint64(uint(cpu.Registers[rt]) & 0x00000000FFFFFFFF)
	product := op1 * op2
	cpu.HiLow.hi = defs.Word(uint(product) & uint(0xFFFFFFFF00000000) >> 32)
	cpu.HiLow.lo = defs.Word(uint(product) & uint(0x00000000FFFFFFFF))

	return nil
}

func (cpu *CPU) multInstr(rs, rt, rd, shift uint8) error {
	op1 := signExtend(cpu.Registers[rs])
	op2 := signExtend(cpu.Registers[rt])
	product := op1 * op2
	cpu.HiLow.hi = defs.Word(uint(product) & uint(0xFFFFFFFF00000000) >> 32)
	cpu.HiLow.lo = defs.Word(uint(product) & uint(0x00000000FFFFFFFF))

	return nil
}

func (cpu *CPU) subUInstr(rs, rt, rd, shift uint8) error {
	op1 := uint32(cpu.Registers[rs])
	op2 := uint32(cpu.Registers[rt])
	
	if rd == 0 {
		return errors.New("cannot write to $zero register")
	}

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
	if rd == 0 {
		return errors.New("cannot write to $zero register")
	}

	cpu.Registers[rd] = defs.Word(check)

	return nil
}

func (cpu *CPU) addUInstr(rs, rt, rd, shift uint8) error {
		op1 := uint32(cpu.Registers[rs])
		op2 := uint32(cpu.Registers[rt])
		if rd == 0 {
			return errors.New("cannot write to $zero register")
		}

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
	if rd == 0 {
		return errors.New("cannot write to $zero register")
	}


	cpu.Registers[rd] = defs.Word(check)

	return nil
}

func (cpu *CPU) andInstr(rs, rt, rd, shift uint8) error {
	op1 := cpu.Registers[rs]
	op2 := cpu.Registers[rt]
	if rd == 0 {
		return errors.New("cannot write to $zero register")
	}

	cpu.Registers[rd] = op1 & op2

	return nil
}

func (cpu *CPU) orInstr(rs, rt, rd, shift uint8) error {
	op1 := cpu.Registers[rs]
	op2 := cpu.Registers[rt]
	if rd == 0 {
		return errors.New("cannot write to $zero register")
	}

	cpu.Registers[rd] = op1 | op2

	return nil
}

func (cpu *CPU) xorInstr(rs, rt, rd, shift uint8) error {
	op1 := cpu.Registers[rs]
	op2 := cpu.Registers[rt]
	if rd == 0 {
		return errors.New("cannot write to $zero register")
	}

	cpu.Registers[rd] = op1 ^ op2

	return nil
}

func (cpu *CPU) norInstr(rs, rt, rd, shift uint8) error {
	op1 := cpu.Registers[rs]
	op2 := cpu.Registers[rt]
	if rd == 0 {
		return errors.New("cannot write to $zero register")
	}

	cpu.Registers[rd] = ^(op1 | op2)

	return nil
}

