package main

type Register struct {
	a, f, b, c, d, e, h, l uint8
	sp uint16
	pc uint16
}

// to simplify code https://gb-archive.github.io/salvage/decoding_gbz80_opcodes/Decoding%20Gamboy%20Z80%20Opcodes.html
type registerTables struct {
	r []*uint8
	rpGet []func() uint16
	rpSet []func(uint16)
	rpGetBE []func() uint16
	rpSetBE []func(uint16)
	rp2Get []func() uint16
	rp2Set []func(uint16)
	cc []func() uint16
	alu []func() uint16
	rot []func() uint16
}

func join8to16(a, b uint8) uint16 {
	c := uint16(a)
	d := uint16(b)
	c = c<<8
	return c + d
}

func split16to8(x uint16) (uint8, uint8) {
	a := x & 0xFF00
	b := x & 0x00FF
	a = a>>8
	return uint8(a), uint8(b)
}

func (r *Register) getAF() uint16 {
	return join8to16(r.f, r.a) // little endian order
}

func (r *Register) setAF(x uint16) {
	a, b := split16to8(x)
	r.f = a // little endian order
	r.a = b
}

func (r *Register) getBC() uint16 {
	return join8to16(r.c, r.b)
}

func (r *Register) setBC(x uint16) {
	a, b := split16to8(x)
	r.c = a
	r.b = b
}

func (r *Register) getDE() uint16 {
	return join8to16(r.e, r.d)
}

func (r *Register) setDE(x uint16) {
	a, b := split16to8(x)
	r.e = a
	r.d = b
}

func (r *Register) getHL() uint16 {
	return join8to16(r.l, r.h)
}

func (r *Register) setHL(x uint16) {
	a, b := split16to8(x)
	r.l = a
	r.h = b
}

func (r *Register) getSP() uint16 {
	return r.sp
}

func (r *Register) setSP(x uint16) {
	r.sp = x
}

func (r *Register) modifyFlags(z, n, h, c uint8) {
	switch z {
	case set:
		r.f &= 0
	case reset:
		r.f &= 0
	}
	switch n {
	case set:
		r.f &= 0
	case reset:
		r.f &= 0
	}
}

type flagChangeEnum int
const (
	same = iota
	set
	reset
	checkCarry
)

func (r *Register) resetFlag(flag uint8) {
	r.f &= ^flag
}

func (r *Register) setFlag(flag uint8) {
	r.f |= flag
}

func (r *Register) checkZeroFlag(result uint8) {
	if result == 0 {
		r.setFlag(zeroFlag)
	}
}

func (r *Register) checkHalfCarryFlag(register uint16, value uint16) {
	if isHalfCarry(register, value) {
		r.setFlag(halfCarryFlag)
	} else {
		r.resetFlag(halfCarryFlag)
	}
}

func (r *Register) checkCarryFlag(result uint16) {
	if result & 0xff00 != 0 {
		r.setFlag(carryFlag)
	} else {
		r.resetFlag(carryFlag)
	}
}

const (
	zeroFlag = 1 << 7
	negativeFlag = 1 << 6
	halfCarryFlag = 1 << 5
	carryFlag = 1 << 4
)

func isHalfCarry(a, b uint16) bool {
	return ((a & 0x0f) + (b & 0x0f)) > 0x0f
}

//func flagIsZero(r *Register) uint8 {
//	return r.flags & zeroFlag
//}
//
//func flagIsNegative(r *Register) uint8 {
//	return r.flags & negativeFlag
//}
//
//func flagIsHalfCarry(r *Register) uint8 {
//	return r.flags & halfCarryFlag
//}
//
//func flagIsCarry(r *Register) uint8 {
//	return r.flags & carryFlag
//}
//
//func flagIsSet(r *Register, x uint8) uint8 {
//	return r.flags & x
//}
//
//func flagSet(r *Register, x uint8) {
//	r.flags |= x
//}
//
//func flagClear(r *Register, x uint8) {
//	r.flags &= ^x
//}
