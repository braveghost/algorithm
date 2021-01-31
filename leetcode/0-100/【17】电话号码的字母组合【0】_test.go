package __100

import (
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	type args struct {
		digits string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test-1",
			args: args{"23"},
			want: []string{"ad","ae","af","bd","be","bf","cd","ce","cf"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LetterCombinations(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LetterCombinations() = %v, want %v", got, tt.want)
			}else {
				t.Log(got)
			}
		})
	}
}

func TestLetterCombinations_Iter(t *testing.T) {
	type args struct {
		digits string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test-1",
			args: args{"23"},
			want: []string{"ad","bd","cd","ae","be","ce","af","bf","cf"},
		},
		{
			name: "test-2",
			args: args{""},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LetterCombinations_Iter(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LetterCombinations_Iter() = %v, want %v", got, tt.want)
			}
		})
	}
}
