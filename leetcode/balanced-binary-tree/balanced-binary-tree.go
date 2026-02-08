// 110. Balanced Binary Tree
// https://leetcode.com/problems/balanced-binary-tree/

package balancedbinarytree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	_, balanced := height(root)
	return balanced
}

func height(node *TreeNode) (int, bool) {
	if node == nil {
		return 0, true
	}

	leftHeight, leftBalanced := height(node.Left)
	if !leftBalanced {
		return 0, false
	}

	rightHeight, rightBalanced := height(node.Right)
	if !rightBalanced {
		return 0, false
	}

	diff := leftHeight - rightHeight
	if diff < 0 {
		diff = -diff
	}
	if diff > 1 {
		return 0, false
	}

	h := leftHeight
	if rightHeight > leftHeight {
		h = rightHeight
	}
	return h + 1, true
}
