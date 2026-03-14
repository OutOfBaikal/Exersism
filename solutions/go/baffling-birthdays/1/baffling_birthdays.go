package bafflingbirthdays

import (
    "math/rand"
    "time"
    "fmt"
)

func SharedBirthday(dates []time.Time) bool {
	seen := make(map[string]bool)
	for _, d := range dates {
		key := fmt.Sprintf("%d-%d", d.Month(), d.Day())
		if seen[key] {
			return true
		}
		seen[key] = true
	}
	return false
}

func RandomBirthdates(size int) []time.Time {
	year := 2023 
	dates := make([]time.Time, size)
	
	for i := 0; i < size; i++ {
		daysToAdd := rand.Intn(365)
		start := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
		dates[i] = start.AddDate(0, 0, daysToAdd)
	}
	return dates
}

func EstimatedProbability(size int) float64 {
	const simulations = 10000
	successCount := 0
	
	// Инициализируем рандом один раз перед циклом
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < simulations; i++ {
		if SharedBirthday(RandomBirthdates(size)) {
			successCount++
		}
	}

	return float64(successCount) / float64(simulations) * 100
}

func generateRandomDate(start, end time.Time) time.Time {
    diff := end.Unix() - start.Unix()
    randomDuration := time.Duration(rand.Int63n(int64(diff)))

    return start.Add(randomDuration)
}
