package cpu

import (
	"fmt"
	"errors"
	"mips_emulator/memory"
	"mips_emulator/defs"
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
	WORD_BYTES = 4
)

/**
	Globals
*/
var opMap map[uint8]IInstr
var funcMap map[uint8]RFunc 
var funcMapC map[uint8]RFunc

/**
	32-bit architecture using the MIPS ISA
	R-type:
		Op: 6 bits | RS: 5 bits | RT: 5 bits | RD: 5 bits | Shift: 5 bits | Func: 6 bits
	I-type:
		Op: 6 bits | RS: 5 bits | RT: bits | Immediate: 16 bits
	J-type:
		Op: 6 bits | Address: 26 bits
*/

type CPU struct {
	PC uint32 // addr of next instruction
	Registers RegFile
	HiLow HiLowRegs
	Instruction uint32 // encoded instruction
	MainMemory *memory.MainMemory
	Exit bool
}

func InitCPU(mem *memory.MainMemory, gp uint32) *CPU {
	cpu := &CPU{}
	cpu.Exit = false
	cpu.HiLow = HiLowRegs{}
	var regFile RegFile
	cpu.Registers = regFile
	
	// load GP (global pointer)
	cpu.Registers[28] = defs.Word(gp)
	// set stack pointer to highest word-aligned address
	cpu.Registers[29] = memory.DATA_SIZE - 4

	cpu.MainMemory = mem

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
		0x00: cpu.sllInstr,
		0x02: cpu.srlInstr,
		0x03: cpu.sraInstr,
		0x0C: cpu.syscall,
		0x2A: cpu.sltInstr, 
		0x2B: cpu.sltuInstr, 
	}

	funcMapC = map[uint8]RFunc{
		0x21: cpu.cloInstr,
	}

	opMap = map[uint8]IInstr{
		0x01: cpu.regImm,
		0x23: cpu.lwInstr,
		0x2B: cpu.swInstr,
		0x08: cpu.addiInstr,
		0x09: cpu.addiuInstr, 
		0x04: cpu.beqInstr,
		0x05: cpu.bneInstr,
		0x06: cpu.blezInstr,
		0x07: cpu.bgtzInstr,
		0x0C: cpu.andiInstr, 
		0x0D: cpu.oriInstr, 
		0x0E: cpu.xoriInstr,
	}
	

	return cpu
}

func (cpu *CPU) Run(initialAddr uint32) error {
	// load instruction from memory using PC address
	cpu.PC = initialAddr
	for {
		instruction, err := cpu.MainMemory.FetchInstruction(cpu.PC)
		if err != nil {
			log.Fatal(err.Error())
			return err
		}

		cpu.Instruction = uint32(instruction)
		cpu.PC += WORD_BYTES 
		err = cpu.DecodeInstr()
		if err != nil {
			log.Fatal(err.Error())
			return err
		}
		fmt.Printf("Instruction completed: 0x%X\n", uint32(instruction))
		if cpu.Exit {
			break
		}
	}

	return nil
}


func (cpu *CPU) decodeRType(op uint8) error {
	funcCode := uint8(cpu.Instruction & 0x0000003F)
	// get registers
	rs := uint8(cpu.Instruction & 0x03E00000 >> (WORD_BITS - OP_BITS - RS_BITS))
	rt := uint8(cpu.Instruction & 0x001F0000 >> (WORD_BITS - OP_BITS - RS_BITS - RT_BITS))
	rd := uint8(cpu.Instruction & 0x0000F800 >> (WORD_BITS - OP_BITS - RS_BITS - RT_BITS - RD_BITS))
	shift := uint8(cpu.Instruction & 0x000007C0 >> (WORD_BITS - OP_BITS - RS_BITS - RT_BITS - RD_BITS - SHIFT_BITS))

	switch(op) {
	case 0x0:
		if value, ok := funcMap[funcCode]; ok {
			return value(rs, rt, rd, shift)
		}
		err := fmt.Sprintf("func code not found: 0x%X", funcCode)
		return errors.New(err)
	case 0x1C:
		if value, ok := funcMapC[funcCode]; ok {
			return value(rs, rt, rd, shift)
		}
		err := fmt.Sprintf("func code not found: 0x%X", funcCode)
		return errors.New(err)
	}

	return errors.New("invalid machine code")
}

func (cpu *CPU) decodeIType(op uint8) error {
	rs := uint8(cpu.Instruction & 0x03E00000 >> (WORD_BITS - OP_BITS - RS_BITS))
	rt := uint8(cpu.Instruction & 0x001F0000 >> (WORD_BITS - OP_BITS - RS_BITS - RT_BITS))
	imm := int16(cpu.Instruction & 0x0000FFFF)

	if value, ok := opMap[op]; ok {
		return value(rs, rt, imm)
	}

	err := fmt.Sprintf("op code not found: 0x%X", op)

	return errors.New(err)
}

func (cpu *CPU) decodeJType(op uint8) error {
	addr := uint32(cpu.Instruction & 0x03FFFFFF)
	switch op {
	case 0x02:
		return cpu.jInstr(addr)
	case 0x03:
		return cpu.jalInstr(addr)
	}
	return errors.New("Invalid machine code")
}


func (cpu *CPU) DecodeInstr() error {
	if cpu.Instruction == 0 {
		return nil // nop
	}

	var instrType InstrType
	op := uint8(cpu.Instruction & 0xFC000000 >> (WORD_BITS - OP_BITS))
	if op == 0 || op == 0x1C {
		instrType = RType
	} else if op == 0x02 || op == 0x03 {
		instrType = JType
	} else {
		instrType = IType
	}

	switch instrType {
	case RType:
		// execute r-type instruction
		return cpu.decodeRType(op)
	case IType:
		// execute i-type instruction
		return cpu.decodeIType(op);
	case JType:
		// execute j-type instruction
		return cpu.decodeJType(op)
	}
	return errors.New("Invalid machine code")
}
