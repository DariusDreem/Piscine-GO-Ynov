package piscine

func IsPrime(nb int) bool {
	if nb%2 == 0 && nb != 2 {
		return false
	} else if nb%3 == 0 && nb != 3 {
		return false
	} else if nb%5 == 0 && nb != 5 {
		return false
	} else if nb%7 == 0 && nb != 7 {
		return false
	} else if nb == 1 {
		return false
	}
	return true
}
