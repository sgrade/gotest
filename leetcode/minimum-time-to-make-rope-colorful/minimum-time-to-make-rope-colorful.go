package minimumtimetomakeropecolorful

func minCost(colors string, neededTime []int) int {
	n := len(colors)
	if colors[n-1] == 'a' {
		colors = colors + string('b')
	} else {
		colors = colors + string('a')
	}
	neededTime = append(neededTime, 0)
	totalTime, curColorTime, maxTimeForCurColor := 0, neededTime[0], neededTime[0]
	for i := range n {
		if colors[i] == colors[i+1] {
			curColorTime += neededTime[i+1]
			maxTimeForCurColor = max(maxTimeForCurColor, neededTime[i+1])
		} else {
			totalTime += curColorTime - maxTimeForCurColor
			curColorTime = neededTime[i+1]
			maxTimeForCurColor = neededTime[i+1]
		}
	}
	return totalTime
}
