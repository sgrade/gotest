package templates

// Backtracking - Subsets
func subsets(nums []int) [][]int {
	result := [][]int{}
	current := []int{}
	var backtrack func(start int)
	backtrack = func(start int) {
		// add current subset
		temp := make([]int, len(current))
		copy(temp, current)
		result = append(result, temp)
		// explore further
		for i := start; i < len(nums); i++ {
			current = append(current, nums[i])
			backtrack(i + 1)
			current = current[:len(current)-1] // backtrack
		}
	}
	backtrack(0)
	return result
}

// Backtracking - Permutations
func permutations(nums []int) [][]int {
	result := [][]int{}
	used := make([]bool, len(nums))
	current := []int{}
	var backtrack func()
	backtrack = func() {
		if len(current) == len(nums) {
			temp := make([]int, len(current))
			copy(temp, current)
			result = append(result, temp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			current = append(current, nums[i])
			used[i] = true
			backtrack()
			used[i] = false
			current = current[:len(current)-1]
		}
	}
	backtrack()
	return result
}

// Backtracking - Combinations (Target Sum)
func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	current := []int{}
	var backtrack func(start, sum int)
	backtrack = func(start, sum int) {
		if sum == target {
			temp := make([]int, len(current))
			copy(temp, current)
			result = append(result, temp)
			return
		}
		if sum > target {
			return
		}
		for i := start; i < len(candidates); i++ {
			current = append(current, candidates[i])
			backtrack(i, sum+candidates[i]) // i for reuse, i+1 for no reuse
			current = current[:len(current)-1]
		}
	}
	backtrack(0, 0)
	return result
}

