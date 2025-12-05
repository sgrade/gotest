// 1214. Two Sum BSTs
// https://leetcode.com/problems/two-sum-bsts/

package twosumbsts

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func twoSumBSTs(root1 *TreeNode, root2 *TreeNode, target int) bool {
	var nodes1, nodes2 []int
	dfs(root1, &nodes1)
	dfs(root2, &nodes2)

	i1, i2 := 0, len(nodes2)-1
	for i1 < len(nodes1) && i2 >= 0 {
		switch sum := nodes1[i1] + nodes2[i2]; {
		case sum == target:
			return true
		case sum < target:
			i1++
		default:
			i2--
		}
	}
	return false
}

func dfs(node *TreeNode, nodes *[]int) {
	if node == nil {
		return
	}
	dfs(node.Left, nodes)
	*nodes = append(*nodes, node.Val)
	dfs(node.Right, nodes)
}
