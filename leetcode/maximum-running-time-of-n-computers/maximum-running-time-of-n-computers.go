// 2141. Maximum Running Time of N Computers
// https://leetcode.com/problems/maximum-running-time-of-n-computers/

package maximumrunningtimeofncomputers

import "sort"

func maxRunTime(n int, batteries []int) int64 {
	// Sort batteries in ascending order
	sort.Ints(batteries)

	// Calculate sum of the smallest batteries (those that will be shared)
	shared := int64(0)
	for i := 0; i < len(batteries)-n; i++ {
		shared += int64(batteries[i])
	}

	// Get the n largest batteries
	nLargest := batteries[len(batteries)-n:]

	// Try to distribute the shared batteries
	for i := 0; i < n-1; i++ {
		need := int64(nLargest[i+1] - nLargest[i])
		sharedPerComputer := shared / int64(i+1)

		if sharedPerComputer < need {
			return int64(nLargest[i]) + sharedPerComputer
		}
		shared -= int64(i+1) * need
	}

	// Distribute remaining shared batteries evenly among all n computers
	return int64(nLargest[n-1]) + shared/int64(n)
}
