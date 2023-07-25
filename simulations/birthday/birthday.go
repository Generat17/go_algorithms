package birthday

import "math/rand"

// https://en.wikipedia.org/wiki/Birthday_problem

// Парадокс дней рождения - утверждение, состоящее в том, что в группе, состоящей из 23 или более человек,
// вероятность совпадения дней рождения (число и месяц) хотя бы у двух людей превышает 50 %. Например,
// если в классе 23 ученика или более, то более вероятно то, что у какой-то пары одноклассников дни
// рождения придутся на один день, чем то, что у каждого будет свой неповторимый день рождения.

// simulateBirthdayMatches возвращает true, если выбран тот же номер
// дважды с помощью генератора случайных чисел, выбираем число между
// 0 и 365 для определенной группы людей.
func simulateBirthdayMatches(numOfPeople int) bool {
	const daysInYear = 365

	seen := make(map[int]bool)
	for i := 0; i < numOfPeople; i++ {
		day := rand.Intn(daysInYear)
		if seen[day] {
			return true
		}
		seen[day] = true
	}

	return false
}

// SimulateBirthdays возвращает вроятность совпадения дней рождения
func SimulateBirthdays(numOfPeople, runs int) float64 {
	same := 0
	for i := 0; i < runs; i++ {
		if simulateBirthdayMatches(numOfPeople) {
			same++
		}
	}

	return float64(same) / float64(runs)
}
