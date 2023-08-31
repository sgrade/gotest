// 637. Average of Levels in Binary Tree
// https://leetcode.com/problems/average-of-levels-in-binary-tree/

package averageoflevelsinbinarytree

import llq "github.com/emirpasic/gods/queues/linkedlistqueue"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	var ans []float64

	q := llq.New()
	q.Enqueue(root)
	for !q.Empty() {
		sum := 0
		cnt := q.Size()
		temp := llq.New()
		for !q.Empty() {
			var x, _ interface{} = q.Dequeue()
			node := x.(*TreeNode)
			sum += node.Val
			if node.Left != nil {
				temp.Enqueue(node.Left)
			}
			if node.Right != nil {
				temp.Enqueue(node.Right)
			}
		}
		ans = append(ans, float64(sum)/float64(cnt))
		q = temp
	}

	return ans
}
