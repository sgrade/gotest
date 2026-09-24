// 3550. Smallest Index With Digit Sum Equal to Index
// https://leetcode.com/problems/smallest-index-with-digit-sum-equal-to-index/

package smallestindexwithdigitsumequaltoindex

func smallestIndex(nums []int) int {
	// First index whose digits sum to the index itself.
	for i, num := range nums {
		if digitSum(num) == i {
			return i
		}
	}
	return -1
}

func digitSum(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}
