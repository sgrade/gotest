// 2265. Count Nodes Equal to Average of Subtree
// https://leetcode.com/problems/count-nodes-equal-to-average-of-subtree/

package countnodesequaltoaverageofsubtree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// postOrder returns the sum of the subtree at node, the number of nodes in it,
// and how many of those nodes equal the average of their own subtree.
func postOrder(node *TreeNode) (int, int, int) {
	totalSum := node.Val
	nodes := 1
	count := 0
	if node.Left != nil {
		childSum, childNodes, childCount := postOrder(node.Left)
		totalSum += childSum
		nodes += childNodes
		count += childCount
	}
	if node.Right != nil {
		childSum, childNodes, childCount := postOrder(node.Right)
		totalSum += childSum
		nodes += childNodes
		count += childCount
	}
	average := totalSum / nodes
	if average == node.Val {
		count++
	}
	return totalSum, nodes, count
}

func averageOfSubtree(root *TreeNode) int {
	_, _, ans := postOrder(root)
	return ans
}
