package piscine

func BasicJoin(elems []string) string {
	world := ""
	for _, letter := range elems {
		world = world + letter
	}
	return world
}
