// 3600. Maximize Spanning Tree Stability with Upgrades
// https://leetcode.com/problems/maximize-spanning-tree-stability-with-upgrades/
//
// Stability of a spanning tree = its minimum edge weight.
// Each edge is either required (must be in the tree) or optional.
// We may upgrade up to k optional edges, doubling their weight.
//
// Approach: binary search on the answer (minimum stability threshold).
// For each threshold, greedily try to build a spanning tree using
// required edges + optional edges sorted by weight descending,
// spending upgrades only when an edge's raw weight < threshold but
// doubled weight >= threshold.

package maximizespanningtreestabilitywithupgrades

import "sort"

// Based on Editorial's Maximize Spanning Tree Stability with Upgrades
type unionFind struct {
	parent []int
}

func newUnionFind(n int) *unionFind {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &unionFind{parent: parent}
}

func (uf *unionFind) clone() *unionFind {
	parent := make([]int, len(uf.parent))
	copy(parent, uf.parent)
	return &unionFind{parent: parent}
}

func (uf *unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *unionFind) union(x, y int) bool {
	rx, ry := uf.find(x), uf.find(y)
	if rx == ry {
		return false
	}
	uf.parent[rx] = ry
	return true
}

func maxStability(n int, edges [][]int, k int) int {
	if len(edges) < n-1 {
		return -1
	}

	var required, optional [][]int
	for _, e := range edges {
		if e[3] == 1 {
			required = append(required, e)
		} else {
			optional = append(optional, e)
		}
	}

	if len(required) > n-1 {
		return -1 // Too many required edges to form a tree.
	}

	sort.Slice(optional, func(i, j int) bool {
		return optional[i][2] > optional[j][2]
	})

	// Build the initial spanning forest from required edges.
	base := newUnionFind(n)
	requiredCount := 0
	minRequiredStability := 200_000
	for _, e := range required {
		src, dst, stability := e[0], e[1], e[2]
		if !base.union(src, dst) {
			return -1 // Required edges form a cycle.
		}
		requiredCount++
		minRequiredStability = min(minRequiredStability, stability)
	}

	// Answer can't exceed the weakest required edge (it's always in the tree).
	lo, hi := 0, minRequiredStability
	result := -1
	for lo <= hi {
		threshold := lo + (hi-lo)/2

		// Greedy: start from required edges, add optional ones by decreasing weight.
		uf := base.clone()
		treeEdges := requiredCount
		upgrades := 0

		for _, e := range optional {
			src, dst, stability := e[0], e[1], e[2]
			if uf.find(src) == uf.find(dst) {
				continue
			}
			if stability >= threshold {
				uf.union(src, dst)
				treeEdges++
			} else if upgrades < k && stability*2 >= threshold {
				uf.union(src, dst)
				treeEdges++
				upgrades++
			} else {
				break // Sorted descending; no further edge can meet the threshold.
			}
			if treeEdges == n-1 {
				break
			}
		}

		if treeEdges == n-1 {
			result = threshold
			lo = threshold + 1
		} else {
			hi = threshold - 1
		}
	}

	return result
}
