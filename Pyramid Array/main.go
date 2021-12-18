package kata

func Pyramid(n int) [][]int {
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, i+1)
		for k := range a[i] {
			a[i][k] = 1
		}
	}
	return a
}
