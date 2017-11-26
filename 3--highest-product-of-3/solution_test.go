package solution

import (
	"fmt"
	"testing"
)

func TestSolutions(t *testing.T) {
	cases := []struct {
		description               string
		integers                  []int
		wantHighestProductOfThree int
	}{
		{"example", []int{1, 7, 3, 4}, 84},
		{"zeroes", []int{1, 0, 3, 4}, 12},
		{"negatives", []int{-10, -10, 1, 3, 2}, 300},
		{"negatives", []int{-10, -3, 10, 1, 3, 2}, 300},
		{"negatives", []int{-10, 10, 1, 3, 2}, 60},
	}
	solutions := []func([]int) int{
		Solution0,
		Solution1,
		Solution2,
		Solution3,
		Solution4,
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%s %v", c.description, c.integers), func(t *testing.T) {
			for si, s := range solutions {
				t.Run(fmt.Sprintf("Solution%d", si), func(t *testing.T) {
					if g, w := s(c.integers), c.wantHighestProductOfThree; g == w {
						t.Logf("got %d", g)
					} else {
						t.Errorf("got %d, want %d", g, w)
					}
				})
			}
		})
	}
}
