package piscine

func Atoi(s string) int {
	rev := ""
	for i := len([]rune(s)) - 1; i >= 0; i-- {
		rev = rev + string(s[i])
	}
	moins := 0
	plus := 0
	a := 0
	i := 1
	isnegative := false
	for _, element := range rev {
		if int(element) < 48 || int(element) > 57 {
			if int(element) == 43 && plus < 1 {
				plus++
			} else if int(element) == 45 && moins == 1 {
				return 0
			} else if int(element) == 43 && plus == 1 {
				return 0
			} else if int(element) == 45 && moins < 1 {
				moins++
				//z01.PrintRune(45)
				isnegative = true

			} else {
				return 0
			}
		} else {
			a = a + int(element-48)*i
			i = i * 10
		}
	}
	if isnegative {
		a = -a
	}
	return a
}
