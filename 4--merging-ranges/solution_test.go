package solution

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSolutions(t *testing.T) {
	cases := []struct {
		description               string
		integers                  [][2]int
		wantHighestProductOfThree [][2]int
	}{
		{
			"example",
			[][2]int{{0, 1}, {3, 5}, {4, 8}, {10, 12}, {9, 10}},
			[][2]int{{0, 1}, {3, 8}, {9, 12}},
		},
		{
			"out of order",
			[][2]int{{0, 1}, {3, 5}, {10, 12}, {4, 8}, {9, 10}},
			[][2]int{{0, 1}, {3, 8}, {9, 12}},
		},
		{
			"side-by-side meetings",
			[][2]int{{1, 2}, {2, 3}},
			[][2]int{{1, 3}},
		},
		{
			"second starts later but ends before first",
			[][2]int{{1, 5}, {2, 3}},
			[][2]int{{1, 5}},
		},
		{
			"many merged to one",
			[][2]int{{1, 10}, {2, 6}, {3, 5}, {7, 9}},
			[][2]int{{1, 10}},
		},
	}
	solutions := []func([][2]int) [][2]int{
		Solution0,
		Solution1,
		Solution2,
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%s %v", c.description, c.integers), func(t *testing.T) {
			for si, s := range solutions {
				t.Run(fmt.Sprintf("Solution%d", si), func(t *testing.T) {
					if g, w := s(c.integers), c.wantHighestProductOfThree; reflect.DeepEqual(g, w) {
						t.Logf("got %d", g)
					} else {
						t.Errorf("got %d, want %d", g, w)
					}
				})
			}
		})
	}
}
