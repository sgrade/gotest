// 999. Available Captures for Rook
// https://leetcode.com/problems/available-captures-for-rook/

package availablecapturesforrook

func numRookCaptures(board [][]byte) int {
	row, col := -1, -1
	for r := 0; r < 8; r++ {
		found := false
		for c := 0; c < 8; c++ {
			if board[r][c] == 'R' {
				found = true
				row = r
				col = c
				break
			}
		}
		if found {
			break
		}
	}

	check := func(r, c int) int {
		if board[r][c] == 'B' {
			return 0
		} else if board[r][c] == 'p' {
			return 1
		}
		return -1
	}

	pawns := 0
	for r := row + 1; r < 8; r++ {
		if v := check(r, col); v != -1 {
			pawns += v
			break
		}
	}
	for r := row - 1; r >= 0; r-- {
		if v := check(r, col); v != -1 {
			pawns += v
			break
		}
	}
	for c := col + 1; c < 8; c++ {
		if v := check(row, c); v != -1 {
			pawns += v
			break
		}
	}
	for c := col - 1; c >= 0; c-- {
		if v := check(row, c); v != -1 {
			pawns += v
			break
		}
	}
	return pawns
}
