// 3517. Smallest Palindromic Rearrangement I
// https://leetcode.com/problems/smallest-palindromic-rearrangement-i/

package smallestpalindromicrearrangementi

func smallestPalindrome(s string) string {
	n := len(s)
	var cnt [26]int
	for _, c := range s {
		cnt[c-'a']++
	}

	ans := make([]byte, n)
	left := 0
	mid := -1
	for i := range 26 {
		if cnt[i]&1 != 0 {
			mid = i
		}
		half := cnt[i] / 2
		ch := byte('a' + i)
		for j := 0; j < half; j++ {
			ans[left] = ch
			ans[n-1-left] = ch
			left++
		}
	}
	if mid != -1 {
		ans[n/2] = byte('a' + mid)
	}
	return string(ans)
}
