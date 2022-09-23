package piscine

func MakeRange(min, max int) []int {
	size := max - min
	if min < max {
		array := make([]int, size)
		for i := 0; min < max; i++ {
			array[i] = min
			min++
		}
		return array
	}
	return nil
}
