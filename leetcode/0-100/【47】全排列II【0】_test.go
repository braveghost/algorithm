package __100

import (
	"reflect"
	"testing"
)

func TestPermuteUnique(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test-1",
			args: args{[]int{1,1,3}},
			want: [][]int{{1, 1, 3}, {1, 3, 1}, {3, 1, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PermuteUnique(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Permute() = %v, want %v", got, tt.want)
			}
		})
	}
}
