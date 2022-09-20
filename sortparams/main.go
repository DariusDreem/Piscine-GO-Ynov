package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	resultList := []string{}
	for i := len(arg) - 1; i != 0; i-- {
		resultList = append(resultList, arg[i])
	}
	resultList = sortParams(resultList)
	for _, el := range resultList {
		z01.PrintRune(rune(el[0]))
		z01.PrintRune('\n')
	}
}

func sortParams(array []string) []string {
	for i := 0; i < len(array)-1; i++ {
		if array[i] > array[i+1] {
			array[i], array[i+1] = array[i+1], array[i]
			i = -1
		}
	}
	return array
}
