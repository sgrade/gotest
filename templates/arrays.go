package templates

// Two Pointers - Opposite Ends
func twoPointers(arr []int) {
	left, right := 0, len(arr)-1
	for left < right {
		// process arr[left] and arr[right]
		left++
		right--
	}
}

// Two Pointers - Same Direction (Fast & Slow)
func fastSlow(arr []int) {
	slow := 0
	for fast := 0; fast < len(arr); fast++ {
		// condition to move slow
		slow++
	}
}

// Sliding Window - Fixed Size
func slidingWindowFixed(arr []int, k int) {
	for i := 0; i <= len(arr)-k; i++ {
		// process window arr[i:i+k]
	}
}

// Sliding Window - Variable Size
func slidingWindowVariable(arr []int) {
	left := 0
	for right := 0; right < len(arr); right++ {
		// expand window: add arr[right]
		for /* condition violated */ false {
			// shrink window: remove arr[left]
			left++
		}
		// update result with valid window
	}
}

