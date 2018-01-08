// Package solution contains solutions to the problem described at https://www.interviewcake.com/question/temperature-tracker.
package solution

type TempTracker interface {
	Insert(temp int)
	Max() int
	Min() int
	Mean() float32
	Mode() int
}

// Solution0 stores the temperatures using O(n) space.
type Solution0 struct {
	tempCounts map[int]int
	count      int
	max        int
	min        int
	sum        int
}

// Insert inserts the temperature in O(1) time and O(1) space.
func (s *Solution0) Insert(temp int) {
	if s.count == 0 || temp > s.max {
		s.max = temp
	}
	if s.count == 0 || temp < s.min {
		s.min = temp
	}
	s.sum += temp
	if s.tempCounts == nil {
		s.tempCounts = map[int]int{}
	}
	s.tempCounts[temp]++
	s.count++
}

// Max returns the maximum temperature in O(1) time and O(1) space.
func (s *Solution0) Max() int {
	return s.max
}

// Min returns the minimum temperature in O(1) time and O(1) space.
func (s *Solution0) Min() int {
	return s.min
}

// Mean returns the average temperature in O(1) time and O(1) space.
func (s *Solution0) Mean() float32 {
	if s.count == 0 {
		return 0
	}
	return float32(s.sum) / float32(s.count)
}

// Mode returns the most common temperature in O(n) and O(1) space.
func (s *Solution0) Mode() int {
	mostCommonTemp := 0
	mostCommonTempCount := 0
	for temp, count := range s.tempCounts {
		if count > mostCommonTempCount {
			mostCommonTemp = temp
			mostCommonTempCount = count
		}
	}
	return mostCommonTemp
}

// Solution1 stores the temperatures using O(1) space.
// Note: The temperatures are stored in a map, but since temperatures are
// limited to integers in the range of 0 to 110, the map is at worse O(111)
// space regardless of the number of temperatures, and the space is therefore
// O(1).
type Solution1 struct {
	tempCounts         map[int]int
	tempWithMostCounts int

	count int
	max   int
	min   int
	sum   int
}

// Insert inserts the temperature in O(1) time and O(1) space.
func (s *Solution1) Insert(temp int) {
	if s.count == 0 || temp > s.max {
		s.max = temp
	}

	if s.count == 0 || temp < s.min {
		s.min = temp
	}

	s.sum += temp

	if s.tempCounts == nil {
		s.tempCounts = map[int]int{}
	}

	s.tempCounts[temp]++
	if s.tempCounts[temp] > s.tempCounts[s.tempWithMostCounts] {
		s.tempWithMostCounts = temp
	}

	s.count++
}

// Max returns the maximum temperature in O(1) time and O(1) space.
func (s *Solution1) Max() int {
	return s.max
}

// Min returns the minimum temperature in O(1) time and O(1) space.
func (s *Solution1) Min() int {
	return s.min
}

// Mean returns the average temperature in O(1) time and O(1) space.
func (s *Solution1) Mean() float32 {
	if s.count == 0 {
		return 0
	}
	return float32(s.sum) / float32(s.count)
}

// Mode returns the most common temperature in O(1) and O(1) space.
func (s *Solution1) Mode() int {
	return s.tempWithMostCounts
}

// Solution2 stores the temperatures using O(1) space.
// Note: The temperatures counts are stored in a fixed size array, since
// temperatures are limited to integers in the range of 0 to 110.
type Solution2 struct {
	tempCounts         [111]int
	tempWithMostCounts int

	count int
	max   int
	min   int
	sum   int
}

// Insert inserts the temperature in O(1) time and O(1) space.
func (s *Solution2) Insert(temp int) {
	if s.count == 0 || temp > s.max {
		s.max = temp
	}

	if s.count == 0 || temp < s.min {
		s.min = temp
	}

	s.sum += temp

	s.tempCounts[temp]++
	if s.tempCounts[temp] > s.tempCounts[s.tempWithMostCounts] {
		s.tempWithMostCounts = temp
	}

	s.count++
}

// Max returns the maximum temperature in O(1) time and O(1) space.
func (s *Solution2) Max() int {
	return s.max
}

// Min returns the minimum temperature in O(1) time and O(1) space.
func (s *Solution2) Min() int {
	return s.min
}

// Mean returns the average temperature in O(1) time and O(1) space.
func (s *Solution2) Mean() float32 {
	if s.count == 0 {
		return 0
	}
	return float32(s.sum) / float32(s.count)
}

// Mode returns the most common temperature in O(1) and O(1) space.
func (s *Solution2) Mode() int {
	return s.tempWithMostCounts
}
