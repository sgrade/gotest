// 207. Course Schedule
// https://leetcode.com/problems/course-schedule/

package courseschedule

func canFinish(numCourses int, prerequisites [][]int) bool {
	adj := make([][]int, numCourses)
	for _, p := range prerequisites {
		course := p[0]
		prerequisite := p[1]
		adj[prerequisite] = append(adj[prerequisite], course)
	}

	visited := make([]bool, numCourses)
	onPath := make([]bool, numCourses)

	var hasCycle func(course int) bool
	hasCycle = func(course int) bool {
		if onPath[course] {
			return true
		}
		if visited[course] {
			return false
		}

		onPath[course] = true
		visited[course] = true

		for _, nextCourse := range adj[course] {
			if hasCycle(nextCourse) {
				return true
			}
		}
		onPath[course] = false
		return false
	}

	for course := 0; course < numCourses; course++ {
		if hasCycle(course) {
			return false
		}
	}
	return true
}
