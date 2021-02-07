package _01_200

import "testing"

func TestNumIslands_DFS(t *testing.T) {
	type args struct {
		grid [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test-1",
			args: args{[][]byte{
			{'1','1','1','1','0'},
			{'1','1','0','1','0'},
			{'1','1','0','0','0'},
			{'0','0','0','0','0'},
			}},
			want: 1,
		},{
			name: "test-2",
			args: args{[][]byte{
			{'1','1','1','1','0'},
			{'1','1','0','1','0'},
			{'1','1','0','0','0'},
			{'0','0','0','0','1'},
			}},
			want: 2,
		},{
			name: "test-3",
			args: args{[][]byte{
			{'1','1','1','1','1'},
			{'1','1','0','1','1'},
			{'1','1','0','1','1'},
			{'0','0','0','0','1'},
			}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumIslands_DFS(tt.args.grid); got != tt.want {
				t.Errorf("NumIslands_DFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumIslands_BFS(t *testing.T) {
	type args struct {
		grid [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test-1",
			args: args{[][]byte{
			{'1','1','1','1','0'},
			{'1','1','0','1','0'},
			{'1','1','0','0','0'},
			{'0','0','0','0','0'},
			}},
			want: 1,
		},{
			name: "test-2",
			args: args{[][]byte{
			{'1','1','1','1','0'},
			{'1','1','0','1','0'},
			{'1','1','0','0','0'},
			{'0','0','0','0','1'},
			}},
			want: 2,
		},{
			name: "test-3",
			args: args{[][]byte{
			{'1','1','1','1','1'},
			{'1','1','0','1','1'},
			{'1','1','0','1','1'},
			{'0','0','0','0','1'},
			}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumIslands_BFS(tt.args.grid); got != tt.want {
				t.Errorf("NumIslands_DFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
