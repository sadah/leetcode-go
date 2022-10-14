/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
 */

package code

import "testing"

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				&TreeNode{
					5,
					&TreeNode{1, nil, nil},
					&TreeNode{
						4,
						&TreeNode{3, nil, nil},
						&TreeNode{6, nil, nil}}},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
