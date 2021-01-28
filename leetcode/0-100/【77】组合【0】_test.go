package __100

import (
	"testing"
)

func TestCombine(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test-1",
			args: args{7,4},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(Combine(tt.args.n, tt.args.k))
			//if got := Combine(tt.args.n, tt.args.k); !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("Combine() = %v, want %v", got, tt.want)
			//}
		})
	}
}

func TestCombine_Tag(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test-1",
			args: args{4,2},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(Combine_Tag(tt.args.n, tt.args.k))
			//if got := Combine(tt.args.n, tt.args.k); !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("Combine() = %v, want %v", got, tt.want)
			//}
		})
	}
}
