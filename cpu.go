package main

import "fmt"

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

func (r *Register) getSP() uint16 {
	return r.sp
}

type CPU struct {
	register Register
	flags uint8
	regList []*uint8
	doubleRegList1 []func() uint16
}

func (cpu *CPU) step(mem *Memory) {
	op := mem.readOperation(cpu.register.pc)
	cpu.register.pc++
	cpu.exec(op, mem)
}

func (cpu *CPU) exec(op uint8, mem *Memory) {
	fmt.Printf("%d %x %b\n", op, op, op)

	switch op {
	case 0x9, 0x19, 0x29, 0x39:
		p := op >> 4 & 3
		val := cpu.doubleRegList1[p]()
		result := cpu.register.getHL() + val

		cpu.register.setHL(result)

	case 0xc3: // jp nn
		nn := mem.readAddress(cpu.register.pc)
		cpu.register.pc = nn
	default:
		if op != 0 { // if not noop
			panic(op)
		}
	}



	//y := op >> 3 & 7
	//z := op & 7
	//dst := cpu.regList[int(y)]
	//src := cpu.regList[int(z)]
	//
	//switch x := op >> 6; x {
	//case 0:
	//	fmt.Println("0")
	//case 1:
	//	fmt.Println("1")
	//case 2:
	//	fmt.Println("2") // alu
	//	//dst = add(*dst, *src)
	//case 3: // jumps and loads
	//	fmt.Println("3")
	//	*dst = add(*dst, *src)
	//}
}

func jump(x uint8) {

}

func add(x, y uint8) uint8 {
	return x + y
}
