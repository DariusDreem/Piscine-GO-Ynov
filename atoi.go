package piscine

func Atoi(s string) int {
	rev := ""
	for i := len([]rune(s)) - 1; i >= 0; i-- {
		rev = rev + string(s[i])
	}
	a := 0
	i := 1
	ispositive := false
	isnegative := false
	for _, element := range rev {
		if int(element) < 48 || int(element) > 57 {
			if int(element) == 43 && ispositive == false {
				ispositive = true
			} else if ispositive == true && isnegative == true {
				return 0
			} else if int(element) == 43 && ispositive == true {
				return 0
			} else if int(element) == 45 && isnegative == false {
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
	if isnegative == true && ispositive == true {
		return 0
	}
	return a
}
