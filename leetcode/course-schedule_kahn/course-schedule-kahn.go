// 207. Course Schedule
// https://leetcode.com/problems/course-schedule/

package courseschedulekahn

func canFinish(numCourses int, prerequisites [][]int) bool {
	adj := make([][]int, numCourses)
	inDegree := make([]int, numCourses)
	for _, p := range prerequisites {
		course, dependency := p[0], p[1]
		adj[dependency] = append(adj[dependency], course)
		inDegree[course]++
	}

	coursesVisited := 0
	queue := make([]int, 0)
	for i := range numCourses {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		coursesVisited++
		for _, neighbor := range adj[course] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}
	return coursesVisited == numCourses
}
