/*
 * @lc app=leetcode id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
 */

package code

import (
	"reflect"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "1",
			args: args{
				head: &ListNode{
					1,
					&ListNode{
						1,
						&ListNode{
							2,
							&ListNode{
								3,
								&ListNode{
									3,
									nil,
								},
							},
						},
					},
				},
			},
			want: &ListNode{
				1,
				&ListNode{
					2,
					&ListNode{
						3,
						nil,
					},
				},
			},
		},
		{
			name: "1",
			args: args{
				head: &ListNode{
					1,
					&ListNode{
						1,
						&ListNode{
							1,
							nil,
						},
					},
				},
			},
			want: &ListNode{
				1,
				nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
