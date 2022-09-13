package piscine

func SortIntegerTable(table []int) {
	tempo := 0
	for i := 0; i < len(table); i++ {
		for j := i + 1; j < len(table); j++ {
			if table[j] >= table[i] {
				continue
			} else {
				tempo = table[j]
				table[j] = table[i]
				table[i] = tempo
			}
		}
	}
}
