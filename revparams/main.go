package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	for i := len(arg) - 1; i != 0; i-- {
		remove_string(arg[i])
	}
}

func remove_string(str string) {
	for _, el := range str {
		z01.PrintRune(el)
	}
	z01.PrintRune('\n')
	return
}
