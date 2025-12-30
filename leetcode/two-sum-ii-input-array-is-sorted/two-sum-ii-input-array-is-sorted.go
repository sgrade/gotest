// 167. Two Sum II - Input Array Is Sorted
// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/

package twosumiiinputarrayissorted

func twoSum(numbers []int, target int) []int {
	lo, hi := 0, len(numbers)-1
	for lo < hi {
		total := numbers[lo] + numbers[hi]
		if total == target {
			return []int{lo + 1, hi + 1}
		} else if total < target {
			lo++
		} else {
			hi--
		}
	}
	return []int{}
}
