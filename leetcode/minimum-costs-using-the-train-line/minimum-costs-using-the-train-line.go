// 2361. Minimum Costs Using the Train Line
// https://leetcode.com/problems/minimum-costs-using-the-train-line/

package minimumcostsusingthetrainline

func minimumCosts(regular []int, express []int, expressCost int) []int64 {
	numStops := len(regular)
	costs := make([]int64, numStops)

	totalRegularCost := int64(0)
	totalExpressCost := int64(expressCost)

	for stop := 0; stop < numStops; stop++ {
		curRegularCost := min(totalRegularCost, totalExpressCost) + int64(regular[stop])
		curExpressCost := min(totalRegularCost+int64(expressCost), totalExpressCost) + int64(express[stop])

		totalRegularCost = curRegularCost
		totalExpressCost = curExpressCost
		costs[stop] = min(totalRegularCost, totalExpressCost)
	}

	return costs
}
