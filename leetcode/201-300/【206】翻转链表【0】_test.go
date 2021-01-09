package _01_300

import (
	"testing"
)

func TestReverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	l1 := &ListNode{
		Val:  45,
		Next: nil,
	}
	l1.Next = l1
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test-1",
			args: args{head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			}},
		},
		{
			name: "test-2",
			args: args{head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  3,
						Next: l1,
					},
				},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseList(tt.args.head)
			t.Log(got)
			t.Log(got.Next)
			t.Log(got.Next.Next)
			t.Log(got.Next.Next.Next)
		})
	}
}

func TestReverseList_Recursive(t *testing.T) {
	type args struct {
		head *ListNode
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "test-1",
			args: args{head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseList_Recursive(tt.args.head)
			t.Log(got)
			t.Log(got.Next)
			t.Log(got.Next.Next)
			t.Log(got.Next.Next.Next)
		})
	}
}
