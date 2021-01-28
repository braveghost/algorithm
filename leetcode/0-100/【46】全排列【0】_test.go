package __100

import (
	"testing"
)

func TestPermute(t *testing.T) {
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
			args: args{[]int{1,2,3}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(Permute(tt.args.nums))
			//if got := Permute(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("Permute() = %v, want %v", got, tt.want)
			//}
		})
	}
}

func TestPermute_Swap(t *testing.T) {
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
			args: args{[]int{1,2,3}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(Permute_Swap(tt.args.nums))
			//if got := Permute(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("Permute() = %v, want %v", got, tt.want)
			//}
		})
	}
}
