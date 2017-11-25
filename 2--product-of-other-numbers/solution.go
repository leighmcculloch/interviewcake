// Package solution contains solutions to the problem described at https://www.interviewcake.com/question/product-of-other-numbers.
package solution

// Solution0 solves the problem in O(n2) time and O(n) space.
func Solution0(integers []int) []int {
	count := len(integers)
	if count < 2 {
		return nil
	}
	multiples := make([]int, count)

	for i := 0; i < count; i++ {
		multiples[i] = 1
		for j := 0; j < count; j++ {
			if i == j {
				continue
			}
			multiples[i] *= integers[j]
		}
	}

	return multiples
}

// Solution1 solves the problem in O(n) time and O(n) space.
func Solution1(integers []int) []int {
	count := len(integers)
	if count < 2 {
		return nil
	}
	multiples := make([]int, count)

	{
		multiple := 1
		for i := 0; i < count; i++ {
			multiples[i] = multiple
			multiple *= integers[i]
		}
	}

	{
		multiple := 1
		for i := count - 1; i >= 0; i-- {
			multiples[i] *= multiple
			multiple *= integers[i]
		}
	}

	return multiples
}
