package main

func (r *Register) getAFBE() uint16 {
	return join8to16(r.a, r.f)
}

func (r *Register) setAFBE(x uint16) {
	r.a, r.f = split16to8(x)
}

func (r *Register) getBCBE() uint16 {
	return join8to16(r.b, r.c)
}

func (r *Register) setBCBE(x uint16) {
	r.b, r.c = split16to8(x)
}

func (r *Register) getDEBE() uint16 {
	return join8to16(r.d, r.e)
}

func (r *Register) setDEBE(x uint16) {
	r.d, r.e = split16to8(x)
}

func (r *Register) getHLBE() uint16 {
	return join8to16(r.h, r.l)
}

func (r *Register) setHLBE(x uint16) {
	r.h, r.l = split16to8(x)
}
