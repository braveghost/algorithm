package __100

import "testing"

func TestMaxArea(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test-1",
			args: args{[]int{4,3,2,1,4}},
		},		{
			name: "test-2",
			args: args{[]int{2,3,4,5,18,17,6}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(MaxArea(tt.args.height))
		})
	}
}
func TestMaxArea_Violence(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test-1",
			args: args{[]int{4,3,2,1,4}},
		},		{
			name: "test-2",
			args: args{[]int{2,3,4,5,18,17,6}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(MaxArea_Violence(tt.args.height))
		})
	}
}
