// 2483. Minimum Penalty for a Shop
// https://leetcode.com/problems/minimum-penalty-for-a-shop/

package minimumpenaltyforashop

func bestClosingTime(customers string) int {
	minPenalty, curPenalty, closingHour := 0, 0, 0
	for i, c := range customers {
		if c == 'Y' {
			curPenalty--
		} else {
			curPenalty++
		}
		if curPenalty < minPenalty {
			minPenalty = curPenalty
			closingHour = i + 1
		}
	}
	return closingHour
}
