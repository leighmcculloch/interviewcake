package solution

import (
	"fmt"
	"testing"
)

func TestSolutions(t *testing.T) {
	cases := []struct {
		description                     string
		amount                          int
		denominations                   []int
		wantPossibleWaysToMakeTheAmount int
	}{
		{"example", 4, []int{1, 2, 3}, 4},
	}
	solutions := []func(int, []int) int{
		Solution0,
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%s %v %d", c.description, c.amount, c.denominations), func(t *testing.T) {
			for si, s := range solutions {
				t.Run(fmt.Sprintf("Solution%d", si), func(t *testing.T) {
					if g, w := s(c.amount, c.denominations), c.wantPossibleWaysToMakeTheAmount; g == w {
						t.Logf("got %d", g)
					} else {
						t.Errorf("got %d, want %d", g, w)
					}
				})
			}
		})
	}
}
