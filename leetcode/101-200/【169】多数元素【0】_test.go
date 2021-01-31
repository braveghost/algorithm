package _01_200

import "testing"

func TestMajorityElementBoyerMoore(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test-1",
			args: args{[]int{1,1,1,1,2,2,2,3,4,5,1,1,1,1}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MajorityElementBoyerMoore(tt.args.nums); got != tt.want {
				t.Errorf("MajorityElementBoyerMoore() = %v, want %v", got, tt.want)
			}
		})
	}
}
