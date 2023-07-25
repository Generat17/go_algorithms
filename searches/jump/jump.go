package jump

import (
	"math"
)

func JumpSearch(sortedList []int, target int) int {
	var index int

	jump := int(math.Sqrt(float64(len(sortedList))))

	if len(sortedList) <= 0 {
		return -1
	}

loop:
	for index < len(sortedList) {
		switch {

		case sortedList[index] == target:
			return index

		case sortedList[index] > target:
			break loop

		default:
			index = index + jump
		}
	}

	previous := index - jump

	if index > len(sortedList)-1 {
		index = len(sortedList) - 1
	}

	for sortedList[index] >= target {
		switch {

		case sortedList[index] == target:
			return index

		case index == previous:
			return -1

		default:
			index--
		}
	}

	return -1
}
