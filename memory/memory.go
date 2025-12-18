package memory

import (
	"errors"
	"mips_emulator/defs"
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

func getWord[T *[DATA_SIZE]byte | *[INSTR_SIZE]byte](addr uint32, mem T) defs.Word {
	var word defs.Word = 0
	word = word | defs.Word(mem[addr]) // least significant byte
	word = word | (defs.Word(mem[addr + 1]) << 8)
	word = word | (defs.Word(mem[addr + 2]) << 16)
	word = word | (defs.Word(mem[addr + 3]) << 24) // most significant byte
	
	return word
}

func storeWord[T *[DATA_SIZE]byte | *[INSTR_SIZE]byte](addr uint32, val defs.Word, mem T) {
	mem[addr] = byte(val) // least significant byte
	mem[addr + 1] = byte((val & 0x0000FF00) >> 8)
	mem[addr + 2] = byte((val & 0x00FF0000) >> 16)
	mem[addr + 3] = byte((uint(val) & uint(0xFF000000)) >> 24) // most significant byte
}

func (m MainMemory) FetchInstruction(addr uint32) (defs.Word, error) {
	if int(addr) > INSTR_SIZE - 4 {
		return 0, errors.New("invalid address")
	}	
	return getWord(addr, m.Instruction), nil
}

func (m *MainMemory) LoadInstruction(addr uint32, instr defs.Word) error {
	if int(addr) > INSTR_SIZE - 4 {
		return errors.New("invalid address")
	}
	storeWord(addr, instr, m.Instruction)
	return nil
}

func (m MainMemory) LoadWord(addr uint32) (defs.Word, error) {
	if int(addr) > DATA_SIZE - 4 {
		return 0, errors.New("invalid address")	
	}
	
	return getWord(addr, m.Data), nil
}

func (m *MainMemory) StoreWord(addr uint32, val defs.Word) error {
	if int(addr) > DATA_SIZE - 4 {
		return errors.New("invalid address")
	}

	storeWord(addr, val, m.Data)
	return nil
}
