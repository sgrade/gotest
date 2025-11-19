package templates

// Hash Map - Frequency Counter
func frequencyCounter(arr []int) {
	freq := make(map[int]int)
	for _, num := range arr {
		freq[num]++
	}
}

// Hash Set - Existence Check
func hashSet(arr []int) {
	seen := make(map[int]bool)
	for _, num := range arr {
		if seen[num] {
			// duplicate found
		}
		seen[num] = true
	}
}

// Hash Map - Two Sum Pattern
func twoSumPattern(arr []int, target int) {
	seen := make(map[int]int) // value -> index
	for i, num := range arr {
		complement := target - num
		if _, exists := seen[complement]; exists {
			// found pair
		}
		seen[num] = i
	}
}

