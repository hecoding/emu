package main

import "fmt"

type CPU struct {
	register Register
	registerTable registerTables
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

func (cpu *CPU) readImmediate8(mem *Memory) uint8 {
	n := mem.readByte(cpu.register.pc)
	cpu.register.pc++
	return n
}

func (cpu *CPU) readStack16(mem *Memory) uint16 {
	n := mem.read2Bytes(cpu.register.pc)
	cpu.register.pc += 2
	return n
}

func (cpu *CPU) writeStack16(mem *Memory, n uint16) {
	cpu.register.sp -= 2
	mem.write2Bytes(cpu.register.sp, n)
}

func (cpu *CPU) step(mem *Memory) {
	op := cpu.readInstruction(mem)
	cpu.exec(op, mem)
}

func (cpu *CPU) getPField(op uint8) uint8 {
	return op >> 4 & 3
}

func (cpu *CPU) getYField(op uint8) uint8 {
	return op >> 3 & 7
}

func (cpu *CPU) getZField(op uint8) uint8 {
	return op & 7
}

func (cpu *CPU) exec(op uint8, mem *Memory) {
	fmt.Printf("%d %x %b\n", op, op, op)

	switch op {
	// 16-bit arithmetic
	case 0xe8: // add sp, n
		//d := cpu.readImmediate8(mem)
		//https://stackoverflow.com/questions/5159603/gbz80-how-does-ld-hl-spe-affect-h-and-c-flags
		fmt.Println("do this")
		panic(op)

		cpu.register.resetFlag(zeroFlag)
		cpu.register.resetFlag(negativeFlag)
		//the other flags

	case 0x03, 0x13, 0x23, 0x33:  // inc nn
		p := cpu.getPField(op)
		cpu.registerTable.rpSetBE[p](cpu.registerTable.rpGetBE[p]() + 1)


	case 0x0B, 0x1B, 0x2B, 0x3B: // dec nn
		p := cpu.getPField(op)
		cpu.registerTable.rpSet[p](cpu.registerTable.rpGet[p]() - 1)


	case 0x06, 0x0e, 0x16, 0x1e, 0x26, 0x2e: // ld reg, n
		y := cpu.getYField(op)
		n := cpu.readImmediate8(mem)
		regY := cpu.registerTable.r[y]//check if 0x06 lands on b, etc

		*regY = n

	// 8-bit loads
	case 0x40, 0x41, 0x42, 0x43, 0x44, 0x45, 0x47, // ld reg,reg (non-double, i.e HL)
	0x48, 0x49, 0x4a, 0x4b, 0x4c, 0x4d, 0x4f,
	0x50, 0x51, 0x52, 0x53, 0x54, 0x55, 0x57,
	0x58, 0x59, 0x5a, 0x5b, 0x5c, 0x5d, 0x5f,
	0x60, 0x61, 0x62, 0x63, 0x64, 0x65, 0x67,
	0x68, 0x69, 0x6a, 0x6b, 0x6c, 0x6d, 0x6f,
	0x78, 0x79, 0x7a, 0x7b, 0x7c, 0x7d, 0x7f:
		y := cpu.getYField(op)
		z := cpu.getZField(op)
		regY := cpu.registerTable.r[y]
		regZ := cpu.registerTable.r[z]

		*regY = *regZ

	case 0x46, 0x4e, 0x56, 0x5e, 0x66, 0x6e, 0x7e: // ld reg, (HL) (same as before, just doubles)
		y := cpu.getYField(op)
		regY := cpu.registerTable.r[y]
		valZ := mem.readByte(cpu.register.getHL())

		*regY = valZ

	case 0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x36: // ld HL, reg
		y := cpu.getYField(op)
		z := cpu.getZField(op)
		addrY := cpu.registerTable.rpGet[y]()
		valZ := cpu.registerTable.r[z]

		fmt.Println(addrY, valZ)
		panic(op)
		//*addrY = valZ

	case 0xe0: // ldh (n), a
		n := uint16(cpu.readImmediate8(mem))
		mem.writeByte(0xff00 + n, cpu.register.a)

	case 0xf0: // ldh a, (n)
		n := uint16(cpu.readImmediate8(mem))
		cpu.register.a = mem.readByte(0xff00 + n)

	// 16-bit loads
	case 0xf5, 0xc5, 0xd5, 0xe5: // push nn
		p := cpu.getPField(op)
		cpu.writeStack16(mem, cpu.registerTable.rp2Get[p]())

	case 0xf1, 0xc1, 0xd1, 0xe1: // pop nn
		p := cpu.getPField(op)
		nn := cpu.readStack16(mem)
		cpu.registerTable.rp2Set[p](nn)

	case 0x9, 0x19, 0x29, 0x39: // add hl, n
		p := cpu.getPField(op)
		register := cpu.register.getHL()
		val := cpu.registerTable.rpGet[p]()
		result := register + val

		cpu.register.resetFlag(negativeFlag)
		cpu.register.checkHalfCarryFlag(register, val)
		cpu.register.checkCarryFlag(result)

		cpu.register.setHL(result)

	// jumps
	case 0xc3: // jp nn
		nn := cpu.readOperand(mem)
		cpu.register.pc = nn

	case 0xe9: // jp (hl)
		cpu.register.pc = cpu.register.getHLBE()

	// 8-bit alu
	case 0xaf, 0xa8, 0xa9, 0xaa, 0xab, 0xac, 0xad: // xor reg
		z := cpu.getZField(op)
		regZ := cpu.registerTable.r[z]
		result := cpu.register.a ^ *regZ

		cpu.register.checkZeroFlag(result)
		cpu.register.resetFlag(negativeFlag)
		cpu.register.resetFlag(halfCarryFlag)
		cpu.register.resetFlag(carryFlag)

		cpu.register.a = result

	case 0xae: // xor (hl)
		panic(op)
		//flags

	case 0xee: // xor n
		//n := cpu.readImmediate8(mem)
		panic(op)
		//flags

	default:
		if op != 0 { // if not noop
			panic(op)
		}
	}
}
