/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
 */

package code

import (
	"reflect"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{
				&TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:   9,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   20,
						Left:  &TreeNode{Val: 15, Left: nil, Right: nil},
						Right: &TreeNode{Val: 7, Left: nil, Right: nil},
					},
				},
			},
			want: [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			name: "2",
			args: args{
				&TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:   9,
						Left:  &TreeNode{Val: 14, Left: nil, Right: nil},
						Right: &TreeNode{Val: 6, Left: nil, Right: nil},
					},
					Right: &TreeNode{
						Val:   20,
						Left:  &TreeNode{Val: 15, Left: nil, Right: nil},
						Right: &TreeNode{Val: 7, Left: nil, Right: nil},
					},
				},
			},
			want: [][]int{{3}, {9, 20}, {14, 6, 15, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
