// 3637. Trionic Array I
// https://leetcode.com/problems/trionic-array-i/

package trionicarrayi

func isTrionic(nums []int) bool {
	peaks, valleys := 0, 0
	for i := 1; i < len(nums)-1; i++ {
		if nums[i-1] < nums[i] && nums[i] > nums[i+1] {
			if peaks > 0 {
				return false
			}
			peaks++
		} else if nums[i-1] > nums[i] && nums[i] < nums[i+1] {
			if peaks == 0 {
				return false
			}
			valleys++
		} else if nums[i-1] == nums[i] || nums[i] == nums[i+1] {
			return false
		}
	}
	return peaks == 1 && valleys == 1
}
