// 496. Next Greater Element I
// https://leetcode.com/problems/next-greater-element-i/

package nextgreaterelementi

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	st := []int{}
	mp := map[int]int{}
	for _, el := range nums2 {
		for len(st) > 0 && el > st[len(st)-1] {
			mp[st[len(st)-1]] = el
			st = st[:len(st)-1]
		}
		st = append(st, el)
	}
	for i := len(st) - 1; i >= 0; i-- {
		mp[st[i]] = -1
	}
	for i := range nums1 {
		nums1[i] = mp[nums1[i]]
	}
	return nums1
}
