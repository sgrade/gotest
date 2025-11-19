package keepmultiplyingfoundvaluesbytwo

// 2154. Keep Multiplying Found Values by Two
// https://leetcode.com/problems/keep-multiplying-found-values-by-two/

func findFinalValue(nums []int, original int) int {
	st := make(map[int]bool)
	for _, num := range nums {
		st[num] = true
	}
	for st[original] {
		original *= 2
	}
	return original
}
