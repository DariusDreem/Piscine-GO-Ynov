package piscine

import (
	"fmt"
)

func SplitWhiteSpaces(s string) []string {
	array := []string{}
	str := ""
	i := 0
	for _, word := range s {
		word = word
		i++
		if word == 32 {
			fmt.Println("nique")
			array = append(array, str)
			str = ""
		} else {
			str += string(word)
		}
	}
	array = append(array, str)
	return array
}
