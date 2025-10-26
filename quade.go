package qds

func QuadE(x, y int) string {
	result := ""
	if x <= 0 || y <= 0 {
		return ""
	}
	for row := 0; row < y; row++ {
		for col := 0; col < x; col++ {
			if row == 0 && col == 0 {
				result += string('A')
			} else if row == 0 && col == x-1 {
				result += string('C')
			} else if row == y-1 && col == 0 {
				result += string('C')
			} else if row == y-1 && col == x-1 {
				result += string('A')
			} else if row == 0 || row == y-1 || col == 0 || col == x-1 {
				result += string('B')
			} else {
				result += string(' ')
			}
		}
		result += string('\n')
	}
	return result
}
