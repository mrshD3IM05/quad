package quad

func QuadA(x, y int) string {
	result := ""
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			r := ' '
			if (j == 1 || j == x) && (i == 1 || i == y) {
				r = 'o'
			} else if (j == 1 || j == x) && (i != 1 && i != y) {
				r = '|'
			} else if (j != 1 && j != x) && (i == 1 || i == y) {
				r = '-'
			}
			result += string(r)
		}
		result += string('\n')
	}
	return result
}
