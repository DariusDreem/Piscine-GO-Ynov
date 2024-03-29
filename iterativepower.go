package piscine

func IterativePower(nb int, power int) int {
	tempo := nb
	if power == 0 {
		return 1
	} else if nb == 0 || power < 0 {
		return 0
	}
	for i := 1; i < power; i++ {
		nb = tempo * nb
	}
	return nb
}
