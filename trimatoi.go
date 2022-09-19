package piscine

func TrimAtoi(s string) int {
	rev := ""
	for i := len([]rune(s)) - 1; i >= 0; i-- {
		rev = rev + string(s[i])
	}
	pabo := 0
	a := 0
	i := 1
	ispositive := false
	isnegative := false
	for _, element := range s {
		if int(element) == 43 && ispositive == false {
			ispositive = true
		} else if ispositive == true && isnegative == true {
			continue
		} else if int(element) == 43 && ispositive == true {
			continue
		} else if int(element) == 45 && isnegative == false {
			isnegative = true
		} else if int(element) >= 48 && int(element) <= 57 {
			break
		}
	}
	for _, element := range rev {
		if int(element) >= 48 && int(element) <= 57 {
			pabo++
			a = a + int(element-48)*i
			i = i * 10
		} else {
			continue
		}
	}

	if isnegative == true {
		a = -a
	}
	return a
}
