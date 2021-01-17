package __100

import (
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "test-1",
			args: args{[]string{"abc","cba","ac","ca"}},
			want: [][]string{{"abc","cba"},{"ac","ca"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GroupAnagrams(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
