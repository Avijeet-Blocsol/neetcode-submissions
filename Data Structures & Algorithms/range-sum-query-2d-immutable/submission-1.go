// The key is to realize that for a 2D case, the prefix sum depends on the immidiate top and left of
// a particular cell. We subtract the diagnol element because it is counted twice.

// We add it when calculating result because it is subtracted twice.

type NumMatrix struct {
	prefixMatrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m := len(matrix)
	n := len(matrix[0])

	p := make([][]int, m+1)

	for i := 0; i < m+1; i++ {
		p[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			p[i+1][j+1] = matrix[i][j] + p[i][j+1] + p[i+1][j] - p[i][j]
		}
	}

	return NumMatrix{
		prefixMatrix: p,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	total := this.prefixMatrix[row2+1][col2+1]
	upper := this.prefixMatrix[row1][col2+1]
	left := this.prefixMatrix[row2+1][col1]
	double := this.prefixMatrix[row1][col1]

	return total - upper - left + double
}

// Your NumMatrix object will be instantiated and called as such:
// obj := Constructor(matrix)
// param_1 := obj.SumRegion(row1,col1,row2,col2)
