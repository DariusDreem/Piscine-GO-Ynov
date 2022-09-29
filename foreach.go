package piscine

func ForEach(f func(int), a []int) {
	for _, nbr := range a {
		print(nbr)
	}
}
