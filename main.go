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

	gb.cpu.regList = []*uint8{&gb.cpu.register.b, &gb.cpu.register.c, &gb.cpu.register.d, &gb.cpu.register.e,
		&gb.cpu.register.h, &gb.cpu.register.l, &gb.cpu.register.f, &gb.cpu.register.a}
	gb.cpu.doubleRegList1 = []func() uint16{gb.cpu.register.getBC, gb.cpu.register.getDE, gb.cpu.register.getHL,
		gb.cpu.register.getSP}

	for {
		gb.cpu.step(&gb.mem)
	}

	fmt.Println("finish execution")
}
