package highscores

import (
    "sort"
)

type HighScores struct {
    scores []int
}

// NewHighScores returns a new HighScores object.
func NewHighScores(scores []int) *HighScores {
	return &HighScores{scores: scores}
}

// Scores returns all the scores.
func (s *HighScores) Scores() []int {
	return s.scores
}

// Latest returns the latest (last) score.
func (s *HighScores) Latest() int {
    return s.scores[len(s.scores) - 1]
}

// PersonalBest returns the best (highest) score.
func (s *HighScores) PersonalBest() int {
	max := 0
	for _, score := range s.scores {
		if score > max {
			max = score
		}
	}
	return max
}

// TopThree returns the top three scores.
func (s *HighScores) TopThree() []int {
	temp := make([]int, len(s.scores))
	copy(temp, s.scores)
	
	sort.Slice(temp, func(i, j int) bool {
		return temp[i] > temp[j]
	})

	n := 3
	if len(temp) < 3 {
		n = len(temp)
	}
	return temp[:n]
}
