package __100

import "testing"

func TestMerge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test-1",
			args: args{
				nums1: []int{1,2,7,0,0,0},
				m:     3,
				nums2: []int{2,5,6},
				n:     3,
			},
		},		{
			name: "test-2",
			args: args{
				nums1: []int{0,0,0,0,0,0,0},
				m:     0,
				nums2: []int{-50,-50,-2,1,1,3,4},
				n:    7,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Merge(tt.args.nums1, tt.args.m, tt.args.nums2,tt.args.n)
			t.Log(tt.args.nums1)
		})
	}
}

func TestMerge_Tail_Insert(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test-1",
			args: args{
				nums1: []int{1,2,7,0,0,0},
				m:     3,
				nums2: []int{2,5,6},
				n:     3,
			},
		},		{
			name: "test-2",
			args: args{
				nums1: []int{0,0,0,0,0,0,0},
				m:     0,
				nums2: []int{-50,-50,-2,1,1,3,4},
				n:    7,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Merge_Tail_Insert(tt.args.nums1, tt.args.m, tt.args.nums2,tt.args.n)
			t.Log(tt.args.nums1)
		})
	}
}
