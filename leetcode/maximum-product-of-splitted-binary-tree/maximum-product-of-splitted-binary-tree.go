// 1339. Maximum Product of Splitted Binary Tree
// https://leetcode.com/problems/maximum-product-of-splitted-binary-tree/

package maximumproductofsplittedbinarytree

const mod = 1_000_000_007

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxProduct(root *TreeNode) int {
	var sums []int
	var getSum func(*TreeNode) int
	getSum = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		cur := node.Val + getSum(node.Left) + getSum(node.Right)
		sums = append(sums, cur)
		return cur
	}
	total := getSum(root)
	maxProd := int64(0)
	for _, sum := range sums {
		maxProd = max(maxProd, int64(sum)*int64(total-sum))
	}
	return int(maxProd % mod)
}
