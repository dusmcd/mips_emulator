package cpu

import (
	"fmt"
	"errors"
	"mips_emulator/memory"
	"log"
)

type InstrType int

const (
	RType InstrType = iota
	IType
	JType
)

const (
	WORD_BITS = 32
	OP_BITS = 6
	RS_BITS = 5
	RT_BITS = 5
	RD_BITS = 5
	SHIFT_BITS = 5
	FUNC_BITS = 6
	IMM_BITS = 16
)

var opMap map[uint8]IInstr




	
/**
	32-bit architecture using the MIPS ISA
	R-type:
		Op: 6 bits | RS: 5 bits | RT: 5 bits | RD: 5 bits | Shift: 5 bits | Func: 6 bits
	I-type:
		Op: 6 bits | RS: 5 bits | RT: bits | Immediate: 16 bits
*/

type CPU struct {
	PC uint32 // addr of next instruction
	Registers *RegFile
	HiLow *HiLowRegs
	Instruction uint32 // encoded instruction
	MainMemory *memory.MainMemory
}

func InitCPU() CPU {
	cpu := CPU{}
	cpu.HiLow = &HiLowRegs{}
	var regFile RegFile
	cpu.Registers = &regFile
	cpu.MainMemory = memory.InitMemory()

	funcMap = map[uint8]RFunc{
	0x20: cpu.addInstr,
	0x22: cpu.subInstr,
	0x21: cpu.addUInstr,
	0x23: cpu.subUInstr,
	0x18: cpu.multInstr,
	0x19: cpu.multUInstr,
	0x1A: cpu.divInstr,
	0x1B: cpu.divUInstr,
	0x24: cpu.andInstr,
	0x25: cpu.orInstr,
	0x26: cpu.xorInstr,
	0x27: cpu.norInstr,
	}

	opMap = map[uint8]IInstr{
		0x23: cpu.lwInstr,
		0x2B: cpu.swInstr,
		0x08: cpu.addiInstr,
	}


	return cpu
}

func (cpu *CPU) Run(numInstructions int, initialAddr uint32) error {
	// load instruction from memory using PC address
	cpu.PC = initialAddr
	for range numInstructions {
		instruction, err := cpu.MainMemory.FetchInstruction(cpu.PC)
		if err != nil {
			log.Fatal(err.Error())
			return err
		}

		cpu.Instruction = uint32(instruction)
		cpu.PC += 4
		err  = cpu.DecodeInstr()
		if err != nil {
			log.Fatal(err.Error())
			return err
		}
		fmt.Printf("Instruction completed: 0x%X\n", instruction)
	}

	return nil
}


func (cpu *CPU) decodeRType()  error {
	funcCode := uint8(cpu.Instruction & 0x0000003F)
	// get registers
	rs := uint8(cpu.Instruction & 0x03E00000 >> (WORD_BITS - OP_BITS - RS_BITS))
	rt := uint8(cpu.Instruction & 0x001F0000 >> (WORD_BITS - OP_BITS - RS_BITS - RT_BITS))
	rd := uint8(cpu.Instruction & 0x0000F800 >> (WORD_BITS - OP_BITS - RS_BITS - RT_BITS - RD_BITS))
	shift := uint8(cpu.Instruction & 0x000007C0 >> (WORD_BITS - OP_BITS - RS_BITS - RT_BITS - RD_BITS - SHIFT_BITS))

	// execute operation
	funcMap[funcCode](rs, rt, rd, shift)	
	return nil
}

func (cpu *CPU) decodeIType(op uint8) error {
	rs := uint8(cpu.Instruction & 0x03E00000 >> (WORD_BITS - OP_BITS - RS_BITS))
	rt := uint8(cpu.Instruction & 0x001F0000 >> (WORD_BITS - OP_BITS - RS_BITS - RT_BITS))
	imm := int16(cpu.Instruction & 0x0000FFFF)

	opMap[op](rs, rt, imm)
	return nil

}


func (cpu *CPU) DecodeInstr() error {
	// need to look up op code in static memory
	var instrType InstrType
	op := uint8(cpu.Instruction & 0xFC000000 >> (WORD_BITS - OP_BITS))
	if op == 0 {
		instrType = RType
	} else if op == 0x02 || op == 0x03 {
		instrType = JType
	} else {
		instrType = IType
	}

	switch instrType {
	case RType:
		// execute r-type instruction
		return cpu.decodeRType()
	case IType:
		// execute i-type instruction
		return cpu.decodeIType(op);
	case JType:
		// execute j-type instruction
		break
	}
	return errors.New("Invalid machine code")
}
