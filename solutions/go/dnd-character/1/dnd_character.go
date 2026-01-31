package dndcharacter

import (
    "math"
    "math/rand"
    "time"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return int(math.Floor((float64(score) - 10.0) / 2.0))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	rand.Seed(time.Now().UnixNano())
    data := []int{}
    res := 0
    for i := 0; i < 4; i++ {
        data = append(data, 1 + rand.Intn(6))
        res += data[i]
    }
    min := 7
    for i := 0; i < 4; i++ {
    	if min > data[i] {
            min = data[i]
        }
    }
    return res - min
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	var ch Character
    ch.Strength = Ability()
	ch.Dexterity = Ability()
	ch.Constitution = Ability()
	ch.Intelligence = Ability()
	ch.Wisdom = Ability()
	ch.Charisma = Ability()
	ch.Hitpoints = 10 + Modifier(ch.Constitution)
    return ch
}
