package main

import "fmt"

type GB struct {
	cpu CPU
	mem Memory
}

func main() {
	var gb GB
	gb.mem.loadROM("romfile/Tetris (World) (Rev A).gb")
	// boot-rom 0x0000 to 0x0099
	for {
		gb.cpu.step(&gb.mem)
	}

	fmt.Println("finish execution")
}
