package interpolation

func InterpolationSearchIterative(sortedList []int, target int) int {
	var leftIdx int
	rightIdx := len(sortedList) - 1

	for leftIdx <= rightIdx &&
		target >= sortedList[leftIdx] &&
		target <= sortedList[rightIdx] &&
		len(sortedList) > 0 {

		a := int(float64(rightIdx-leftIdx) / float64(sortedList[rightIdx]-sortedList[leftIdx]))
		b := target - sortedList[leftIdx]
		positionIdx := leftIdx + a*b

		value := sortedList[positionIdx]

		switch {

		case value == target:
			return positionIdx

		case value > target:
			rightIdx = positionIdx - 1

		case value < target:
			leftIdx = positionIdx - 1
		}
	}

	return -1
}

func InterpolationSearchRecursive(sortedList []int, target int, leftIdx int, rightIdx int) int {

	if leftIdx <= rightIdx &&
		target >= sortedList[leftIdx] &&
		target <= sortedList[rightIdx] &&
		len(sortedList) > 0 {

		a := int(float64(rightIdx-leftIdx) / float64(sortedList[rightIdx]-sortedList[leftIdx]))
		b := target - sortedList[leftIdx]
		positionIdx := leftIdx + a*b

		value := sortedList[positionIdx]

		switch {

		case value == target:
			return positionIdx

		case value > target:
			return InterpolationSearchRecursive(sortedList, target, leftIdx, positionIdx-1)

		case value < target:
			return InterpolationSearchRecursive(sortedList, target, positionIdx+1, rightIdx)
		}

	}

	return -1
}
