package piscine

import "github.com/01-edu/z01"

func ToUpper(s string) string {
	world := ""
	for _, letter := range s {
		if letter >= 97 && letter <= 122 {
			letter = letter - 32
			z01.PrintRune(letter)
		} else {
			z01.PrintRune(letter)
		}
	}
	return world
}
