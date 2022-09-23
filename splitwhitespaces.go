package piscine

func SplitWhiteSpaces(s string) []string {
	array := []string{}
	str := ""
	i := 0
	yee := false
	for _, word := range s {
		word = word
		i++
		if word == 32 {
			array = append(array, str)
			str = ""
			yee = true
		} else if word == 32 && yee {
			yee = false
			continue
		} else {
			str += string(word)
		}
	}
	array = append(array, str)
	return array
}
