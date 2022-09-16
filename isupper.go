package piscine

func IsUpper(s string) bool {
	for _, letter := range s {
		if letter < 61 || letter > 90 {
			return false
		} else {
			continue
		}
	}
	return true
}
