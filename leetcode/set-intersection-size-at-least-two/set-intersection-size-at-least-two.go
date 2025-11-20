package setintersectionsizeatleasttwo

import "sort"

// 757. Set Intersection Size At Least Two
// https://leetcode.com/problems/set-intersection-size-at-least-two/

func intersectionSizeTwo(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		s1, e1, s2, e2 := intervals[i][0], intervals[i][1], intervals[j][0], intervals[j][1]
		if s1 != s2 {
			return s1 < s2
		}
		return e1 > e2
	})

	to_pick_from_interval := make([]int, len(intervals))
	for i := range len(intervals) {
		to_pick_from_interval[i] = 2
	}
	min_len := 0

	for i := len(intervals) - 1; i >= 0; i-- {
		cur_start, to_pick := intervals[i][0], to_pick_from_interval[i]
		for point := cur_start; point < cur_start+to_pick; point++ {
			for j := 0; j <= i; j++ {
				cur_end := intervals[j][1]
				if to_pick_from_interval[j] > 0 && point <= cur_end {
					to_pick_from_interval[j]--
				}
			}
			min_len++
		}
	}
	return min_len
}
