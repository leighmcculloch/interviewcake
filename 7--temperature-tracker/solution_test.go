package solution

import (
	"fmt"
	"testing"
)

func TestTempTracker(t *testing.T) {
	cases := []struct {
		description string
		temps       []int
		wantMax     int
		wantMin     int
		wantMean    float32
		wantMode    int
	}{
		{"zero", []int{}, 0, 0, 0, 0},
		{"one", []int{70}, 70, 70, 70, 70},
		{"multiple", []int{32, 70, 70}, 70, 32, 57.3333333333333, 70},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%s %v", c.description, c.temps), func(t *testing.T) {
			solutions := []TempTracker{
				&Solution0{},
				&Solution1{},
				&Solution2{},
			}

			for si, tempTracker := range solutions {
				t.Run(fmt.Sprintf("Solution%d", si), func(t *testing.T) {
					for _, temp := range c.temps {
						tempTracker.Insert(temp)
					}

					t.Run("Max", func(t *testing.T) {
						if g, w := tempTracker.Max(), c.wantMax; g == w {
							t.Logf("got %d", g)
						} else {
							t.Errorf("got %d, want %d", g, w)
						}
					})

					t.Run("Min", func(t *testing.T) {
						if g, w := tempTracker.Min(), c.wantMin; g == w {
							t.Logf("got %d", g)
						} else {
							t.Errorf("got %d, want %d", g, w)
						}
					})

					t.Run("Mean", func(t *testing.T) {
						if g, w := tempTracker.Mean(), c.wantMean; g == w {
							t.Logf("got %f", g)
						} else {
							t.Errorf("got %f, want %f", g, w)
						}
					})

					t.Run("Mode", func(t *testing.T) {
						if g, w := tempTracker.Mode(), c.wantMode; g == w {
							t.Logf("got %d", g)
						} else {
							t.Errorf("got %d, want %d", g, w)
						}
					})
				})
			}
		})
	}
}
