package piscine

func IsAlpha(s string) bool {
	for _, letter := range s {
		if letter >= 97 && letter <= 122 || letter >= 65 && letter <= 90 || letter >= 48 && letter <= 57 {
			continue
		} else {
			return false
		}
	}
	return true
}
