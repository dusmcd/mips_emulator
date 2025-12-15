package cpu

type RegFile [32]int32

type HiLowRegs struct {
	hi int32
	lo int32
}
