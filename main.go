package main

import (
	"algo/simulations/birthday"
	"algo/sorting/bubble"
	"fmt"
	"math/rand"
	"time"
)

func compArr(arr []int, n int) bool {
	check := true

	var trueArr = make([]int, n)
	for i := 0; i < n; i++ {
		trueArr[i] = i
	}

	for i := 0; i < n; i++ {
		if arr[i] != trueArr[i] {
			check = false
		}
	}

	// if arrays compare then true, else false
	return check
}

func testSort() {
	rand.Seed(time.Now().Unix())

	const n = 600  // длинна массива
	const m = 4000 // кол-во повторений

	start := time.Now() // начало измерения

	flag := true
	for i := 0; i < m; i++ {
		randArr := rand.Perm(n)
		//quick.QuickSort[int](randArr, 0, len(randArr)-1)
		bubble.BubbleSort(randArr)
		if !compArr(randArr, n) {
			flag = false
		}
	}

	duration := time.Since(start)
	fmt.Println(duration)

	fmt.Println(flag)
}

func main() {
	fmt.Println(birthday.SimulateBirthdays(23, 100000))
}
