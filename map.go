package piscine

func Map(f func(int) bool, a []int) []bool {
	var Yee []bool
	bol := false
	for i, nbr := range a {
		i++

		if IsPrime(nbr) == true {
			bol = false
		} else if IsPrime(nbr) == false {
			bol = true
		}
		Yee = append(Yee, bol)
	}
	return Yee
}
