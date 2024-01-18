// 2320. Count Number of Ways to Place Houses
// https://leetcode.com/problems/count-number-of-ways-to-place-houses/

package countnumberofwaystoplacehouses

const MOD = 1e9 + 7

func countHousePlacements(n int) int {
	prev_prev, prev, cur := 1, 2, 2
	for i := 2; i <= n; i++ {
		cur = (prev_prev + prev) % MOD
		prev_prev = prev
		prev = cur
	}
	cur = cur * cur % MOD
	return cur
}
