package __100

import "testing"

func TestTrap_Violence(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test-1",
			args: args{[]int{0,1,0,2,1,0,1,3,2,1,2,1}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Trap_Violence(tt.args.height); got != tt.want {
				t.Errorf("Trap_Violence() = %v, want %v", got, tt.want)
			}
		})
	}
}