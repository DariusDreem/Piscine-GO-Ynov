package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var list []int
	length := 0
	if n == 0 {
		list = append(list, 0)
		z01.PrintRune(rune(list[0]) + '0')
	}
	for n > 0 {
		last := n % 10
		n = n / 10
		list = append(list, last)
		length++
	}
	for {
		counter := 0
		for i := 1; i < length; i++ {
			if list[i] < list[i-1] {
				list[i], list[i-1] = list[i-1], list[i]
				counter++
			}
		}
		if counter == 0 {
			break
		}
	}
	for j := 0; j < length; j++ {
		z01.PrintRune(rune(list[j]) + '0')
	}
}
