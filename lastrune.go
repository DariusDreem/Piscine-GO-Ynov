package piscine

func LastRune(s string) rune {
	for index, letter := range s {
		if StrLen(s)-1 == index {
			return letter
		}
	}
	return 0
}
