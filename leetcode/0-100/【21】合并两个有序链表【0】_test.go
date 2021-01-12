package __100

import (
	"fmt"
	"testing"
)

func TestMergeTwoLists_Iteration(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test-1",
			args: args{
				l1: &ListNode{
					Val:  0,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
				l2:&ListNode{
					Val:  1,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		},
		{
			name: "test-2",
			args: args{
				l1: &ListNode{
					Val:  0,
				},
				l2:&ListNode{
					Val:  1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeTwoLists_Iteration(tt.args.l1, tt.args.l2)
			fmt.Println(got)
			fmt.Println(got.Next)
			fmt.Println(got.Next.Next)
			fmt.Println(got.Next.Next.Next)
			fmt.Println(got.Next.Next.Next.Next)
		})
	}
}
