package piscine

func AtoiBase(s string, base string) int {
	loop := 0
	baseList := []string{}
	power := 1
	result := 0
	listResult := []int{}
	if len(base) < 2 {
		return 0
	}
	for _, element := range base {
		loop = 0
		if element == 43 || element == 45 {
			return 0
		}
		if len(baseList) == 0 {
			baseList = append(baseList, string(element))
			continue
		}
		for _, baseLetter := range baseList {
			if string(element) == baseLetter {
				return 0
			} else {
				continue
			}
		}
		if loop == 0 {
			baseList = append(baseList, string(element))
		}
	}
	for _, element := range s {
		for index, el := range base {
			if element == el {
				listResult = append(listResult, index)
			} else {
				continue
			}
		}
	}
	for i := len(listResult) - 1; i != -1; i-- {
		result = result + listResult[i]*power
		power *= len(base)
	}
	return result
}
