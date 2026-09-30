// 1111. Maximum Nesting Depth of Two Valid Parentheses Strings
// https://leetcode.com/problems/maximum-nesting-depth-of-two-valid-parentheses-strings/

package maximumnestingdepthoftwovalidparenthesesstrings

func maxDepthAfterSplit(seq string) []int {
	ans := make([]int, len(seq))
	depth := 0
	for i, c := range seq {
		if c == '(' {
			depth++
			ans[i] = depth % 2
		} else {
			ans[i] = depth % 2
			depth--
		}
	}
	return ans
}
