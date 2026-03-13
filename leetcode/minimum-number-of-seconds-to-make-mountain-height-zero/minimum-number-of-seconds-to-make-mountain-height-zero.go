// 3296. Minimum Number of Seconds to Make Mountain Height Zero
// https://leetcode.com/problems/minimum-number-of-seconds-to-make-mountain-height-zero/
//
// Binary search on total seconds. For a candidate time, each worker
// with cost t can reduce height by k units if t·k·(k+1)/2 ≤ time.
// Sum all workers' reductions; feasible if the sum ≥ mountainHeight.

package minimumnumberofsecondstomakemountainheightzero

import "math"

func minNumberOfSeconds(mountainHeight int, workerTimes []int) int64 {
	maxTime := 0
	for _, t := range workerTimes {
		maxTime = max(maxTime, t)
	}

	h := int64(mountainHeight)
	lo, hi := int64(1), int64(maxTime)*h*(h+1)/2
	var ans int64

	for lo <= hi {
		seconds := lo + (hi-lo)/2
		var totalReduction int64

		for _, t := range workerTimes {
			budget := seconds / int64(t)
			// Largest k where k*(k+1)/2 ≤ budget, via quadratic formula.
			reduction := int64(math.Sqrt(float64(2 * budget)))
			if reduction*(reduction+1)/2 > budget {
				reduction--
			}
			totalReduction += reduction
		}

		if totalReduction >= h {
			ans = seconds
			hi = seconds - 1
		} else {
			lo = seconds + 1
		}
	}

	return ans
}
