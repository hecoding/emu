package main

import (
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func loadROM(filepath string) []byte {
	data, err := ioutil.ReadFile(filepath)
	check(err)

	// some rom checks
	return data
}
