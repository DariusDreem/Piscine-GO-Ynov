package piscine

import (
	"github.com/01-edu/z01"
	_ "github.com/01-edu/z01"
)

func PrintNbr(nb int) {
	newSlice := []int{}

	if nb > 0 || nb < 0 {
		a := nb % 10
		newSlice = append(newSlice, a+48)
		z01.PrintRune(rune(a + 48))
		nb /= 10
		PrintNbr(nb)
	} else {
		z01.PrintRune('0')
	}
}
