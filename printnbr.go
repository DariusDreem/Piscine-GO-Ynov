package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	resultMod := 0
	oldResultList := []int{}
	isNegative := false
	if n == 0 {
		z01.PrintRune('0')
	}
	if n < 0 {
		isNegative = true
		z01.PrintRune('-')
	}
	for i := 0; n != 0; i++ {
		resultMod = n % 10
		if isNegative == true {
			resultMod *= -1
		}
		oldResultList = append(oldResultList, resultMod)
		n /= 10
	}
	for i := len(oldResultList) - 1; i != -1; i-- {
		z01.PrintRune(rune(oldResultList[i] + 48))
	}
}
