package __100

import "testing"

func TestIsValid_Violence(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test-1",
			args: args{s: "{}[]"},
			want: true,
		},		{
			name: "test-2",
			args: args{s: ""},
			want: false,
		},	{
			name: "test-3",
			args: args{s: "{{}"},
			want: false,
		},	{
			name: "test-4",
			args: args{s: "{{[]}}"},
			want: true,
		},{
			name: "test-5",
			args: args{s: "{{[}}"},
			want: false,
		},{
			name: "test-6",
			args: args{s: "{"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValid_Violence(tt.args.s); got != tt.want {
				t.Errorf("IsValid_Violence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test-1",
			args: args{s: "{}[]"},
			want: true,
		},		{
			name: "test-2",
			args: args{s: ""},
			want: false,
		},	{
			name: "test-3",
			args: args{s: "{{}"},
			want: false,
		},	{
			name: "test-4",
			args: args{s: "{{[]}}"},
			want: true,
		},{
			name: "test-5",
			args: args{s: "{{[}}"},
			want: false,
		},{
			name: "test-6",
			args: args{s: "{{"},
			want: false,
		},{
			name: "test-7",
			args: args{s: "(()("},
			want: false,
		},{
			name: "test-9",
			args: args{s: "("},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValid(tt.args.s); got != tt.want {
				t.Errorf("IsValid_Violence() = %v, want %v", got, tt.want)
			}
		})
	}
}
