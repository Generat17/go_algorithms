package monty

import (
	"math/rand"
)

// https://en.wikipedia.org/wiki/Monty_Hall_problem

func stayWinsGame() bool {
	carDoor := rand.Intn(3)
	playerDoor := rand.Intn(3)

	return carDoor == playerDoor
}

func fraction(a, b int) float64 {
	return float64(a) / float64(b)
}

func Simulation(n int) (float64, float64) {
	stayWin, switchWin := 0, 0
	for i := 0; i < n; i++ {
		if stayWinsGame() {
			stayWin++
		} else {
			switchWin++
		}
	}

	return fraction(stayWin, n), fraction(switchWin, n)
}
