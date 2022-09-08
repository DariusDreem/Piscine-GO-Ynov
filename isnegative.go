package piscine

import (
	"github.com/01-edu/z01"
	_ "github.com/01-edu/z01"
)

func IsNegative(nb int) {

	i := nb

	if i < 0 {
		z01.PrintRune('T')
		z01.PrintRune('\n')

	} else if i >= 0 {
		z01.PrintRune('F')
		z01.PrintRune('\n')

	}

}
