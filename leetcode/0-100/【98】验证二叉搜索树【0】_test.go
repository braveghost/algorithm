package __100

import "testing"

func Test_IsValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test-1",
			args: args{&TreeNode{
				Val:   15,
				Left:  &TreeNode{
					Val:   6,
					Left:  &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   17,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val:   18,
					Left:  &TreeNode{
						Val:   5,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   20,
						Left:  nil,
						Right: nil,
					},
				},
			}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_IsValidBST_InorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		//{
		//	name: "test-1",
		//	args: args{&TreeNode{
		//		Val:   15,
		//		Left:  &TreeNode{
		//			Val:   6,
		//			Left:  &TreeNode{
		//				Val:   4,
		//				Left:  nil,
		//				Right: nil,
		//			},
		//			Right: &TreeNode{
		//				Val:   17,
		//				Left:  nil,
		//				Right: nil,
		//			},
		//		},
		//		Right: &TreeNode{
		//			Val:   18,
		//			Left:  &TreeNode{
		//				Val:   5,
		//				Left:  nil,
		//				Right: nil,
		//			},
		//			Right: &TreeNode{
		//				Val:   20,
		//				Left:  nil,
		//				Right: nil,
		//			},
		//		},
		//	}},
		//	want: false,
		//},
		{
			name: "test-1",
			args: args{&TreeNode{
				Val:   2,
				Left:  &TreeNode{
					Val:   1,

				},
				Right: &TreeNode{
					Val:   3,

				},
			}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidBST_InorderTraversal2(tt.args.root); got != tt.want {
			//if got := IsValidBST_InorderTraversal(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
