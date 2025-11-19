package templates

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DFS - Preorder (Root -> Left -> Right)
func preorder(root *TreeNode) {
	if root == nil {
		return
	}
	// process root
	preorder(root.Left)
	preorder(root.Right)
}

// DFS - Inorder (Left -> Root -> Right)
func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	// process root
	inorder(root.Right)
}

// DFS - Postorder (Left -> Right -> Root)
func postorder(root *TreeNode) {
	if root == nil {
		return
	}
	postorder(root.Left)
	postorder(root.Right)
	// process root
}

// BFS - Level Order
func levelOrder(root *TreeNode) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			// process node
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
}

