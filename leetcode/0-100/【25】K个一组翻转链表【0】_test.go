package __100

import (
	"fmt"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test-1",
			args: args{
				head: 	&ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  3,
							Next: &ListNode{
								Val:  4,
								Next: nil,
							},
						},
					},
				},
				k:    1,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseKGroup(tt.args.head, tt.args.k)
			fmt.Println(got)
			fmt.Println(got.Next)
			fmt.Println(got.Next.Next)
			fmt.Println(got.Next.Next.Next)
			fmt.Println(got.Next.Next.Next.Next)
		})
	}
}