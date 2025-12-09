// 3583. Count Special Triplets
// https://leetcode.com/problems/count-special-triplets/

package countspecialtriplets

func specialTriplets(nums []int) int {
	counter := make(map[int]int)
	for _, num := range nums {
		counter[num]++
	}

	const MOD = 1e9 + 7
	ans := 0
	left_counter := make(map[int]int)
	for _, num := range nums {
		target := num * 2
		left_count := left_counter[target]
		left_counter[num]++
		right_count := counter[target] - left_counter[target]
		ans = (ans + left_count*right_count) % MOD
	}
	return ans
}
