package countoddnumbersinanintervalrange

func countOdds(low int, high int) int {
	ans := (high + 1 - low) / 2
	if low%2 == 1 && high%2 == 1 {
		ans += 1
	}
	return ans
}
