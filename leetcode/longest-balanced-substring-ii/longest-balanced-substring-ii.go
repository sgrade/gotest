// 3714. Longest Balanced Substring II
// https://leetcode.com/problems/longest-balanced-substring-ii/

package longestbalancedsubstringii

func longestBalanced(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	// Prefix sums for counts of 'a', 'b', 'c'
	a := make([]int, n+1)
	b := make([]int, n+1)
	c := make([]int, n+1)
	for i := 0; i < n; i++ {
		a[i+1] = a[i]
		b[i+1] = b[i]
		c[i+1] = c[i]
		switch s[i] {
		case 'a':
			a[i+1]++
		case 'b':
			b[i+1]++
		case 'c':
			c[i+1]++
		}
	}

	ans := 0

	// Case 1: substring consists of a single repeated letter
	curLen := 0
	var last byte
	for i := 0; i < n; i++ {
		if i == 0 || s[i] != last {
			last = s[i]
			curLen = 1
		} else {
			curLen++
		}
		if curLen > ans {
			ans = curLen
		}
	}

	// If the entire string is one letter, this is optimal already.
	if ans == n {
		return ans
	}

	// Case 2: substrings consisting of exactly two different letters.
	longestPairBalanced := func(x, y, z byte) int {
		best := 0
		xCount, yCount := 0, 0
		// diff -> earliest index (1-based) where this diff occurred
		mp := map[int]int{0: 0}

		for i := 0; i < n; i++ {
			ch := s[i]
			if ch == z {
				// Third letter encountered – reset all state.
				mp = map[int]int{0: i + 1}
				xCount, yCount = 0, 0
				continue
			}
			if ch == x {
				xCount++
			} else if ch == y {
				yCount++
			}
			diff := xCount - yCount
			if left, ok := mp[diff]; ok {
				if i-left+1 > best {
					best = i - left + 1
				}
			} else {
				// Record the earliest index for this diff.
				mp[diff] = i + 1
			}
		}
		return best
	}

	// Try all pairs for the two-letter case.
	if v := longestPairBalanced('a', 'b', 'c'); v > ans {
		ans = v
	}
	if v := longestPairBalanced('a', 'c', 'b'); v > ans {
		ans = v
	}
	if v := longestPairBalanced('b', 'c', 'a'); v > ans {
		ans = v
	}

	// Case 3: substrings with all three letters.
	// We use the condition:
	//   between j and i, counts equal if:
	//   a[i]-a[j] == b[i]-b[j] == c[i]-c[j]
	// This can be rearranged to:
	//   2*a[k] - b[k] - c[k] is equal at i and j.
	type idxList []int
	mp3 := map[int]idxList{}
	mp3[0] = append(mp3[0], 0)

	for i := 0; i < n; i++ {
		diff := 2*a[i+1] - b[i+1] - c[i+1]
		if prev, ok := mp3[diff]; ok {
			// Indices in prev are increasing; the earliest one that satisfies
			// the equal-count condition will give the longest substring for this i.
			for _, idx := range prev {
				ac := a[i+1] - a[idx]
				bc := b[i+1] - b[idx]
				cc := c[i+1] - c[idx]
				if ac == bc && ac == cc {
					if i-idx+1 > ans {
						ans = i - idx + 1
					}
					// Earlier indices only give longer substrings; since we
					// iterate from earliest, we can stop after the first hit.
					break
				}
			}
		}
		mp3[diff] = append(mp3[diff], i+1)
	}

	return ans
}
