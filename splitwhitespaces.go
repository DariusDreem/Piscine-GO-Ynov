package piscine

func SplitWhiteSpaces(s string) []string {
	str := ""
	i := 0
	array := []string{}
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
		} else if str == "" {
			str += string(word)
			yee = false
		} else {
			str += string(word)
			yee = false
		}
	}
	array = append(array, str)
	return array
}
