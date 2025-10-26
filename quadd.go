package quad

func QuadD(x, y int) string { //ahaddou
	result := ""
	if x < 1 || y < 1 {
		return ""
	}
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			r := ' '
			if (j == 1 || j == x) && (i == 1 || i == y) {
				if j != 1 {
					r = 'C'
				} else {
					r = 'A'
				}
			} else if (j == 1 || j == x) && (i != 1 && i != y) || (j != 1 && j != x) && (i == 1 || i == y) {
				r = 'B'
			}
			result += string(r)
		}
		result += string('\n')
	}
	return result
}
