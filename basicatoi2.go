package piscine

func BasicAtoi2(s string) int {
	rev := ""
	for i := len([]rune(s)) - 1; i >= 0; i-- {
		rev = rev + string(s[i])
	}
	a := 0
	i := 1
	for _, element := range rev {
		if int(element) < 48 || int(element) > 57 {
			return 0
		} else {
			a = a + int(element-48)*i
			i = i * 10
		}
	}
	return a
}
