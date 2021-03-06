package main

import (
	"io/ioutil"
)

type Memory struct {
	cart []uint8
}

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

func (m *Memory) readOperation(i uint16) uint8 {
	return m.cart[i]
}

func (m *Memory) readAddress(i uint16) uint16 {
	a := uint16(m.cart[i])
	b := uint16(m.cart[i + 1])
	b = b << 4
	return a + b
}
