func isValidSudoku(board [][]byte) bool {
	rows := make([][9]bool, 9)
	columns := make([][9]bool, 9)
	boxes := make([][9]bool, 9)

	bi := 0

	for ri,row := range board {
		if ri == 3 || ri == 6 {
			bi++
		} else if bi != 0 {
			bi -= 2
		}
		for di, digit := range row {
			if di == 3 || di == 6 {
				bi++
			}
			if digit == '.' {
				continue
			}
			digitIndex := digit - '1'
			if rows[ri][digitIndex] || columns[di][digitIndex] || boxes[bi][digitIndex] {
				return false
			}
			rows[ri][digitIndex] = true
			columns[di][digitIndex] = true
			boxes[bi][digitIndex] = true
		}
	}

	return true;
}
