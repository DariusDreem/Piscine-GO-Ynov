package piscine

func NRune(s string, n int) rune {
	for index, letter := range s {
		if StrLen(s) < n || n < 0 {
			return 0
		} else if index+1 == n {
			return letter
		}
	}
	return 0
}
