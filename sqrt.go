package piscine

func Sqrt(nb int) int {
	i := 0
	for {
		if i*i > nb {
			return 0
		} else if i*i == nb {
			return i
		} else {
			i++
		}
	}
}
