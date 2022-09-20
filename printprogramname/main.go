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
		if 65 <= el && 90 >= el || 97 <= el && 122 >= el || el == 46 {
			resultList = append(resultList, el)
		} else {
			break
		}
	}
	for i := len(resultList) - 1; i != -1; i-- {
		z01.PrintRune(rune(resultList[i]))
	}
	return
}
