package insertion

func InsertionSort(randomList []int) []int {

	for i := 1; i < len(randomList); i++ {
		for j := i; j > 0 && randomList[j-1] > randomList[j]; j-- {
			randomList[j-1], randomList[j] = randomList[j], randomList[j-1]
		}
	}

	return randomList
}
