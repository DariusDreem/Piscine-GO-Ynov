package piscine

func IterativePower(nb int, power int) int {

	if nb < 0 || power < 0 {
		return 0
	}

	result := nb ^ power
	return result
}
