// 744. Find Smallest Letter Greater Than Target
// https://leetcode.com/problems/find-smallest-letter-greater-than-target/

package findsmallestlettergreaterthantarget

import "sort"

func nextGreatestLetter(letters []byte, target byte) byte {
	n := len(letters)
	idx := sort.Search(n, func(i int) bool {
		return letters[i] > target
	})

	if idx == n {
		return letters[0]
	}
	return letters[idx]
}
