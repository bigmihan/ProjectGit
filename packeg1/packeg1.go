package packeg1

// commit2
func Sum(a int, b int) int {
	return a + b
}

func MySubtract(a int, b int) int {
	return a - b
}

func Mod(a int) int {
	if a > 0 {
		return a
	}
	return -1 * a
}

func Zero() int {
	return 0
}
