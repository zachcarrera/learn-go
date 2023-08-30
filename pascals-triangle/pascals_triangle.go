package pascal

func Triangle(n int) [][]int {
	triangle := make([][]int, n)
	for i := range triangle {
		triangle[i] = make([]int, i+1)
		triangle[i][0] = 1
		triangle[i][i] = 1
		if i == 0 {
			continue
		}
		for j := range triangle[i][1:i] {
			triangle[i][j+1] = triangle[i-1][j] + triangle[i-1][j+1]
		}
	}
	return triangle
}
