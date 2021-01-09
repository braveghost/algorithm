package __100

import (
	"fmt"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test-1",
			args: args{head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  3,
						Next: &ListNode{
							Val:  4,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SwapPairs(tt.args.head)
			fmt.Println(1,got)
			fmt.Println(1,got.Next)
			fmt.Println(1,got.Next.Next)
			fmt.Println(1,got.Next.Next.Next)
			fmt.Println(1,got.Next.Next.Next.Next)
		})
	}
}

func TestSwapPairs_Recursive(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test-1",
			args: args{head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  3,
						Next: &ListNode{
							Val:  4,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SwapPairs_Recursive(tt.args.head)
			fmt.Println(1,got)
			fmt.Println(1,got.Next)
			fmt.Println(1,got.Next.Next)
			fmt.Println(1,got.Next.Next.Next)
			fmt.Println(1,got.Next.Next.Next.Next)
		})
	}
}
