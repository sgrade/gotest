// 2996. Smallest Missing Integer Greater Than Sequential Prefix Sum
// https://leetcode.com/problems/smallest-missing-integer-greater-than-sequential-prefix-sum/

package smallestmissingintegergreaterthansequentialprefixsum

func missingInteger(nums []int) int {
	sm, i := nums[0], 1
	for i < len(nums) {
		if nums[i]-nums[i-1] == 1 {
			sm += nums[i]
			i++
		} else {
			break
		}
	}

	st := make(map[int]bool, len(nums))
	for _, num := range nums {
		st[num] = true
	}
	for st[sm] {
		sm++
	}
	return sm
}
