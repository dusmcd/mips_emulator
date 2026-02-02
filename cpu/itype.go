package cpu

import (
	"mips_emulator/defs"
	"errors"
)

const (
	BLTZ = 0x00
	BGEZ = 0x01
	BLTZAL = 0x10
	BGEZAL = 0x11
)

type IInstr func(rs, rt uint8, imm int16) error

/**
* load the portion of the word starting at the effective address
* and ending at the most significant byte of the word
*/
func (cpu *CPU) lwlInstr(rs, rt uint8, imm int16) error {
	if rt == 0 {
		return errors.New("cannot write to $zero register")
	}

	baseAddr := cpu.Registers[rs]
	fullAddr := defs.Word(imm) + baseAddr
	k := fullAddr % 4
	size := int(4 - k)

	for i := range size {
		addr := uint32(fullAddr) + uint32(i)
		current, err := cpu.MainMemory.LoadByte(addr)	
		if err != nil {
			return err
		}

		bitMask := uint32(current) & (0x000000FF << (4 - size + i))
		cpu.Registers[rt] |= defs.Word(bitMask)
	}


	return nil
}

/**
* load the portion of the word starting at the effective address
* and ending at the least significant byte of the word
*/
func (cpu *CPU) lwrInstr(rs, rt uint8, imm int16) error {
	if rt == 0 {
		return errors.New("cannot write to $zero register")
	}
	baseAddr := cpu.Registers[rs]
	fullAddr := defs.Word(imm) + baseAddr
	k := fullAddr % 4
	size := k + 1

	for i := range size {
		addr := uint32(fullAddr) - uint32(i)
		current, err := cpu.MainMemory.LoadByte(addr)
		if err != nil {
			return err
		}

		bitMask := uint32(current) & (0xFF000000 >> (4 - size + i))
		cpu.Registers[rt] |= defs.Word(bitMask)
	}

	return nil
}

func (cpu *CPU) luiInstr(rs, rt uint8, imm int16) error {
	if rt == 0 {
		return errors.New("cannot write to $zero register")
	}

	// this will ensure that the low bits are zero
	cpu.Registers[rt] = 0

	// loading high 16 bits
	cpu.Registers[rt] |= defs.Word(imm) << 16

	return nil
}

func (cpu *CPU) lhuInstr(rs, rt uint8, imm int16) error {
	baseAddr := cpu.Registers[rs]
	fullAddr := defs.Word(imm) + baseAddr

	if rt == 0 {
		return errors.New("cannot write to $zero register")
	}

	memoryVal, err := cpu.MainMemory.LoadHalfWord(uint32(fullAddr))
	if err != nil {
		return err
	}

	result := defs.Word(memoryVal) & 0x0000FFFF
	cpu.Registers[rt] = result

	return nil
}

func (cpu *CPU) lhInstr(rs, rt uint8, imm int16) error {
	baseAddr := cpu.Registers[rs]
	fullAddr := defs.Word(imm) + baseAddr

	if rt == 0 {
		return errors.New("cannot write to $zero register")
	}

	memoryVal, err := cpu.MainMemory.LoadHalfWord(uint32(fullAddr))
	if err != nil {
		return err
	}

	cpu.Registers[rt] = defs.Word(memoryVal)
	return nil
}

func (cpu *CPU) lbuInstr(rs, rt uint8, imm int16) error {
	baseAddr := cpu.Registers[rs]
	fullAddr := defs.Word(imm) + baseAddr

	if rt == 0 {
		return errors.New("cannot write to $zero register")
	}

	memoryVal, err := cpu.MainMemory.LoadByte(uint32(fullAddr))
	if err != nil {
		return err
	}

	val := defs.Word(memoryVal) & 0x000000FF
	cpu.Registers[rt] = val

	return nil
}

func (cpu *CPU) lbInstr(rs, rt uint8, imm int16) error {
	baseAddr := cpu.Registers[rs]
	fullAddr := defs.Word(imm) + baseAddr

	if rt == 0 {
		return errors.New("cannot write to $zero register")
	}

	memoryVal, err := cpu.MainMemory.LoadByte(uint32(fullAddr))

	if err != nil {
		return err
	}
	cpu.Registers[rt] = defs.Word(memoryVal)


	return nil
}

func (cpu *CPU) regImm(rs, rt uint8, imm int16) error {
	check := cpu.Registers[rs]
	newPC := defs.Word(cpu.PC) + defs.Word(imm << 2)

	switch(cpu.Registers[rt]) {
	case BLTZ:
		if check < 0 {
			cpu.PC = uint32(newPC)
		}
		break
	case BGEZ:
		if check >= 0 {
			cpu.PC = uint32(newPC)
		}
		break
	case BLTZAL:
		if check < 0 {
			cpu.Registers[31] = defs.Word(cpu.PC)
			cpu.PC = uint32(newPC)
		}
		break
	case BGEZAL:
		if check >= 0 {
			cpu.Registers[31] = defs.Word(cpu.PC)
			cpu.PC = uint32(newPC)
		}
		break
	}
	return nil
}

func (cpu *CPU) oriInstr(rs, rt uint8, imm int16) error {
	op1 := cpu.Registers[rs]

	if rt == 0 {
		return errors.New("cannot write to $zero register")
	}

	cpu.Registers[rt] = op1 | defs.Word(imm)
	return nil
}

func (cpu *CPU) xoriInstr(rs, rt uint8, imm int16) error {
	op1 := cpu.Registers[rs]

	if rt == 0 {
		return errors.New("cannot write to $zero register")
	}

	cpu.Registers[rt] = op1 ^ defs.Word(imm)
	return nil
}

func (cpu *CPU) bgtzInstr(rs, rt uint8, imm int16) error {
	op1 := cpu.Registers[rs]

	if op1 > 0 {
		newPC := defs.Word(cpu.PC) + defs.Word(imm << 2)
		cpu.PC = uint32(newPC)
	}

	return nil
}


func (cpu *CPU) blezInstr(rs, rt uint8, imm int16) error {
	op1 := cpu.Registers[rs]

	if op1 <= 0 {
		newPC := defs.Word(cpu.PC) + defs.Word(imm << 2)
		cpu.PC = uint32(newPC)
	}

	return nil
}

func (cpu *CPU) bneInstr(rs, rt uint8, imm int16) error {
	op1 := cpu.Registers[rs]
	op2 := cpu.Registers[rt]
	if op1 != op2 {
		newPC := defs.Word(cpu.PC) + defs.Word(imm << 2)
		cpu.PC = uint32(newPC)
	}

	return nil
}

func (cpu *CPU) beqInstr(rs, rt uint8, imm int16) error {
	op1 := cpu.Registers[rs]
	op2 := cpu.Registers[rt]
	if op1 == op2 {
		newPC := defs.Word(cpu.PC) + defs.Word(imm << 2)
		cpu.PC = uint32(newPC)
	}

	return nil
}

func (cpu *CPU) lwInstr(rs, rt uint8, imm int16) error {
	baseAddr := cpu.Registers[rs]
	fullAddr := uint32(baseAddr + defs.Word(imm))
	memoryVal, err := cpu.MainMemory.LoadWord(fullAddr)
	if err != nil {
		return err
	}

	if rt == 0 {
		return errors.New("cannot write to $zero register")
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

func (cpu *CPU) andiInstr(rs, rt uint8, imm int16) error {
	op1 := cpu.Registers[rs]

	if rt == 0 {
		return errors.New("cannot write to $zero register")
	}
	cpu.Registers[rt] = op1 & defs.Word(imm)

	return nil
}

func (cpu *CPU) addiuInstr(rs, rt uint8, imm int16) error {
	op1 := uint32(cpu.Registers[rs])
	check := op1 + uint32(imm)

	if rt == 0 {
		return errors.New("cannot write to $zero register")
	}

	cpu.Registers[rt] = defs.Word(check)
	return nil
}

func (cpu *CPU) addiInstr(rs, rt uint8, imm int16) error {
	op1 := cpu.Registers[rs]
	check := op1 + defs.Word(imm)
	if isOverflow(op1, defs.Word(imm), check) {
		return errors.New("signed overflow exception")
	}

	if rt == 0 {
		return errors.New("cannot write to $zero register")
	}

	cpu.Registers[rt] = check
	return nil
}
