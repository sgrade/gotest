// 549. Binary Tree Longest Consecutive Sequence II
// https://leetcode.com/problems/binary-tree-longest-consecutive-sequence-ii/

package binarytreelongestconsecutivesequenceii

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type pathLengths struct {
	increasing int
	decreasing int
}

// longestConsecutive returns the length of the longest consecutive elements in the binary tree.
// It considers paths that are either strictly increasing or strictly decreasing by 1.
func longestConsecutive(root *TreeNode) int {
	maxLen := 0
	getLongestPath(root, &maxLen)
	return maxLen
}

func getLongestPath(node *TreeNode, maxLen *int) pathLengths {
	if node == nil {
		return pathLengths{0, 0}
	}

	incrLen, decrLen := 1, 1

	if node.Left != nil {
		left := getLongestPath(node.Left, maxLen)
		if node.Val == node.Left.Val-1 {
			incrLen = left.increasing + 1
		} else if node.Val == node.Left.Val+1 {
			decrLen = left.decreasing + 1
		}
	}

	if node.Right != nil {
		right := getLongestPath(node.Right, maxLen)
		if node.Val == node.Right.Val-1 {
			incrLen = max(incrLen, right.increasing+1)
		} else if node.Val == node.Right.Val+1 {
			decrLen = max(decrLen, right.decreasing+1)
		}
	}

	*maxLen = max(*maxLen, incrLen+decrLen-1)
	return pathLengths{incrLen, decrLen}
}
