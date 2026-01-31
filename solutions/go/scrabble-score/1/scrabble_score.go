package scrabble

import "strings"

func Score(word string) int {
    res := 0
    word = strings.ToLower(word)
    data := make(map[rune]int, 0)
    data['a'] = 1
    data['b'] = 3
    data['c'] = 3
    data['d'] = 2
    data['e'] = 1
    data['f'] = 4
    data['g'] = 2
    data['h'] = 4
    data['i'] = 1
    data['j'] = 8
    data['k'] = 5
    data['m'] = 3
    data['l'] = 1
    data['n'] = 1
    data['o'] = 1
    data['p'] = 3
    data['q'] = 10
    data['r'] = 1
    data['s'] = 1
    data['t'] = 1
    data['u'] = 1
    data['v'] = 4
    data['w'] = 4
    data['x'] = 8
    data['y'] = 4
    data['z'] = 10
    for _, x := range(word) {
        res += data[x]
    }
    return res
}
