package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x rune
	y rune
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func PrintStr(s string) {
	for _, a := range s {
		z01.PrintRune(a)
	}
}

func main() {
	points := &point{}
	setPoint(points)
	PrintStr("x = ")
	z01.PrintRune(points.x/10 + 48)
	z01.PrintRune(points.x%10 + 48)
	PrintStr(",")
	PrintStr(" ")
	PrintStr("y = ")
	z01.PrintRune(points.y/10 + 48)
	z01.PrintRune(points.y%10 + 48)
	z01.PrintRune('\n')
}
