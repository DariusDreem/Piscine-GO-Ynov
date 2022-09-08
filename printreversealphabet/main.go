package main

import (
	"github.com/01-edu/z01"
	_ "github.com/01-edu/z01"
)

func main() {
	for i := 122; i >= 97; {

		z01.PrintRune(rune(i))
		i = i - 1
	}
	z01.PrintRune('\n')
}
