package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	str := ""
	strList := []int32{}
	indexList := []int{}
	if len(arg) == 1 {
		z01.PrintRune('\n')
		return
	}
	for i := 1; i != len(arg); i++ {
		if i == 1 {
			str = arg[i]
		} else {
			str = str + " " + arg[i]
		}
	}
	for _, letter := range str {
		strList = append(strList, letter)
	}
	for i, letter := range strList {
		if letter == 65 || letter == 69 || letter == 73 || letter == 79 || letter == 85 || letter == 97 || letter == 101 || letter == 105 || letter == 111 || letter == 117 {
			indexList = append(indexList, i)
		}
	}
	for i := 0; i != len(indexList)/2; i++ {
		strList[indexList[i]], strList[indexList[len(indexList)-i-1]] = strList[indexList[len(indexList)-i-1]], strList[indexList[i]]
	}
	for _, lettre := range strList {
		z01.PrintRune(rune(lettre))
	}
	z01.PrintRune(' ')
	z01.PrintRune('\n')
}
