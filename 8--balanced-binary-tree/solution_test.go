package solution

import (
	"fmt"
	"testing"
)

func TestSolutions(t *testing.T) {
	cases := []struct {
		description       string
		tree              *BinaryTreeNode
		wantSuperbalanced bool
	}{
		{
			"root",
			&BinaryTreeNode{},
			true,
		},
		{
			"balanced depth diff 0",
			&BinaryTreeNode{
				value: 1,
				left: &BinaryTreeNode{
					value: 0,
				},
				right: &BinaryTreeNode{
					value: 2,
				},
			},
			true,
		},
		{
			"balanced depth diff 1",
			&BinaryTreeNode{
				value: 1,
				left: &BinaryTreeNode{
					value: 0,
				},
				right: &BinaryTreeNode{
					value: 2,
					right: &BinaryTreeNode{
						value: 3,
					},
				},
			},
			true,
		},
		{
			"unbalanced depth diff 2",
			&BinaryTreeNode{
				left: &BinaryTreeNode{
					value: 0,
				},
				value: 1,
				right: &BinaryTreeNode{
					value: 2,
					right: &BinaryTreeNode{
						left: &BinaryTreeNode{
							value: 3,
						},
						value: 4,
					},
				},
			},
			false,
		},
		{
			"unbalanced root with balanced subtrees",
			&BinaryTreeNode{ // 1
				left: &BinaryTreeNode{ // 2
					left: &BinaryTreeNode{}, // 3
					right: &BinaryTreeNode{ // 3
						left:  &BinaryTreeNode{}, // 4
						right: &BinaryTreeNode{}, // 4
					},
				},
				right: &BinaryTreeNode{ // 2
					left: &BinaryTreeNode{ // 3
						left:  &BinaryTreeNode{}, // 4
						right: &BinaryTreeNode{}, // 4
					},
					right: &BinaryTreeNode{ // 3
						left: &BinaryTreeNode{}, // 4
						right: &BinaryTreeNode{ // 4
							left: &BinaryTreeNode{}, // 5
						},
					},
				},
			},
			false,
		},
	}
	solutions := []func(*BinaryTreeNode) bool{
		Solution0,
		Solution1,
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%s %v", c.description, c.tree), func(t *testing.T) {
			for si, s := range solutions {
				t.Run(fmt.Sprintf("Solution%d", si), func(t *testing.T) {
					if g, w := s(c.tree), c.wantSuperbalanced; g == w {
						t.Logf("got %v", g)
					} else {
						t.Errorf("got %v, want %v", g, w)
					}
				})
			}
		})
	}
}
