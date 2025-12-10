// 3577. Count the Number of Computer Unlocking Permutations
// https://leetcode.com/problems/count-the-number-of-computer-unlocking-permutations/

package countthenumberofcomputerunlockingpermutations

// Based on Editorial's Approach: Brain Teaser
func countPermutations(complexity []int) int {
	n := len(complexity)
	for i := 1; i < n; i++ {
		if complexity[i] <= complexity[0] {
			return 0
		}
	}

	const MOD = 1e9 + 7
	ans := 1
	for i := 2; i < n; i++ {
		ans = ans * i % MOD
	}
	return ans
}
