package piscine

func ConcatParams(args []string) string {
	string := ""
	i := 0
	for _, word := range args {
		i++
		if i != len(args) {
			string += word
			string += "\n"
		} else {
			string += word
		}
	}
	return string
}
