// 1502. Can Make Arithmetic Progression From Sequence
// https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence/

package canmakearithmeticprogressionfromsequence

import "sort"

func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr[:])
	ans := true
	d := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != d {
			ans = false
			break
		}
	}
	return ans
}
