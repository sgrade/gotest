// 2110. Number of Smooth Descent Periods of a Stock
// https://leetcode.com/problems/number-of-smooth-descent-periods-of-a-stock/

package numberofsmoothdescentperiodsofastock

func getDescentPeriods(prices []int) int64 {
	ans, curLen := int64(1), int64(1)
	for i := 1; i < len(prices); i++ {
		if prices[i-1]-prices[i] == 1 {
			curLen++
		} else {
			curLen = 1
		}
		ans += curLen
	}
	return ans
}
