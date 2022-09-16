package piscine

func IsLower(s string) bool {
	for _, letter := range s {
		if letter < 97 || letter > 122 {
			return false
		} else {
			continue
		}
	}
	return true
}
