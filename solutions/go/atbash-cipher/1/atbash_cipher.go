package atbash

import "strings"

func Atbash(s string) string {
    s = strings.ToLower(s)
	res := []rune{}
    var a int
    a = 0
    for i := 0; i < len(s); i++ {
        if s[i] < 'a' || s[i] > 'z' {
            if s[i] >= '0' && s[i] <= '9' {
                res = append(res, rune(s[i]))
                a++
            }
            continue
        }
        if a == 5 {
            res = append(res, ' ')
            a = 0
        }
        res = append(res, rune(26 - (s[i] - 'a') + 'a' - 1))
        a++
    }
    return string(res)
}
