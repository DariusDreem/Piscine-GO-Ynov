package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	remove_string(arg[0])
}

func remove_string(str string) {
	result := 0
	for i, el := range str {
		if el == 47 || el == 92 {
			result = i
		}
	}
	str = str[result+1:]
	for _, el := range str {
		z01.PrintRune(el)
	}
	z01.PrintRune('\n')
	return
}
