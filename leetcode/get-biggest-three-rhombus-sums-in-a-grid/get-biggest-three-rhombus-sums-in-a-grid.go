// 1878. Get Biggest Three Rhombus Sums in a Grid
// https://leetcode.com/problems/get-biggest-three-rhombus-sums-in-a-grid/

package getbiggestthreerhombussumsinagrid

// Based on Editorial's Approach: Enumerate All Rhombuses
// top3 tracks the three largest distinct values.
type top3 [3]int

func (t *top3) add(x int) {
	switch {
	case x > t[0]:
		t[2], t[1], t[0] = t[1], t[0], x
	case x != t[0] && x > t[1]:
		t[2], t[1] = t[1], x
	case x != t[0] && x != t[1] && x > t[2]:
		t[2] = x
	}
}

func (t *top3) result() []int {
	var r []int
	for _, v := range t {
		if v > 0 {
			r = append(r, v)
		}
	}
	return r
}

func getBiggestThree(grid [][]int) []int {
	m, n := len(grid), len(grid[0])

	// Prefix sums along \ and / diagonals (1-indexed, with padding).
	d1 := make([][]int, m+1)
	d2 := make([][]int, m+1)
	for i := range d1 {
		d1[i] = make([]int, n+2)
		d2[i] = make([]int, n+2)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			d1[i][j] = d1[i-1][j-1] + grid[i-1][j-1]
			d2[i][j] = d2[i-1][j+1] + grid[i-1][j-1]
		}
	}

	var best top3
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			best.add(grid[i][j])
			// Enumerate rhombuses with top vertex at (i, j) and half-height sz.
			for sz := 1; i+2*sz < m; sz++ {
				l, r := j-sz, j+sz
				if l < 0 || r >= n {
					break
				}
				mid, bot := i+sz, i+2*sz
				// Sum four edges via diagonal prefix sums, subtract double-counted vertices.
				s := (d2[mid+1][l+1] - d2[i][j+2]) +
					(d1[mid+1][r+1] - d1[i][j]) +
					(d1[bot+1][j+1] - d1[mid][l]) +
					(d2[bot+1][j+1] - d2[mid][r+2]) -
					(grid[i][j] + grid[bot][j] + grid[mid][l] + grid[mid][r])
				best.add(s)
			}
		}
	}
	return best.result()
}
