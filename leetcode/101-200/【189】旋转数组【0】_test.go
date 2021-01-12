package _01_200

import "testing"

func TestRotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test-1",
			args: args{nums: []int{1,2,3,4,5,6,7}, k: 1},
		},
		{
			name: "test-2",
			args: args{nums: []int{1,2,3,4,5,6,7}, k:3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Rotate(tt.args.nums, tt.args.k)
			t.Log(tt.args.nums)
		})
	}
}

func TestRotate_Segmented_Flip(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test-1",
			args: args{nums: []int{1,2,3,4,5,6,7}, k: 1},
		},
		{
			name: "test-2",
			args: args{nums: []int{1,2,3,4,5,6,7}, k:5},
		},		{
			name: "test-3",
			args: args{nums: []int{1,}, k:5},
		},	{
			name: "test-4",
			args: args{nums: []int{}, k:5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Rotate_Segmented_Flip(tt.args.nums, tt.args.k)
			t.Log(tt.args.nums)
		})
	}
}

func TestRotate_Rotate_2(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test-1",
			args: args{nums: []int{1,2,3,4,5,6,7}, k: 1},
		},
		{
			name: "test-2",
			args: args{nums: []int{1,2,3,4,5,6,7}, k:2},
		},		{
			name: "test-3",
			args: args{nums: []int{1,4,5,7,9}, k:3},
		},
		{
			name: "test-4",
			args: args{nums: []int{-1,-100,3,99}, k:2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Rotate_2(tt.args.nums, tt.args.k)
			t.Log(tt.args.nums)
		})
	}
}
