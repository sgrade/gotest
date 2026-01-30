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
	for _, c := range word {
		i := c - 'a'
		if node.children[i] == nil {
			node.children[i] = newTrie()
		}
		node = node.children[i]
	}
	if node.id == -1 {
		*index++
		node.id = *index
	}
	return node.id
}

func update(current *int64, newVal int64) {
	if *current == -1 || newVal < *current {
		*current = newVal
	}
}

func minimumCost(source string, target string, original []string, changed []string, cost []int) int64 {
	n := len(source)
	numStrings := len(original)
	root := newTrie()

	idx := -1
	nodeCount := numStrings * 2
	graph := make([][]int, nodeCount)
	for i := range graph {
		graph[i] = make([]int, nodeCount)
		for j := range graph[i] {
			graph[i][j] = math.MaxInt32 / 2
		}
		graph[i][i] = 0
	}

	// Add all strings to trie and populate graph with conversion costs
	for i := 0; i < numStrings; i++ {
		fromID := add(root, original[i], &idx)
		toID := add(root, changed[i], &idx)
		graph[fromID][toID] = min(graph[fromID][toID], cost[i])
	}

	// Floyd-Warshall to find shortest paths
	size := idx + 1
	for k := 0; k < size; k++ {
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				graph[i][j] = min(graph[i][j], graph[i][k]+graph[k][j])
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
			update(&dp[j], base)
		}

		srcNode, tgtNode := root, root
		for i := j; i < n; i++ {
			srcNode = srcNode.children[source[i]-'a']
			tgtNode = tgtNode.children[target[i]-'a']
			if srcNode == nil || tgtNode == nil {
				break
			}
			if srcNode.id != -1 && tgtNode.id != -1 && graph[srcNode.id][tgtNode.id] != math.MaxInt32/2 {
				newVal := base + int64(graph[srcNode.id][tgtNode.id])
				update(&dp[i], newVal)
			}
		}
	}

	return dp[n-1]
}
