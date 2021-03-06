package main

type Register struct {
	a, f, b, c, d, e, h, l uint8
	sp, pc uint16
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
