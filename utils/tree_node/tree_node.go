package tree_node

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ArrayToTreeNode(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	root := &TreeNode{Val: arr[0]}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for i := 1; i < len(arr); i += 2 {
		cur := queue[0]
		queue = queue[1:]
		if arr[i] != 0 {
			cur.Left = &TreeNode{Val: arr[i]}
			queue = append(queue, cur.Left)
		}
		if i+1 < len(arr) && arr[i+1] != 0 {
			cur.Right = &TreeNode{Val: arr[i+1]}
			queue = append(queue, cur.Right)
		}
	}
	return root
}

func TreeNodeToArray(root *TreeNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	ret = append(ret, root.Val)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			continue
		}
		if node.Left == nil {
			ret = append(ret, 0)
		} else {
			ret = append(ret, node.Left.Val)
		}
		if node.Right == nil {
			ret = append(ret, 0)
		} else {
			ret = append(ret, node.Right.Val)
		}
		queue = append(queue, node.Left, node.Right)
	}
	i := len(ret) - 1
	for i >= 0 {
		if ret[i] != 0 {
			break
		}
		i--
	}
	return ret[:i+1]
}
