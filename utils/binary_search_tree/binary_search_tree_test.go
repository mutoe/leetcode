package binary_search_tree

import (
	"fmt"
	"testing"

	. "../tree_node"
)

func TestBST_Add(t *testing.T) {
	type fields struct {
		root *TreeNode
		size int
	}
	type args struct {
		num int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		{
			name: "test1",
			fields: fields{
				root: &TreeNode{Val: 2},
				size: 1,
			},
			args: args{1},
			want: []byte{'l'},
		},
		{
			name: "test2",
			fields: fields{
				root: &TreeNode{
					Val:  5,
					Left: &TreeNode{Val: 2},
				},
				size: 2,
			},
			args: args{3},
			want: []byte{'l', 'r'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectSize := tt.fields.size + 1
			b := &BinarySearchTree{
				root: tt.fields.root,
				size: tt.fields.size,
			}
			b.Add(tt.args.num)
			target := b.root
			for _, condition := range tt.want {
				if target == nil {
					t.Errorf("Target not match")
					break
				}
				if condition == 'l' {
					target = target.Left
				} else {
					target = target.Right
				}
			}
			if target == nil {
				t.Errorf("Target not match")
				return
			}
			if target.Val != tt.args.num {
				t.Errorf("Failed")
			}
			if b.size != expectSize {
				t.Errorf("Size not match: expect %v, but got %v", expectSize, b.size)
			}
		})
	}
}

func TestBST_Order(t *testing.T) {
	b := &BinarySearchTree{}
	elements := []int{5, 3, 1, 4, 6, 2, 9, 8, 7}
	for _, element := range elements {
		b.Add(element)
	}
	fmt.Println("Level order")
	b.LevelOrder()
	fmt.Println("Previous order")
	b.PrevOrder()
	fmt.Println("Middle order")
	b.MidOrder()
	fmt.Println("Next order")
	b.NextOrder()
}
