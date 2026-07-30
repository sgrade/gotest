// 3014. Minimum Number of Pushes to Type Word I
// https://leetcode.com/problems/minimum-number-of-pushes-to-type-word-i/

package minimumnumberofpushestotypewordi

func minimumPushes(word string) int {
	ans := 0
	for i := 0; i < len(word); i++ {
		ans += (i/8 + 1)
	}
	return ans
}
