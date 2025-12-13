package couponcodevalidator

import (
	"regexp"
	"slices"
)

// Compile regex once at package level for performance
var alphanumericRegex = regexp.MustCompile("^[a-zA-Z0-9_]+$")

// Define valid business lines as a constant
var validBusinessLines = []string{"electronics", "grocery", "pharmacy", "restaurant"}

func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
	// Initialize map with nil slices (idiomatic Go)
	validCoupons := map[string][]string{
		"electronics": nil,
		"grocery":     nil,
		"pharmacy":    nil,
		"restaurant":  nil,
	}

	for i := range code {
		// Early returns for invalid conditions
		if !isActive[i] {
			continue
		}
		if _, ok := validCoupons[businessLine[i]]; !ok {
			continue
		}
		if len(code[i]) == 0 {
			continue
		}

		// Use pre-compiled regex
		if !alphanumericRegex.MatchString(code[i]) {
			continue
		}
		validCoupons[businessLine[i]] = append(validCoupons[businessLine[i]], code[i])
	}

	// Pre-allocate result slice with estimated capacity
	var ans []string
	for _, line := range validBusinessLines {
		slices.Sort(validCoupons[line])
		ans = append(ans, validCoupons[line]...)
	}
	return ans
}
