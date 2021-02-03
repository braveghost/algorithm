package _01_500

import "testing"

func TestMinMutation(t *testing.T) {
	type args struct {
		start string
		end   string
		bank  []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test-1",
			args: args{
				start: "AACCGGTT",
				end:   "AACCGCTA",
				bank:  []string{"AACCGGTA","AACCGCTA","AAACGGTA"},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinMutation(tt.args.start, tt.args.end, tt.args.bank); got != tt.want {
				t.Errorf("MinMutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
