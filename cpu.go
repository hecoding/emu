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
	regList []*uint8
	doubleRegList1 []func() uint16
}

func (cpu *CPU) readInstruction(mem *Memory) uint8 {
	op := mem.readByte(cpu.register.pc)
	cpu.register.pc++
	return op
}

func (cpu *CPU) readOperand(mem *Memory) uint16 {
	n := mem.read2Bytes(cpu.register.pc)
	cpu.register.pc += 2
	return n
}

func (cpu *CPU) step(mem *Memory) {
	op := cpu.readInstruction(mem)
	cpu.exec(op, mem)
}

func (cpu *CPU) exec(op uint8, mem *Memory) {
	fmt.Printf("%d %x %b\n", op, op, op)

	switch op {
	case 0x9, 0x19, 0x29, 0x39: // add hl, n
		p := op >> 4 & 3
		register := cpu.register.getHL()
		val := cpu.doubleRegList1[p]()
		result := register + val

		cpu.register.clearFlag(negativeFlag)
		cpu.register.checkHalfCarryFlag(register, val)
		cpu.register.checkCarryFlag(result)

		cpu.register.setHL(result)

	case 0xc3: // jp nn
		nn := cpu.readOperand(mem)
		cpu.register.pc = nn
	default:
		if op != 0 { // if not noop
			panic(op)
		}
	}
}

func (r *Register) clearFlag(flag uint8) {
	r.f &= ^flag
}

func (r *Register) setFlag(flag uint8) {
	r.f |= flag
}

func (r *Register) checkHalfCarryFlag(register uint16, value uint16) {
	if isHalfCarry(register, value) {
		r.setFlag(halfCarryFlag)
	} else {
		r.clearFlag(halfCarryFlag)
	}
}

func (r *Register) checkCarryFlag(result uint16) {
	if result & 0xff00 != 0 {
		r.setFlag(carryFlag)
	} else {
		r.clearFlag(carryFlag)
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
