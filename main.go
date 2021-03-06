package main

import (
	"fmt"
)

func main() {
	cart := loadROM("romfile/Tetris (World) (Rev A).gb")
	fmt.Println(cart)
}
