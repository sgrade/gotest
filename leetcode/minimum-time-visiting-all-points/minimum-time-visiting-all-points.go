// 1266. Minimum Time Visiting All Points
// https://leetcode.com/problems/minimum-time-visiting-all-points/

package minimumtimevisitingallpoints

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minTimeToVisitAllPoints(points [][]int) int {
	ans := 0
	for i := 1; i < len(points); i++ {
		xDiff := abs(points[i][0] - points[i-1][0])
		yDiff := abs(points[i][1] - points[i-1][1])
		ans += max(xDiff, yDiff)
	}
	return ans
}
