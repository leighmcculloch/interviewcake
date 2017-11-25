package solution

import (
	"fmt"
	"testing"
)

func TestSolutions(t *testing.T) {
	cases := []struct {
		description          string
		stockPricesYesterday []int
		wantMaxProfit        int
	}{
		{"example", []int{10, 7, 5, 8, 11, 9}, 6},
		{"largest first", []int{20, 7, 5, 8, 11, 9}, 6},
		{"negative price", []int{20, -10, 7, 5, 8, 11, 9}, 21},
		{"no profit", []int{20, 20, 20, 20}, 0},
		{"loss", []int{20, 10, 7, 5}, -2},
	}
	solutions := []func([]int) int{
		Solution0,
		Solution1,
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%s %v", c.description, c.stockPricesYesterday), func(t *testing.T) {
			for si, s := range solutions {
				t.Run(fmt.Sprintf("solution %d", si), func(t *testing.T) {
					if g, w := s(c.stockPricesYesterday), c.wantMaxProfit; g == w {
						t.Logf("got %d", g)
					} else {
						t.Errorf("got %d, want %d", g, w)
					}
				})
			}
		})
	}
}
