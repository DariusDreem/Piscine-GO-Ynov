package piscine

func Index(s string, toFind string) int {
	for index, letter := range s {
		for i := 0; i < index; i++ {
			if int32(toFind[i]) == letter && letter+1 == int32(toFind[i]+1) {
				return index
			} else {
				continue
			}
		}
	}
	return -1
}
