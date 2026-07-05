package savethecow

import (
    "errors"
)

type Game struct{
    word    []rune
	guessed map[rune]bool
	current []rune
	tries   int
}

func NewGame(word string) *Game {
	w := []rune(word)
    cur := make([]rune, len(w))
    for i := range cur {
        cur[i] = '_'
    }

    return &Game{
        word: w,
        guessed: make(map[rune]bool),
        current: cur,
        tries: 0,
    }
}

func (g *Game) Guess(r rune) error {
    // if g.RemainingGuesses() == 0 {
    //     return nil
    // }
    if g.State() == "Lose" {
        return errors.New("cannot guess after the game is lost")
    }
    if g.State() == "Win" {
        return errors.New("cannot guess after the game is won")
    }
    if !(('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z')) {
        return errors.New("Incorrect letter for guessing!")
    }
    
    already := g.guessed[r]
    g.guessed[r] = true

    found := false
    for i, ch := range g.word {
        if ch == r {
            g.current[i] = r
            found = true
        }
    }
    if !found || already {
        if g.tries < 10 {
    		g.tries++
    	} 
    }
    return nil
}

func (g *Game) MaskedWord() string {
	return string(g.current)
}

func (g *Game) RemainingGuesses() int {
    res := 9 - g.tries
    if res < 0 {
        return 0
    }
    return res
}

func (g *Game) State() string {
    // if g.RemainingGuesses() > 0 {
    //     if g.MaskedWord() != string(g.word) {
    //         return "Ongoing"
    //     } 
    //     return "Win"
    // }
    // if g.RemainingGuesses() == 0 {
    //     if g.MaskedWord() != string(g.word) {
    //     	return "Lose"
    //     } 
    // 	return "Win"
    // }
    // return "Lose"
    // Win first
	complete := true
	for i := range g.word {
		if g.current[i] != g.word[i] {
			complete = false
			break
		}
	}
	if complete {
		return "Win"
	}
	if g.tries == 10 {
		return "Lose"
	}
	return "Ongoing"
}
