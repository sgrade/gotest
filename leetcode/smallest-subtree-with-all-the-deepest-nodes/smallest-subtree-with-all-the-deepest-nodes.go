// 865. Smallest Subtree with all the Deepest Nodes
// https://leetcode.com/problems/smallest-subtree-with-all-the-deepest-nodes/

package smallestsubtreewithallthedeepestnodes

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	return getDeepest(root).node
}

type deepest struct {
	depth int
	node  *TreeNode
}

func getDeepest(node *TreeNode) deepest {
	if node == nil {
		return deepest{0, nil}
	}
	left := getDeepest(node.Left)
	right := getDeepest(node.Right)
	if left.depth > right.depth {
		return deepest{left.depth + 1, left.node}
	}
	if right.depth > left.depth {
		return deepest{right.depth + 1, right.node}
	}
	return deepest{left.depth + 1, node}
}
