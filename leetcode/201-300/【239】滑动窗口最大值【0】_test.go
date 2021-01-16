package _01_300

import (
	"reflect"
	"testing"
)

func TestMaxSlidingWindow_Violence(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test-1",
			args: args{
				nums: []int{1, 3, 2, 6, 32, 5, 7, 22, 3, 5, 6, 77, 1},
				k:    1,
			},
			want: []int{1, 3, 2, 6, 32, 5, 7, 22, 3, 5, 6, 77, 1},
		}, {
			name: "test-2",
			args: args{
				nums: []int{1, 3, 2, 6, 32, 5, 7, 22, 3, 5, 6, 77, 1},
				k:    2,
			},
			want: []int{3, 3, 6, 32, 32, 7, 22, 22, 5, 6, 77, 77},
		}, {
			name: "test-3",
			args: args{
				nums: []int{1, 3, 2, 6, 32, 5, 7, 22, 3, 5, 6, 77, 1},
				k:    3,
			},
			want: []int{3, 6, 32, 32, 32, 22, 22, 22, 6, 77, 77},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSlidingWindow_Violence(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxSlidingWindow_Violence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxSlidingWindow(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test-1",
			args: args{
				nums: []int{1, 3, 2, 6, 32, 5, 7, 22, 3, 5, 6, 77, 1},
				k:    1,
			},
			want: []int{1, 3, 2, 6, 32, 5, 7, 22, 3, 5, 6, 77, 1},
		}, {
			name: "test-2",
			args: args{
				nums: []int{1, 3, 2, 6, 32, 5, 7, 22, 3, 5, 6, 77, 1},
				k:    2,
			},
			want: []int{3, 3, 6, 32, 32, 7, 22, 22, 5, 6, 77, 77},
		}, {
			name: "test-3",
			args: args{
				nums: []int{1, 3, 2, 6, 32, 5, 7, 22, 3, 5, 6, 77, 1},
				k:    3,
			},
			want: []int{3, 6, 32, 32, 32, 22, 22, 22, 6, 77,77},
		},{
			name: "test-4",
			args: args{
				nums: []int{1,3,1,2,0,5},
				k:    3,
			},
			want: []int{3,3,2,5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSlidingWindow(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxSlidingWindow_Violence() = %v, want %v", got, tt.want)
			}
		})
	}
}
