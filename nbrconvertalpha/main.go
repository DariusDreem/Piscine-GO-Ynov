package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	resultList := []int{}
	if len(arg) == 1 {
		return
	}
	for i := 1; i < len(arg); i++ {
		a := BasicAtoi(arg[i])
		if a >= 1 && a <= 26 {
			resultList = append(resultList, a)
		} else {
			resultList = append(resultList, 32)
		}
	}
	if arg[1] == "--upper" {
		resultList = resultList[1:]
		nbrconvertalphaUpper(resultList)
	} else {
		nbrconvertalphaLower(resultList)
	}
	z01.PrintRune('\n')
}

func nbrconvertalphaUpper(array []int) {
	for _, el := range array {
		if el != 32 {
			el += 64
		}
		z01.PrintRune(rune(el))
	}
}

func nbrconvertalphaLower(array []int) {
	for _, el := range array {
		if el != 32 {
			el += 96
		}
		z01.PrintRune(rune(el))
	}
}

func BasicAtoi(s string) int {
	rev := ""
	for i := len([]rune(s)) - 1; i >= 0; i-- {
		rev = rev + string(s[i])
	}
	a := 0
	i := 1
	for _, element := range rev {
		a = a + int(element-48)*i
		i = i * 10
	}
	return a
}
