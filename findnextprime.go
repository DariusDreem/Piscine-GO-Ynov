package piscine

func FindNextPrime(nb int) int {
	prime := false
	if nb == 1 || nb <= 0 {
		return 2
	}
	i := nb
	for prime != true {
		if IsPrime(i) == true {
			prime = true
			break
		}
		i++
	}
	return i
}
