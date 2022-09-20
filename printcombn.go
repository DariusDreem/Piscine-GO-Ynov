package piscine

import (
	"github.com/01-edu/z01"
)

func PrintCombN(n int) {
	resultList := []int{}
	if n > 10 {
		return
	}
	for i := 0; i != n; i++ {
		resultList = append(resultList, i)
	}
	if n == 1 {
		for i := 0; i < 10; i++ {
			z01.PrintRune(rune(48 + i))
			if i != 9 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
		return
	}
	if resultList[n-n] == 0 {
		resultList[n-1] = int(resultList[n-1])
		for _, el := range resultList {
			z01.PrintRune(rune(el + 48))
		}
		z01.PrintRune(',')
		z01.PrintRune(' ')
	}
	for {
		if resultList[n-1] != 9 {
			resultList[n-1] = int(resultList[n-1]) + 1
		} else if n >= 2 && resultList[n-2] != 8 && resultList[n-1] == 9 {
			resultList[n-2]++
			resultList[n-1] = resultList[n-2] + 1
		} else if n >= 3 && resultList[n-3] != 7 && resultList[n-2] == 8 {
			resultList[n-3]++
			resultList[n-2] = resultList[n-3] + 1
			resultList[n-1] = resultList[n-2] + 1
		} else if n >= 4 && resultList[n-4] != 6 && resultList[n-3] == 7 {
			resultList[n-4]++
			resultList[n-3] = resultList[n-4] + 1
			resultList[n-2] = resultList[n-3] + 1
			resultList[n-1] = resultList[n-2] + 1
		} else if n >= 5 && resultList[n-5] != 5 && resultList[n-4] == 6 {
			resultList[n-5]++
			resultList[n-4] = resultList[n-5] + 1
			resultList[n-3] = resultList[n-4] + 1
			resultList[n-2] = resultList[n-3] + 1
			resultList[n-1] = resultList[n-2] + 1
		} else if n >= 6 && resultList[n-6] != 4 && resultList[n-5] == 5 {
			resultList[n-6]++
			resultList[n-5] = resultList[n-6] + 1
			resultList[n-4] = resultList[n-5] + 1
			resultList[n-3] = resultList[n-4] + 1
			resultList[n-2] = resultList[n-3] + 1
			resultList[n-1] = resultList[n-2] + 1
		} else if n >= 7 && resultList[n-7] != 3 && resultList[n-6] == 4 {
			resultList[n-7]++
			resultList[n-6] = resultList[n-7] + 1
			resultList[n-5] = resultList[n-6] + 1
			resultList[n-4] = resultList[n-5] + 1
			resultList[n-3] = resultList[n-4] + 1
			resultList[n-2] = resultList[n-3] + 1
			resultList[n-1] = resultList[n-2] + 1
		} else if n >= 8 && resultList[n-8] != 2 && resultList[n-7] == 3 {
			resultList[n-8]++
			resultList[n-7] = resultList[n-8] + 1
			resultList[n-6] = resultList[n-7] + 1
			resultList[n-5] = resultList[n-6] + 1
			resultList[n-4] = resultList[n-5] + 1
			resultList[n-3] = resultList[n-4] + 1
			resultList[n-2] = resultList[n-3] + 1
			resultList[n-1] = resultList[n-2] + 1
		} else {
			resultList[n-9]++
			resultList[n-8] = resultList[n-9] + 1
			resultList[n-7] = resultList[n-8] + 1
			resultList[n-6] = resultList[n-7] + 1
			resultList[n-5] = resultList[n-6] + 1
			resultList[n-4] = resultList[n-5] + 1
			resultList[n-3] = resultList[n-4] + 1
			resultList[n-2] = resultList[n-3] + 1
			resultList[n-1] = resultList[n-2] + 1
		}
		for _, el := range resultList {
			z01.PrintRune(rune(el + 48))
		}
		if resultList[n-n] == 10-n {
			z01.PrintRune('\n')
			break
		} else {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}
}
