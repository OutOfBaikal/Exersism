package affinecipher

import (
    "errors"
    "strings"
    "unicode"
)

func Encode(text string, a, b int) (string, error) {
    if !isCoprime(a) {
		return "", errors.New("a and m must be coprime")
	}
    var res []rune
    count := 0
    for _, r := range strings.ToLower(text) {
        var char rune
        if unicode.IsLetter(r) {
        	char = rune(((a * int(r - 'a') + b) % 26) + 'a')
        } else if unicode.IsDigit(r) {
            char = r
        } else {
            continue
        }

        if count > 0 && count % 5 == 0 {
            res = append(res, ' ')
        }
        res = append(res, char)
        count++
    }

    return string(res), nil
}

func Decode(text string, a, b int) (string, error) {
	if !isCoprime(a) {
		return "", errors.New("a and m must be coprime")
	}

	inv := findInv(a)
	var res []rune
    
    for _, r := range text {
		if unicode.IsLetter(r) {
			y := int(r - 'a')
			val := (inv * (y - b)) % 26
			if val < 0 { val += 26 }
			res = append(res, rune(val+'a'))
		} else if unicode.IsDigit(r) {
			res = append(res, r)
		}
    }

    return string(res), nil
}

func isCoprime(a int) bool {
	coprimes := []int{1, 3, 5, 7, 9, 11, 15, 17, 19, 21, 23, 25}
	for _, v := range coprimes {
		if a == v { return true }
	}
	return false
}

func findInv(a int) int {
	for x := 1; x < 26; x++ {
		if (a * x) % 26 == 1 {
			return x
		}
	}
	return 1
}
