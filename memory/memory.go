package memory

import (
	"errors"
	"mips_emulator/defs"
	"encoding/binary"
)

const (
	DATA_SIZE = 1073741824
	INSTR_SIZE = 1048576
)


/**
	this structure represents DRAM (i.e., slower, main memory)
	mutli-byte data will be stored as little-endian
*/
type MainMemory struct {
	Instruction *[DATA_SIZE]byte
	Data *[INSTR_SIZE]byte
}

func InitMemory() *MainMemory {
	var instruction [DATA_SIZE]byte
	var data [INSTR_SIZE]byte
	memory := &MainMemory{}
	memory.Instruction = &instruction
	memory.Data = &data

	return memory
}

func getWord(addr uint32, mem []byte) defs.Word {
	uData := binary.LittleEndian.Uint32(mem[addr:addr+4])
	word := defs.Word(uData)
	return word
}

func storeWord(addr uint32, val defs.Word, mem []byte) {
	binary.LittleEndian.PutUint32(mem[addr:addr+4], uint32(val))
}

func (m MainMemory) FetchInstruction(addr uint32) (defs.Word, error) {
	if int(addr) > INSTR_SIZE - 4 {
		return 0, errors.New("invalid address")
	}	

	if addr % 4 != 0 {
		return 0, errors.New("address must be word-aligned")
	}
	data := m.Instruction[:]
	return getWord(addr, data), nil
}

func (m *MainMemory) LoadInstruction(addr uint32, instr defs.Word) error {
	if int(addr) > INSTR_SIZE - 4 {
		return errors.New("invalid address")
	}
	if addr % 4 != 0 {
		return errors.New("address must be word-aligned")
	}

	data := m.Instruction[:]
	storeWord(addr, instr, data)
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
	
	data := m.Data[:]
	return getWord(addr, data), nil
}

func (m *MainMemory) StoreWord(addr uint32, val defs.Word) error {
	if int(addr) > DATA_SIZE - 4 {
		return errors.New("invalid address")
	}

	if int(addr) % 4 != 0 {
		return errors.New("address must be word-aligned")
	}

	data := m.Data[:]
	storeWord(addr, val, data)
	return nil
}
