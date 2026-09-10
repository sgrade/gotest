// 2265. Count Nodes Equal to Average of Subtree
// https://leetcode.com/problems/count-nodes-equal-to-average-of-subtree/

package countnodesequaltoaverageofsubtree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfSubtree(root *TreeNode) int {
	_, _, count := postOrder(root)
	return count
}

// postOrder returns the sum of the subtree at node, the number of nodes in it,
// and how many of those nodes equal the average of their own subtree.
func postOrder(node *TreeNode) (sum, nodes, count int) {
	if node == nil {
		return 0, 0, 0
	}

	leftSum, leftNodes, leftCount := postOrder(node.Left)
	rightSum, rightNodes, rightCount := postOrder(node.Right)

	sum = node.Val + leftSum + rightSum
	nodes = 1 + leftNodes + rightNodes
	count = leftCount + rightCount
	// Values are non-negative, so integer division truncates toward the floor.
	if sum/nodes == node.Val {
		count++
	}
	return sum, nodes, count
}
