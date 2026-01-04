// 1390. Four Divisors
// https://leetcode.com/problems/four-divisors/

package fourdivisors

// Based on
// https://leetcode.com/problems/four-divisors/solutions/7463933/b-ks-solutionjust-math-shortest-code-of-07gxj
func sumFourDivisors(nums []int) int {
	totalSum := 0
	for _, n := range nums {
		totalSum += getSumOfFactors(n)
	}
	return totalSum
}

func getSumOfFactors(n int) int {
	curSum, factors := 0, 0
	for divisor1 := 2; divisor1*divisor1 <= n; divisor1++ {
		if n%divisor1 == 0 {
			divisor2 := n / divisor1
			if divisor1 == divisor2 || factors > 0 {
				return 0
			}
			curSum += divisor1 + divisor2
			factors += 2
		}
	}
	if factors == 0 {
		return 0
	}
	return 1 + curSum + n
}
