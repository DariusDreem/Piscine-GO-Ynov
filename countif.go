package piscine

func CountIf(f func(string) bool, tab []string) int {
	bu := 0
	for _, str := range tab {
		oui := f(str)
		if oui == true {
			bu++
			continue
		} else {
			continue
		}
	}
	return bu
}
