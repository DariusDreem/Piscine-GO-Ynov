package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func even(nbr2 int) bool {
	if nbr2%2 == 0 {
		return true
	} else {
		return false
	}
}

func isEven(nbr int) bool {
	if even(nbr) == true {
		return true
	} else {
		return false
	}
}

func main() {
	arg := os.Args[1:]
	length := len(arg)
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"
	if isEven(length) == true {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
