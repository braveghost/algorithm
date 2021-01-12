package _01_200

import (
	"reflect"
	"testing"
)

func TestDetectCycle_Violence(t *testing.T) {
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
		want *ListNode
	}{
		{
			name: "test-1",
			args: args{head: l1},
			want: l1,
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
			want: nil,
		},		{
			name: "test-3",
			args: args{head: &ListNode{
				Val: 1,
			}},
			want: nil,
		},
		{
			name: "test-2",
			args: args{head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
				},
			}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DetectCycle_Violence(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DetectCycle_Violence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDetectCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}

	l1 := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  12,
			Next: nil,
		},
	}
	l1.Next.Next = l1
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test-1",
			args: args{head: l1},
			want: l1,
		},		{
			name: "test-5",
			args: args{head: &ListNode{
				Val:  5,
				Next: l1,
			}},
			want: l1,
		},
		{
			name: "test-2",
			args: args{head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  3,
					},
				},
			}},
			want: nil,
		},		{
			name: "test-3",
			args: args{head: &ListNode{
				Val: 1,
			}},
			want: nil,
		},
		{
			name: "test-2",
			args: args{head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
				},
			}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DetectCycle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DetectCycle_Violence() = %v, want %v", got, tt.want)
			}
		})
	}
}
