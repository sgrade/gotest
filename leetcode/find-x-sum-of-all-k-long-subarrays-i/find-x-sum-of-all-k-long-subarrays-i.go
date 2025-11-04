package findxsumofallklongsubarraysi

import "sort"

func findXSum(nums []int, k int, x int) []int {
	n := len(nums)
	cntr := make(map[int]int)
	for i := range k - 1 {
		cntr[nums[i]]++
	}

	type Pair struct {
		cnt int
		num int
	}

	ans := make([]int, 0, n-k+1)
	for i := k - 1; i < n; i++ {
		cntr[nums[i]]++
		freq := make([]Pair, 0, len(cntr))
		for num, cnt := range cntr {
			freq = append(freq, Pair{cnt, num})
		}
		sort.Slice(freq, func(i, j int) bool {
			if freq[i].cnt != freq[j].cnt {
				return freq[i].cnt > freq[j].cnt
			}
			return freq[i].num > freq[j].num
		})
		sm := 0
		for j := 0; j < x && j < len(freq); j++ {
			sm += freq[j].cnt * freq[j].num
		}
		ans = append(ans, sm)
		if cntr[nums[i-k+1]] == 1 {
			delete(cntr, nums[i-k+1])
		} else {
			cntr[nums[i-k+1]]--
		}
	}
	return ans
}
