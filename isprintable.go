package piscine

func IsPrintable(s string) bool {
	for _, letter := range s {
		if letter <= 31 && letter >= 0 {
			return false
		} else {
			continue
		}
	}
	return true
}
