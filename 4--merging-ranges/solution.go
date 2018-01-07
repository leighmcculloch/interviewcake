// Package solution contains solutions to the problem described at https://www.interviewcake.com/question/merging-ranges.
package solution

import "sort"

type sortableMeetings [][2]int

func (m sortableMeetings) Len() int {
	return len(m)
}

func (m sortableMeetings) Less(i, j int) bool {
	return m[i][0] < m[j][0]
}

func (m sortableMeetings) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

// Solution0 solves the problem in O(nlogn) time and O(n) space.
func Solution0(meetings [][2]int) [][2]int {
	sortedMeetings := make([][2]int, len(meetings))
	copy(sortedMeetings, meetings)
	sort.Sort(sortableMeetings(sortedMeetings))

	merged := [][2]int{sortedMeetings[0]}

	for _, m := range sortedMeetings {
		lastIndex := len(merged) - 1
		last := merged[lastIndex]

		if m[0] <= last[1] {
			if m[1] > last[1] {
				last[1] = m[1]
			}
			merged[lastIndex] = last
		} else {
			merged = append(merged, m)
		}
	}
	return merged
}

// Solution1 solves the problem in O(n) time and O(n) space, assuming a reasonable upper bound.
func Solution1(meetings [][2]int) [][2]int {
	upperBound := 0
	for _, m := range meetings {
		if m[1] > upperBound {
			upperBound = m[1]
		}
	}

	timeSlots := make([]bool, upperBound)
	for _, m := range meetings {
		for i := m[0]; i < m[1]; i++ {
			timeSlots[i] = true
		}
	}

	merged := [][2]int{}
	for timeSlot, hasMeeting := range timeSlots {
		if !hasMeeting {
			continue
		}
		if len(merged) == 0 || timeSlot > merged[len(merged)-1][1] {
			merged = append(merged, [2]int{timeSlot, timeSlot + 1})
		} else {
			merged[len(merged)-1][1] = timeSlot + 1
		}
	}

	return merged
}

// Solution2 solves the problem in O(nlogn) time and O(n) space, in place.
func Solution2(meetings [][2]int) [][2]int {
	sort.Sort(sortableMeetings(meetings))

	lastIndex := 0

	for _, m := range meetings {
		last := meetings[lastIndex]
		if m[0] <= last[1] {
			if m[1] > last[1] {
				last[1] = m[1]
			}
			meetings[lastIndex] = last
		} else {
			lastIndex++
			meetings[lastIndex] = m
		}
	}

	return meetings[:lastIndex+1]
}
