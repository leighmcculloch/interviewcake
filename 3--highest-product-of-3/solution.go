// Package solution contains solutions to the problem described at https://www.interviewcake.com/question/highest-product-of-3.
package solution

import (
	"sort"
)

// Solution0 solves the problem in O(n3) time and O(1) space.
func Solution0(integers []int) int {
	count := len(integers)
	if count < 3 {
		return 0
	}

	max := integers[0] * integers[1] * integers[2]

	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			for k := 0; k < count; k++ {
				if i == j || i == k || j == k {
					continue
				}
				mult := integers[i] * integers[j] * integers[k]
				if mult > max {
					max = mult
				}
			}
		}
	}

	return max
}

// Solution1 solves the problem in O(nlogn) time and O(1) space.
func Solution1(integers []int) int {
	count := len(integers)
	if count < 3 {
		return 0
	}

	sort.Ints(integers)

	max1 := integers[count-1] * integers[count-2] * integers[count-3]
	max2 := integers[count-1] * integers[0] * integers[1]
	max := max1
	if max2 > max {
		max = max2
	}
	return max
}

// Solution2 solves the problem in O(n) time and O(1) space.
func Solution2(ints []int) int {
	if len(ints) < 3 {
		return 0
	}

	h1 := -1
	l1 := -1
	for i, n := range ints {
		if h1 == -1 || n > ints[h1] {
			h1 = i
		}
		if l1 == -1 || n < ints[l1] {
			l1 = i
		}
	}

	h2 := -1
	l2 := -1
	for i, n := range ints {
		if i != h1 && (h2 == -1 || n > ints[h2]) {
			h2 = i
		}
		if i != l1 && (l2 == -1 || n < ints[l2]) {
			l2 = i
		}
	}

	h3 := -1
	l3 := -1
	for i, n := range ints {
		if i != h1 && i != h2 && (h3 == -1 || n > ints[h3]) {
			h3 = i
		}
		if i != l1 && i != l2 && (l3 == -1 || n > ints[l3]) {
			l3 = i
		}
	}

	hMul := ints[h1] * ints[h2] * ints[h3]
	lMul := ints[l1] * ints[l2] * ints[l3]
	maxMul := hMul
	if lMul > maxMul {
		maxMul = lMul
	}

	return maxMul
}

// Solution3 solves the problem in O(n) time and O(1) space.
func Solution3(ints []int) int {
	if len(ints) < 3 {
		return 0
	}

	highest := ints[0]
	if ints[1] > highest {
		highest = ints[1]
	}
	lowest := ints[0]
	if ints[1] < lowest {
		lowest = ints[1]
	}

	highestOf2 := ints[0] * ints[1]
	lowestOf2 := ints[0] * ints[1]

	highestOf3 := ints[0] * ints[1] * ints[2]

	for i := 2; i < len(ints); i++ {
		n := ints[i]

		if x := n * highestOf2; x > highestOf3 {
			highestOf3 = x
		}
		if x := n * lowestOf2; x > highestOf3 {
			highestOf3 = x
		}

		if x := n * highest; x > highestOf2 {
			highestOf2 = x
		}
		if x := n * lowest; x > highestOf2 {
			highestOf2 = x
		}

		if x := n * highest; x < lowestOf2 {
			lowestOf2 = x
		}
		if x := n * lowest; x < lowestOf2 {
			lowestOf2 = x
		}

		if n > highest {
			highest = n
		}

		if n < lowest {
			lowest = n
		}
	}

	return highestOf3
}

// Solution4 solves the problem in O(n) time and O(1) space.
func Solution4(ints []int) int {
	if len(ints) < 3 {
		return 0
	}

	max := func(ints ...int) int {
		m := ints[0]
		for _, n := range ints {
			if n > m {
				m = n
			}
		}
		return m
	}

	min := func(ints ...int) int {
		m := ints[0]
		for _, n := range ints {
			if n < m {
				m = n
			}
		}
		return m
	}

	highest := max(ints[:2]...)
	lowest := min(ints[:2]...)

	highestOf2 := ints[0] * ints[1]
	lowestOf2 := ints[0] * ints[1]

	highestOf3 := ints[0] * ints[1] * ints[2]

	for i := 2; i < len(ints); i++ {
		n := ints[i]

		highestOf3 = max(
			highestOf3,
			n*highestOf2,
			n*lowestOf2,
		)

		highestOf2 = max(
			highestOf2,
			n*highest,
			n*lowest,
		)

		lowestOf2 = min(
			highestOf2,
			n*highest,
			n*lowest,
		)

		highest = max(
			highest,
			n,
		)

		lowest = min(
			lowest,
			n,
		)
	}

	return highestOf3
}
