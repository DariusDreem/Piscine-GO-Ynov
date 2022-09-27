package main

import (
	"piscine"

	"github.com/01-edu/z01"
)

type point struct {
	y int
	x int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := point{}

	setPoint(&points)

	piscine.PrintStr("x =")
	z01.PrintRune(rune(points.x/10 + 48))
	z01.PrintRune(rune(points.x%10 + 48))
	piscine.PrintStr(", ")
	piscine.PrintStr("y =")
	z01.PrintRune(rune(points.y/10 + 48))
	z01.PrintRune(rune(points.y%10 + 48))
}
