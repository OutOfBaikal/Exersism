package pangram

import "strings"

func IsPangram(input string) bool {
    input = strings.ToLower(input)
    data := make(map[rune]int, 0)
    data['a'] = 0
    data['b'] = 0
    data['c'] = 0
    data['d'] = 0
    data['e'] = 0
    data['f'] = 0
    data['g'] = 0
    data['h'] = 0
    data['i'] = 0
    data['j'] = 0
    data['k'] = 0
    data['m'] = 0
    data['l'] = 0
    data['n'] = 0
    data['o'] = 0
    data['p'] = 0
    data['q'] = 0
    data['r'] = 0
    data['s'] = 0
    data['t'] = 0
    data['u'] = 0
    data['v'] = 0
    data['w'] = 0
    data['x'] = 0
    data['y'] = 0
    data['z'] = 0
    for _, x := range(input) {
        data[x]++
    }
    for _, v := range(data) {
        if v == 0 {
            return false
        }
    }
	return true
}
