func isValidSudoku(board [][]byte) bool {

n := len(board)

	for i := range n {
		row := []byte{}
		col := []byte{}
		subGrid := []byte{}

		for j := range n {
			row = append(row, board[i][j])
			col = append(col, board[j][i])
		}

		sub_rl := (i / 3) * 3
		sub_cl := (i % 3) * 3

		for k := sub_rl; k < sub_rl+3; k++ {
			for l := sub_cl; l < sub_cl+3; l++ {
				subGrid = append(subGrid, board[k][l])
			}
		}

		if !checkValidBoardsec(row) {
			return false
		}

		if !checkValidBoardsec(col) {
			return false
		}

		if !checkValidBoardsec(subGrid) {
			return false
		}
	}

	return true
}

func checkValidBoardsec(boardSec []byte) bool {
	check_map := make(map[byte]bool)

	for _, val := range boardSec {

		if val != byte('.') {
			if _, ok := check_map[val]; ok {
				return false
			}

			check_map[val] = true
		}
	}

	return true
}
