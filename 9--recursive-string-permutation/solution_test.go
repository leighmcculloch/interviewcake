package recursivestringpermutation

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerm(t *testing.T) {
	testCases := []struct {
		in  string
		out []string
	}{
		{
			in:  "",
			out: []string{""},
		},
		{
			in:  "A",
			out: []string{"A"},
		},
		{
			in:  "AB",
			out: []string{"AB", "BA"},
		},
		{
			in:  "AA",
			out: []string{"AA"},
		},
		{
			in:  "ABC",
			out: []string{"ABC", "ACB", "BAC", "BCA", "CBA", "CAB"},
		},
		{
			in:  "ABA",
			out: []string{"ABA", "AAB", "BAA"},
		},
		{
			in: "PRACTICE",
			out: func() []string {
				lines := []string{}
				f, err := os.Open("testdata/practice.txt")
				if err != nil {
					panic(err)
				}
				s := bufio.NewScanner(f)
				for s.Scan() {
					lines = append(lines, s.Text())
				}
				return lines
			}(),
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			out := Perm(testCase.in)
			t.Log(out)
			assert.ElementsMatch(t, testCase.out, out)
		})
	}
}
