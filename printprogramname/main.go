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
		if el == 92 {
			result = i
		}
	}
	str = str[result+1:]
	for i := 0; i != len(str); i++ {
		z01.PrintRune(rune(str[i]))
	}
	return
}
