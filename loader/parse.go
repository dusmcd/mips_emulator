package loader

import (
	"debug/elf"
	"fmt"
	"errors"
	"encoding/binary"
	"mips_emulator/memory"
	"mips_emulator/defs"
)

const (
	LITTLE_ENDIAN = 1
	BIG_ENDIAN = 2
	UNKNOWN_FORMAT = 0
	SIZE_32 = 1
	SIZE_64 = 2
	EXEC = 2
	MIPS = 8
	MIPS_LE = 10
)

type StartUpData struct {
	EntryAddr uint32
	Memory *memory.MainMemory
}

func ParseProgramHeaders(headers []*elf.Prog, memory *memory.MainMemory) error {
	for _, pHeader := range headers {
		switch(pHeader.Type) {
		case elf.PT_LOAD:
			// load into main memory
			addr := uint32(pHeader.Vaddr)
			size := pHeader.Filesz
			buffer := []byte{}
			bytesRead, err := pHeader.Open().Read(buffer)
			
			if err != nil {
				return err
			}
			if bytesRead != int(size) {
				return errors.New("bytes read does not match given size")
			}

			// assuming data is word aligned
			for i := 0; i < len(buffer); i += 4 {
				word := binary.LittleEndian.Uint32(buffer[i:i+4])
				memory.StoreWord(addr, defs.Word(word))	
			}
		}
	}	
	return nil
}

func ParseFile(path string) (*StartUpData, error) {
	elfFile, err := elf.Open(path)
	if err != nil {
		return nil, err
	}

	if elfFile.Data != LITTLE_ENDIAN {
		return nil, errors.New("data format not supported")
	}

	if elfFile.Class != SIZE_32 {
		return nil, errors.New("word length not 32 bits")
	}

	if elfFile.Type != EXEC {
		return nil, errors.New("must be an executable file")
	}

	if elfFile.Machine != MIPS && elfFile.Machine != MIPS_LE {
		return nil, errors.New("must be a MIPS machine")
	}
	startUpData := StartUpData{}
	memory := memory.InitMemory()
	err = ParseProgramHeaders(elfFile.Progs, memory)
	if err != nil {
		return nil, err
	}

	fmt.Println("ELF Metadata")
	fmt.Println("============")
	fmt.Printf("Class: %d\n", elfFile.Class)
	fmt.Printf("Data: %d\n", elfFile.Data)
	fmt.Printf("Machine: %d\n", elfFile.Machine)

	err = elfFile.Close()
	if err != nil {
		return nil, err
	}

	startUpData.Memory = memory
	startUpData.EntryAddr = uint32(elfFile.Entry)
	return &startUpData, nil
}
