package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) <= 1 {
		return true
	}
	if f(a[0], a[1]) < 0 {
		for i := 0; i < len(a)-1; i++ {
			number := f(a[i], a[i+1])
			if number > 0 {
				return false
			}
		}
	} else {
		for i := 0; i < len(a)-1; i++ {
			number := f(a[i], a[i+1])
			if number < 0 {
				return false
			}
		}
	}
	return true
}
