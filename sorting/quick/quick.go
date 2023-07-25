package quick

type Num interface {
	int | int8 | float64
}

func partition[T Num](arr []T, low int, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++

			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func QuickSort[T Num](arr []T, low int, high int) {
	if low < high {
		pi := partition[T](arr, low, high)

		QuickSort(arr, low, pi-1)
		QuickSort(arr, pi+1, high)
	}
}

func IdiomaticQuickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivotIndex := len(a) / 2

	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	IdiomaticQuickSort(a[:left])
	IdiomaticQuickSort(a[left+1:])

	return a
}
