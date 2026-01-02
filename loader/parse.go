package loader

import (
	"debug/elf"
	"fmt"
	"errors"
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

func ParseFile(path string) error {
	elfFile, err := elf.Open(path)
	if err != nil {
		return err
	}

	if elfFile.Data != LITTLE_ENDIAN {
		return errors.New("data format not supported")
	}

	if elfFile.Class != SIZE_32 {
		return errors.New("word length not 32 bits")
	}

	if elfFile.Type != EXEC {
		return errors.New("must be an executable file")
	}

	if elfFile.Machine != MIPS && elfFile.Machine != MIPS_LE {
		return errors.New("must be a MIPS machine")
	}

	fmt.Println("ELF Metadata")
	fmt.Println("============")
	fmt.Printf("Class: %d\n", elfFile.Class)
	fmt.Printf("Data: %d\n", elfFile.Data)
	fmt.Printf("Machine: %d\n", elfFile.Machine)

	err = elfFile.Close()
	if err != nil {
		return err
	}

	return nil
}
