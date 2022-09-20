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
	resultList := []int{}
	for i := len(str) - 1; i != -1; i-- {
		el := int(str[i])
		if el == 92 || el == 47 || el == 58 || el == 42 || el == 63 || el == 34 || el == 60 || el == 62 || el == 124 {
			break
		} else {
			resultList = append(resultList, el)
		}
	}
	for i := len(resultList) - 1; i != -1; i-- {
		z01.PrintRune(rune(resultList[i]))
	}
	return
}
