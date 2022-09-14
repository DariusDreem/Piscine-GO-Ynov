package piscine

func IsPrime(nb int) bool {
	if nb == 1 || nb < 0 {
		return false
	}
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return false
		} else {
			return true
		}
	}
	return true
}
