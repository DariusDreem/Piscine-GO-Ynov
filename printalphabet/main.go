package main

import (
	"github.com/01-edu/z01"
	_ "github.com/01-edu/z01"
)

func main() {
	for i := 97; i < 122; {

		z01.PrintRune(rune(i))
		i++
	}
	z01.PrintRune('\n')
}
