package piscine

func AlphaCount(s string) int {
	count := 0
	for index, letter := range s {
		if letter >= 65 && letter <= 90 {
			count++
		} else if letter >= 97 && letter <= 122 {
			count++
		} else if index == StrLen(s) {
			return count
		}
	}
	return count
}
