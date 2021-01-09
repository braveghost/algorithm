package _01_200

import (
	"testing"
)

func TestHasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	l1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	l1.Next = l1

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test-1",
			args: args{head: l1},
			want: true,
		},
		{
			name: "test-2",
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
			want: false,
		},		{
			name: "test-3",
			args: args{head: &ListNode{
				Val: 1,
			}},
			want: false,
		},
		{
			name: "test-2",
			args: args{head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
				},
			}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasCycle(tt.args.head); got != tt.want {
				t.Errorf("HasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
