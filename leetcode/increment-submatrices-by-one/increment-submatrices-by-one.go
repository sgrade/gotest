package incrementsubmatricesbyone

func rangeAddQueries(n int, queries [][]int) [][]int {
	diff_arr := make([][]int, n+1)
	for i := range n + 1 {
		diff_arr[i] = make([]int, n+1)
	}
	for _, q := range queries {
		r1, c1, r2, c2 := q[0], q[1], q[2], q[3]
		diff_arr[r1][c1]++
		diff_arr[r2+1][c1]--
		diff_arr[r1][c2+1]--
		diff_arr[r2+1][c2+1]++
	}

	ans := make([][]int, n)
	for i := range n {
		ans[i] = make([]int, n)
	}

	for r := range n {
		for c := range n {
			col_diff, row_diff, diag_diff := 0, 0, 0
			if r != 0 {
				col_diff = ans[r-1][c]
			}
			if c != 0 {
				row_diff = ans[r][c-1]
			}
			if r != 0 && c != 0 {
				diag_diff = ans[r-1][c-1]
			}
			ans[r][c] = diff_arr[r][c] + col_diff + row_diff - diag_diff
		}
	}
	return ans
}
