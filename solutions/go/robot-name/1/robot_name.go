package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Robot представляет робота с именем
type Robot struct {
	name string
}

// globalNames хранит все уже использованные имена
var globalNames = make(map[string]bool)

// init засевает генератор случайных чисел один раз при запуске программы
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Name возвращает имя робота. Если имя не назначено, генерирует уникальное.
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	name, err := generateUniqueName()
	if err != nil {
		return "", err
	}

	r.name = name
	return name, nil
}

// Reset сбрасывает имя робота, заставляя сгенерировать новое при следующем вызове Name()
func (r *Robot) Reset() {
	r.name = ""
}

// generateUniqueName генерирует уникальное имя или возвращает ошибку, если все имена заняты
func generateUniqueName() (string, error) {
	const maxAttempts = 1000 // Ограничиваем попытки, чтобы не зациклиться
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"

	for i := 0; i < maxAttempts; i++ {
		name := fmt.Sprintf("%c%c%c%c%c",
			letters[rand.Intn(len(letters))],
			letters[rand.Intn(len(letters))],
			numbers[rand.Intn(len(numbers))],
			numbers[rand.Intn(len(numbers))],
			numbers[rand.Intn(len(numbers))],
		)

		if !globalNames[name] {
			globalNames[name] = true
			return name, nil
		}
	}

	// Если после maxAttempts имя не найдено, вероятно, все имена заняты
	return "", errors.New("all possible names are exhausted")
}
