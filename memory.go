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
