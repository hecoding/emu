package main

import (
	"io/ioutil"
)

type Memory struct {
	cart []uint8
	stack [0x80]uint8 // stack in GMB Z80 is a part of the regular memory
}

const stackOffset uint16 = 0xff80

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (m *Memory) loadROM(filepath string) {
	data, err := ioutil.ReadFile(filepath)
	check(err)

	// some rom checks
	m.cart = data
}

func (m *Memory) readByte(i uint16) uint8 {
	if i < 0x8000 {
		return m.cart[i]
	} else if i >= 0xff80 && i < 0xffff {
		return m.stack[i - stackOffset]
	} else {
		panic(i)
	}
}

func (m *Memory) writeByte(i uint16, n uint8) {
	if i < 0x8000 {
		m.cart[i] = n
	} else if i >= 0xff80 && i < 0xffff {
		m.stack[i - stackOffset] = n
	} else {
		panic(i)
	}
}

func (m *Memory) read2Bytes(i uint16) uint16 {
	var a, b uint8
	if i < 0x8000 {
		a = m.cart[i]
		b = m.cart[i + 1]
	} else if i >= 0xff80 && i < 0xffff {
		a = m.cart[i - stackOffset]
		b = m.cart[i + 1 - stackOffset]
	} else {
		panic(i)
	}

	return join8to16(b, a)
}

func (m *Memory) write2Bytes(i uint16, n uint16) {
	b, a := split16to8(n)

	if i < 0x8000 {
		m.cart[i] = a
		m.cart[i] = b
	} else if i >= 0xff80 && i < 0xffff {
		m.stack[i - stackOffset] = a
		m.stack[i - stackOffset + 1] = b
	} else {
		panic(i)
	}
}
