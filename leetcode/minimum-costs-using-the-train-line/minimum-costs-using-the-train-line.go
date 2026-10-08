// 2361. Minimum Costs Using the Train Line
// https://leetcode.com/problems/minimum-costs-using-the-train-line/

package minimumcostsusingthetrainline

// minimumCosts uses DP over two states: the cheapest cost to reach the
// current stop ending on the regular line, and ending on the express line.
// Switching regular -> express costs expressCost; express -> regular is free.
func minimumCosts(regular []int, express []int, expressCost int) []int64 {
	costs := make([]int64, len(regular))

	// Start at stop 0 on the regular line; being on express requires paying to switch.
	regularCost, expressLineCost := int64(0), int64(expressCost)

	for i := range regular {
		regularCost, expressLineCost =
			min(regularCost, expressLineCost)+int64(regular[i]),
			min(regularCost+int64(expressCost), expressLineCost)+int64(express[i])
		costs[i] = min(regularCost, expressLineCost)
	}

	return costs
}
