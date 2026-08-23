type NumMatrix struct {
	PrefixMatrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	n := len(matrix)
	m := len(matrix[0])

	p := make([][]int, n+1)

	for i := range n + 1 {
		p[i] = make([]int, m+1)
	}

	for i := range n {
		for j := range m {
			p[i+1][j+1] = p[i][j+1] + p[i+1][j] + matrix[i][j] - p[i][j]
		}
	}

	return NumMatrix{
		PrefixMatrix: p,
	}
}

func (nm *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	top := nm.PrefixMatrix[row1][col2+1]
	left := nm.PrefixMatrix[row2+1][col1]
	total := nm.PrefixMatrix[row2+1][col2+1]
	dedupe := nm.PrefixMatrix[row1][col1]

	return total - top - left + dedupe
}

// Your NumMatrix object will be instantiated and called as such:
// obj := Constructor(matrix)
// param_1 := obj.SumRegion(row1,col1,row2,col2)
