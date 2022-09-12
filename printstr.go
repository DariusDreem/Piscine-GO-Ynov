package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {

	for _, world := range s {
		z01.PrintRune(world)
	}
}

