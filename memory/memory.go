package memory

import (
	"errors"
	"mips_emulator/defs"
	"encoding/binary"
)

const (
	DATA_SIZE = 1073741824 // 1 GB (2^30)
	INSTR_SIZE = 1048576
)


/**
	this structure represents DRAM (i.e., slower, main memory)
	mutli-byte data will be stored as little-endian
*/
type MainMemory struct {
	Instruction *[INSTR_SIZE]byte
	Data *[DATA_SIZE]byte
}

func InitMemory() *MainMemory {
	memory := &MainMemory{}
	memory.Instruction = &[INSTR_SIZE]byte{}
	memory.Data = &[DATA_SIZE]byte{}

	return memory
}

func (m MainMemory) FetchInstruction(addr uint32) (defs.Word, error) {
	if int(addr) > INSTR_SIZE - 4 {
		return 0, errors.New("invalid address")
	}	

	if addr % 4 != 0 {
		return 0, errors.New("address must be word-aligned")
	}

	return defs.Word(binary.LittleEndian.Uint32(m.Instruction[addr:addr+4])), nil
}

func (m *MainMemory) LoadInstruction(addr uint32, instr defs.Word) error {
	if int(addr) > INSTR_SIZE - 4 {
		return errors.New("invalid address")
	}
	if addr % 4 != 0 {
		return errors.New("address must be word-aligned")
	}

	binary.LittleEndian.PutUint32(m.Instruction[addr:addr+4], uint32(instr))
	return nil
}

func (m *MainMemory) StoreByte(addr uint32, data byte) error {
	if int(addr) > INSTR_SIZE {
		return errors.New("invalid address")
	}
	m.Data[addr] = data
	return nil
}

func (m MainMemory) LoadWord(addr uint32) (defs.Word, error) {
	if int(addr) > DATA_SIZE - 4 {
		return 0, errors.New("invalid address")	
	}

	if int(addr) % 4 != 0 {
		return 0, errors.New("address must be word-aligned")
	}
	
	return defs.Word(binary.LittleEndian.Uint32(m.Data[addr:addr+4])), nil
}

func (m *MainMemory) StoreWord(addr uint32, val defs.Word) error {
	if int(addr) > DATA_SIZE - 4 {
		return errors.New("invalid address")
	}

	if int(addr) % 4 != 0 {
		return errors.New("address must be word-aligned")
	}

	binary.LittleEndian.PutUint32(m.Data[addr:addr+4], uint32(val))
	return nil
}
