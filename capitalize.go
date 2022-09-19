package piscine

func Capitalize(s string) string {
	world := ""
	pabo := 0
	for _, letter := range s {
		if letter >= 97 && letter <= 122 && pabo == 0 {
			letter = letter - 32
			pabo++
			world += string(letter)
		} else if letter >= 97 && letter <= 122 && pabo != 0 {
			pabo++
			world += string(letter)
		} else if letter >= 65 && letter <= 90 {
			if pabo != 0 {
				letter = letter + 32
				pabo++
				world += string(letter)
			} else {
				pabo++
				world += string(letter)
			}
		} else if letter >= 48 && letter <= 57 {
			world += string(letter)
			pabo++
		} else {
			pabo = 0
			world += string(letter)
		}
	}
	return world
}
