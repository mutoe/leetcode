package tree_node

import (
	"reflect"
	"testing"
)

func TestTreeNodeToArray(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{&TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val:  4,
					Left: &TreeNode{Val: 9},
					Right: &TreeNode{
						Val:  8,
						Left: &TreeNode{Val: 4},
					},
				},
			}},
			want: []int{2, NULL, 4, 9, 8, NULL, NULL, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TreeNodeToArray(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TreeNodeToArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayToTreeNode(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "test1",
			args: args{[]int{2, NULL, 4, 9, 8, NULL, NULL, 4}},
			want: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val:  4,
					Left: &TreeNode{Val: 9},
					Right: &TreeNode{
						Val:  8,
						Left: &TreeNode{Val: 4},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ArrayToTreeNode(tt.args.arr)
			result := isSameTree(got, tt.want)
			if !result {
				t.Errorf("ArrayToTreeNode()")
			}
		})
	}
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	} else {
		return false
	}
}
