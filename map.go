package piscine

func Map(f func(int) bool, a []int) []bool {
	prime := make([]bool, len(a))
	i := 0
	for _, nbr := range a {
		prime[i] = f(nbr)
		i++
	}
	return prime
}
