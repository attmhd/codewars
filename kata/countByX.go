package kata

func CountByX(x, n int) []int {
	var i int
	a := []int{}
	for i = x; i <= x*n; i += x {
		a = append(a, i)
	}

	return a
}
