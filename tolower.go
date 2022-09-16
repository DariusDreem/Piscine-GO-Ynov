package piscine

func ToLower(s string) string {
	world := ""
	for _, letter := range s {
		if letter >= 65 && letter <= 90 {
			letter = letter + 32
			world += string(letter)
		} else {
			world += string(letter)
		}
	}
	return world
}
