package templates

import "strings"

// Two Pointers - Palindrome Check
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// Reverse String
func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

// Anagram Check - Frequency Counter
func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	freq := make(map[rune]int)
	for _, ch := range s {
		freq[ch]++
	}
	for _, ch := range t {
		freq[ch]--
		if freq[ch] < 0 {
			return false
		}
	}
	return true
}

// Substring Pattern - Sliding Window
func substringPattern(s string) {
	left := 0
	charMap := make(map[byte]int)
	for right := 0; right < len(s); right++ {
		charMap[s[right]]++
		for /* condition */ false {
			charMap[s[left]]--
			left++
		}
		// process window s[left:right+1]
	}
}

// String Builder Pattern
func stringBuilder() string {
	var sb strings.Builder
	sb.WriteString("hello")
	sb.WriteString(" world")
	return sb.String()
}

