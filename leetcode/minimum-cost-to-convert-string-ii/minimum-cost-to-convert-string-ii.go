package minimumcosttoconvertstringii

import "math"

// Based on Editorial's Approach: Trie + Floyd's Algorithm + Dynamic Programming
type Trie struct {
	children [26]*Trie
	id       int
}

func newTrie() *Trie {
	return &Trie{id: -1}
}

func add(node *Trie, word string, index *int) int {
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if node.children[idx] == nil {
			node.children[idx] = newTrie()
		}
		node = node.children[idx]
	}
	if node.id == -1 {
		*index++
		node.id = *index
	}
	return node.id
}

func minimumCost(source string, target string, original []string, changed []string, cost []int) int64 {
	n := len(source)
	numStrings := len(original)
	root := newTrie()
	idx := -1

	// Build trie and collect conversion edges
	type edge struct {
		from, to, cost int
	}
	edges := make([]edge, numStrings)
	for i := 0; i < numStrings; i++ {
		fromID := add(root, original[i], &idx)
		toID := add(root, changed[i], &idx)
		edges[i] = edge{fromID, toID, cost[i]}
	}

	// Allocate graph with actual size
	size := idx + 1
	const inf = math.MaxInt32 / 2
	graph := make([][]int, size)
	for i := range graph {
		graph[i] = make([]int, size)
		for j := range graph[i] {
			graph[i][j] = inf
		}
		graph[i][i] = 0
	}

	// Populate graph with conversion costs
	for i := 0; i < numStrings; i++ {
		e := edges[i]
		if e.cost < graph[e.from][e.to] {
			graph[e.from][e.to] = e.cost
		}
	}

	// Floyd-Warshall with early termination
	for k := 0; k < size; k++ {
		for i := 0; i < size; i++ {
			if graph[i][k] == inf {
				continue
			}
			for j := 0; j < size; j++ {
				if graph[k][j] == inf {
					continue
				}
				if newDist := graph[i][k] + graph[k][j]; newDist < graph[i][j] {
					graph[i][j] = newDist
				}
			}
		}
	}

	dp := make([]int64, n)
	for i := range dp {
		dp[i] = -1
	}

	for j := 0; j < n; j++ {
		if j > 0 && dp[j-1] == -1 {
			continue
		}

		base := int64(0)
		if j > 0 {
			base = dp[j-1]
		}

		if source[j] == target[j] {
			if dp[j] == -1 || base < dp[j] {
				dp[j] = base
			}
		}

		srcNode, tgtNode := root, root
		for i := j; i < n; i++ {
			if srcNode == nil || tgtNode == nil {
				break
			}
			srcNode = srcNode.children[source[i]-'a']
			tgtNode = tgtNode.children[target[i]-'a']
			if srcNode == nil || tgtNode == nil {
				break
			}
			if srcNode.id != -1 && tgtNode.id != -1 {
				if graphCost := graph[srcNode.id][tgtNode.id]; graphCost < inf {
					newVal := base + int64(graphCost)
					if dp[i] == -1 || newVal < dp[i] {
						dp[i] = newVal
					}
				}
			}
		}
	}

	return dp[n-1]
}
