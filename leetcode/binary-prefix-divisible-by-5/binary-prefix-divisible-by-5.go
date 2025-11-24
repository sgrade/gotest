package binaryprefixdivisibleby5

func prefixesDivBy5(nums []int) []bool {
	n, sm := len(nums), 0
	ans := make([]bool, n)
	for i := range nums {
		sm *= 2
		sm = (sm + nums[i]) % 5
		ans[i] = sm == 0
	}
	return ans
}
