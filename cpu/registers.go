package cpu

import (
	"mips_emulator/defs"
)

type RegFile [32]defs.Word

type HiLowRegs struct {
	hi defs.Word
	lo defs.Word
}
