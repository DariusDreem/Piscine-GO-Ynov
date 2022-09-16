package piscine

func ToUpper(s string) string {
	world := ""
	for _, letter := range s {
		if letter >= 97 && letter <= 122 {
			letter = letter - 32
			world += string(letter)
		} else {
			world += string(letter)
		}
	}
	return world
}
