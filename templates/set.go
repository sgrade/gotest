package templates

// Set - Basic Operations
func setBasics() {
	set := make(map[int]bool)
	// add
	set[1] = true
	set[2] = true
	// contains
	if set[1] {
		// found
	}
	// remove
	delete(set, 1)
	// size
	size := len(set)
	_ = size
}

// Union - A ∪ B (all elements from both sets)
func union(set1, set2 map[int]bool) map[int]bool {
	result := make(map[int]bool)
	for k := range set1 {
		result[k] = true
	}
	for k := range set2 {
		result[k] = true
	}
	return result
}

// Intersection - A ∩ B (common elements)
func intersection(set1, set2 map[int]bool) map[int]bool {
	result := make(map[int]bool)
	for k := range set1 {
		if set2[k] {
			result[k] = true
		}
	}
	return result
}

// Difference - A - B (elements in A but not in B)
func difference(set1, set2 map[int]bool) map[int]bool {
	result := make(map[int]bool)
	for k := range set1 {
		if !set2[k] {
			result[k] = true
		}
	}
	return result
}

// Symmetric Difference - A △ B (elements in either but not both)
func symmetricDifference(set1, set2 map[int]bool) map[int]bool {
	result := make(map[int]bool)
	for k := range set1 {
		if !set2[k] {
			result[k] = true
		}
	}
	for k := range set2 {
		if !set1[k] {
			result[k] = true
		}
	}
	return result
}

// Subset - Is set1 ⊆ set2?
func isSubset(set1, set2 map[int]bool) bool {
	for k := range set1 {
		if !set2[k] {
			return false
		}
	}
	return true
}

// Convert slice to set
func sliceToSet(arr []int) map[int]bool {
	set := make(map[int]bool)
	for _, num := range arr {
		set[num] = true
	}
	return set
}

// Convert set to slice
func setToSlice(set map[int]bool) []int {
	result := []int{}
	for k := range set {
		result = append(result, k)
	}
	return result
}

