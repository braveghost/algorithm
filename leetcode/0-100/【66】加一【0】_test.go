package __100

import (
	"testing"
)

func TestPlusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test-1",
			args: args{digits: []int{1,2,3}},
		},
		{
			name: "test-1",
			args: args{digits: []int{1,2,9}},
		},
		{
			name: "test-3",
			args: args{digits: []int{9,9,9}},
		},		{
			name: "test-4",
			args: args{digits: []int{9}},
		},	{
			name: "test-5",
			args: args{digits: []int{0}},
		},{
			name: "test-6",
			args: args{digits: []int{3}},
		},{
			name: "test-7",
			args: args{digits: []int{9,9}},
		},{
			name: "test-7",
			args: args{digits: []int{0,9}},
		},{
			name: "test-7",
			args: args{digits: []int{0,2}},
		},{
			name: "test-7",
			args: args{digits: []int{0,0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(PlusOne(tt.args.digits))
		})
	}
}
