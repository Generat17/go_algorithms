package selection

func selectionSort(randomList []int) []int {
	for i := 0; i < len(randomList)-1; i++ {
		smallestIdx := i

		for j := i + 1; j < len(randomList); j++ {
			if randomList[j] < randomList[smallestIdx] {
				smallestIdx = j
			}
		}

		randomList[i], randomList[smallestIdx] = randomList[smallestIdx], randomList[i]
	}
	return randomList
}
