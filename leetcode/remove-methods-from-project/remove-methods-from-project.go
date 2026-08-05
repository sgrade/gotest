// 3310. Remove Methods From Project
// https://leetcode.com/problems/remove-methods-from-project/

package removemethodsfromproject

// Based on Editorial's Approach: Searching
func remainingMethods(n, k int, invocations [][]int) []int {
	// Build directed graph: caller -> callees.
	graph := make([][]int, n)
	for _, inv := range invocations {
		caller, callee := inv[0], inv[1]
		graph[caller] = append(graph[caller], callee)
	}

	// BFS from k: mark every method reachable from the buggy method as suspicious.
	suspicious := make([]bool, n)
	suspicious[k] = true
	queue := []int{k}
	for len(queue) > 0 {
		method := queue[0]
		queue = queue[1:]
		for _, callee := range graph[method] {
			if !suspicious[callee] {
				suspicious[callee] = true
				queue = append(queue, callee)
			}
		}
	}

	// Unsafe to remove if any non-suspicious method invokes a suspicious one.
	for _, inv := range invocations {
		caller, callee := inv[0], inv[1]
		if !suspicious[caller] && suspicious[callee] {
			ans := make([]int, n)
			for i := range n {
				ans[i] = i
			}
			return ans
		}
	}

	// Otherwise keep only the non-suspicious methods.
	var ans []int
	for i := range n {
		if !suspicious[i] {
			ans = append(ans, i)
		}
	}
	return ans
}
