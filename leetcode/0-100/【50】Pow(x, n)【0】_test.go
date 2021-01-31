package __100

import (
	"fmt"
	"testing"
)

func TestMyPow(t *testing.T) {
	fmt.Println(10/2)
	fmt.Println(5/2)
	fmt.Println(4&1)
	fmt.Println(5&1)
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MyPow(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("MyPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
