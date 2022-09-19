package piscine

func Join(strs []string, sep string) string {
	world := ""
	i := 0
	po := 0
	for _, letter := range strs {
		if po > 0 {
			i++
			world = world + sep + letter
		} else {
			world = world + letter
			po++
		}
	}
	return world
}
