package creategridwithexactlyonepath

import "strings"

func createGrid(m int, n int) []string {
	ans := make([]string, m)
	ans[0] = strings.Repeat(".", n)
	for i := 1; i < m; i++ {
		ans[i] = strings.Repeat("#", n-1) + "."
	}
	return ans
}
