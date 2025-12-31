// 2203. Last Day Where You Can Still Cross
// https://leetcode.com/problems/last-day-where-you-can-still-cross/

package lastdaywhereyoucanstillcross

// Based on Editorial's Approach 3: Disjoint Set Union (on land cells)

var directions = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

// UnionFind (Disjoint Set Union / Union-Find)
type UnionFind struct {
	parent []int
	rank   []int
}

// NewUnionFind creates a new UnionFind with n elements
func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &UnionFind{parent: parent, rank: rank}
}

// Find finds the root of x with path compression
func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // path compression
	}
	return uf.parent[x]
}

// Union merges sets containing x and y using union by rank
func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	if rootX == rootY {
		return
	}
	// union by rank
	if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
	} else if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
	} else {
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
	}
}

// IsConnected checks if x and y are in the same set
func (uf *UnionFind) IsConnected(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

func latestDayToCross(row int, col int, cells [][]int) int {
	dsu := NewUnionFind(row*col + 2)
	grid := make([][]int, row)
	for r := range grid {
		grid[r] = make([]int, col)
		for c := range grid[r] {
			grid[r][c] = 1
		}
	}

	const (
		water = 1
		land  = 0
	)

	topNode := 0
	bottomNode := row*col + 1

	for day := len(cells) - 1; day >= 0; day-- {
		r := cells[day][0] - 1
		c := cells[day][1] - 1
		grid[r][c] = land
		currIdx := r*col + c + 1

		for _, dir := range directions {
			newR, newC := r+dir[0], c+dir[1]
			if newR < 0 || newR >= row || newC < 0 || newC >= col {
				continue
			}
			if grid[newR][newC] == land {
				neighborIdx := newR*col + newC + 1
				dsu.Union(currIdx, neighborIdx)
			}
		}

		if r == 0 {
			dsu.Union(topNode, currIdx)
		}
		if r == row-1 {
			dsu.Union(bottomNode, currIdx)
		}
		if dsu.Find(topNode) == dsu.Find(bottomNode) {
			return day
		}
	}

	return -1
}
