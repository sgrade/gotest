// 3623. Count Number of Trapezoids I
// https://leetcode.com/problems/count-number-of-trapezoids-i/

package countnumberoftrapezoidsi

func countTrapezoids(points [][]int) int {
	MOD := 1000000000 + 7
	pointsOnLevel := make(map[int]int)
	for _, point := range points {
		pointsOnLevel[point[1]]++
	}
	ans, totalEdges := 0, 0
	for _, cnt := range pointsOnLevel {
		edgesOnLevel := cnt * (cnt - 1) / 2
		ans += edgesOnLevel * totalEdges
		totalEdges += edgesOnLevel
	}
	return ans % MOD
}
