package piscine

func Any(f func(string) bool, a []string) bool {
	oui := false
	for _, str := range a {
		oui := f(str)
		if oui == true {
			return oui
		} else {
			continue
		}
	}
	return oui
}
