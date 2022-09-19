package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	isnegativeNbr := 0
	baseList := []string{}
	loop := 0
	resultBase := ""
	if nbr < 0 {
		isnegativeNbr = 1
		nbr = -nbr
	}
	if len(base) < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	for _, element := range base {
		loop = 0
		if len(baseList) == 0 {
			baseList = append(baseList, string(element))
			continue
		}
		for _, baseLetter := range baseList {
			if string(element) == baseLetter || element == 43 || element == 45 {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			} else {
				continue
			}
		}
		if loop == 0 {
			baseList = append(baseList, string(element))
		}
	}
	for nbr = nbr; nbr != 0; {
		resultmod := nbr % len(base)
		resultBase = string(base[resultmod]) + resultBase
		nbr /= len(base)
	}
	for i := 0; i < len(resultBase); i++ {
		if isnegativeNbr == 1 {
			z01.PrintRune('-')
			isnegativeNbr = 0
		}
		z01.PrintRune(rune(resultBase[i]))
	}
}
