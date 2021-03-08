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

	gb.cpu.registerTable.r = []*uint8{&gb.cpu.register.b, &gb.cpu.register.c, &gb.cpu.register.d, &gb.cpu.register.e,
		&gb.cpu.register.h, &gb.cpu.register.l, &gb.cpu.register.f, &gb.cpu.register.a} // f should be HL
	gb.cpu.registerTable.rpGet = []func() uint16{gb.cpu.register.getBC, gb.cpu.register.getDE, gb.cpu.register.getHL,
		gb.cpu.register.getSP}
	gb.cpu.registerTable.rpSet = []func(x uint16){gb.cpu.register.setBC, gb.cpu.register.setDE, gb.cpu.register.setHL,
		gb.cpu.register.setSP}
	gb.cpu.registerTable.rpGetBE = []func() uint16{gb.cpu.register.getBCBE, gb.cpu.register.getDEBE, gb.cpu.register.getHLBE,
		gb.cpu.register.getSP}
	gb.cpu.registerTable.rpSetBE = []func(x uint16){gb.cpu.register.setBCBE, gb.cpu.register.setDEBE, gb.cpu.register.setHLBE,
		gb.cpu.register.setSP}
	gb.cpu.registerTable.rp2Get = []func() uint16{gb.cpu.register.getBC, gb.cpu.register.getDE, gb.cpu.register.getHL,
		gb.cpu.register.getAF}
	gb.cpu.registerTable.rp2Set = []func(x uint16) {gb.cpu.register.setBC, gb.cpu.register.setDE, gb.cpu.register.setHL,
		gb.cpu.register.setAF}
	gb.cpu.register.sp = 0xfffe

	for {
		gb.cpu.step(&gb.mem)
	}

	fmt.Println("finish execution")
}
