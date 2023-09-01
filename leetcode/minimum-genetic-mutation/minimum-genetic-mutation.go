// 433. Minimum Genetic Mutation
// https://leetcode.com/problems/minimum-genetic-mutation/

package minimumgeneticmutation

import (
	llq "github.com/emirpasic/gods/queues/linkedlistqueue"
	"github.com/emirpasic/gods/sets/hashset"
)

func minMutation(startGene string, endGene string, bank []string) int {
	q := llq.New()
	visited := hashset.New()
	q.Enqueue(startGene)
	visited.Add(startGene)

	steps := 0
	for !q.Empty() {
		count := q.Size()
		for i := 0; i < count; i++ {
			var node_int, _ interface{} = q.Dequeue()
			node := node_int.(string)
			if node == endGene {
				return steps
			}

			for _, ch := range "ACGT" {
				for j := 0; j < len(node); j++ {
					candidate_arr := []rune(node)
					candidate_arr[j] = ch
					candidate := string(candidate_arr)
					if !visited.Contains(candidate) {
						for _, gene := range bank {
							if candidate == gene {
								q.Enqueue(candidate)
								visited.Add(candidate)
							}
						}
					}
				}
			}
		}
		steps++
	}

	return -1
}
