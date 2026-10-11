// 2778. Sum of Squares of Special Elements
// https://leetcode.com/problems/sum-of-squares-of-special-elements/

package sumofsquaresofspecialelements

func sumOfSquares(nums []int) int {
	n := len(nums)
	ans := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			ans += nums[i-1] * nums[i-1]
		}
	}
	return ans
}
