// 3507. Minimum Pair Removal to Sort Array I
// https://leetcode.com/problems/minimum-pair-removal-to-sort-array-i/

package minimumpairremovaltosortarrayi

func minimumPairRemoval(nums []int) int {
	operations := 0

	for len(nums) > 1 {
		nonDecreasing := true
		minSum := int(1e9)
		idx := -1

		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				nonDecreasing = false
			}
			curSum := nums[i] + nums[i+1]
			if curSum < minSum {
				minSum = curSum
				idx = i
			}
		}

		if nonDecreasing {
			break
		}

		operations++
		nums[idx] = minSum
		nums = append(nums[:idx+1], nums[idx+2:]...)
	}

	return operations
}
