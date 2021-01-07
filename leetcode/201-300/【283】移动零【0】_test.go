package _01_300

import "testing"

func TestMoveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test-1",
			args: args{nums: []int{0,1,0,3,12}},
		},
		{
			name: "test-2",
			args: args{nums: []int{0,0,0,3,0}},
		},
		{
			name: "test-3",
			args: args{nums: []int{2,1}},
		},		{
			name: "test-4",
			args: args{nums: []int{1,0,2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MoveZeroes(tt.args.nums)
		})
		t.Log(tt.args.nums)
	}
}
