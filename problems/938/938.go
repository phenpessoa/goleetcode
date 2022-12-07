package main

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func (t *TreeNode) insert(v int) {
	if t == nil {
		return
	}
	if v <= t.Val {
		if t.Left == nil {
			t.Left = &TreeNode{nil, nil, v}
		} else {
			t.Left.insert(v)
		}
	} else {
		if t.Right == nil {
			t.Right = &TreeNode{nil, nil, v}
		} else {
			t.Right.insert(v)
		}
	}
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	var sum int
	if low <= root.Val && root.Val <= high {
		sum += root.Val
	}

	if root.Val >= low {
		sum += rangeSumBST(root.Left, low, high)
	}

	if root.Val <= high {
		sum += rangeSumBST(root.Right, low, high)
	}

	return sum
}
