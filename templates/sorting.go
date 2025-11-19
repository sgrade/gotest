package templates

import "sort"

// Sort Integers
func sortInts() {
	nums := []int{3, 1, 4, 1, 5}
	sort.Ints(nums) // ascending
	sort.Sort(sort.Reverse(sort.IntSlice(nums))) // descending
}

// Sort Strings
func sortStrings() {
	words := []string{"banana", "apple", "cherry"}
	sort.Strings(words)
}

// Custom Comparator - Sort Slice
func customSort() {
	nums := []int{3, 1, 4, 1, 5}
	// ascending
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	// descending
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
}

// Sort 2D Array / Intervals
func sortIntervals() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}}
	// sort by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// sort by end time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
}

// Sort Struct
type Person struct {
	Name string
	Age  int
}

func sortStructs() {
	people := []Person{{"Alice", 30}, {"Bob", 25}}
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
}

// Binary Search after Sorting
func searchSorted(nums []int, target int) int {
	sort.Ints(nums)
	return sort.SearchInts(nums, target)
}

