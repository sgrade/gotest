// 1. Two Sum
// https://leetcode.com/problems/two-sum/

package twosum

func twoSum(nums []int, target int) []int {
	numToIdx := make(map[int]int)
	for i, num := range nums {
		need := target - num
		j, ok := numToIdx[need]
		if ok {
			return []int{i, j}
		}
		numToIdx[num] = i
	}
	return []int{}
}
