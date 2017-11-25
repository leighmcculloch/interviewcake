package solution

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSolutions(t *testing.T) {
	cases := []struct {
		description                        string
		integers                           []int
		wantProductsOfAllIntsExceptAtIndex []int
	}{
		{"example", []int{1, 7, 3, 4}, []int{84, 12, 28, 21}},
		{"zeroes", []int{1, 0, 3, 4}, []int{0, 12, 0, 0}},
		{"one", []int{12}, nil},
	}
	solutions := []func([]int) []int{
		Solution0,
		Solution1,
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%s %v", c.description, c.integers), func(t *testing.T) {
			for si, s := range solutions {
				t.Run(fmt.Sprintf("Solution%d", si), func(t *testing.T) {
					if g, w := s(c.integers), c.wantProductsOfAllIntsExceptAtIndex; reflect.DeepEqual(g, w) {
						t.Logf("got %d", g)
					} else {
						t.Errorf("got %d, want %d", g, w)
					}
				})
			}
		})
	}
}
