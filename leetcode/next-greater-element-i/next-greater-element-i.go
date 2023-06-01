// 496. Next Greater Element I
// https://leetcode.com/problems/next-greater-element-i/

package nextgreaterelementi

import "github.com/emirpasic/gods/stacks/arraystack"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	st := arraystack.New()
	mp := make(map[int]int)
	for _, el := range nums2 {
		top, _ := st.Peek()
		for !st.Empty() && el > top.(int) {
			prev, _ := st.Pop()
			mp[prev.(int)] = el
			top, _ = st.Peek()
		}
		st.Push(el)
	}
	for !st.Empty() {
		prev, _ := st.Pop()
		mp[prev.(int)] = -1
	}
	for i := range nums1 {
		nums1[i] = mp[nums1[i]]
	}
	return nums1
}
