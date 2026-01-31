package romannumerals

import "fmt"

func ToRomanNumeral(input int) (string, error) {
    res := []byte{}
    if input < 1 || input > 3999  {
        return string(res), fmt.Errorf("Invalid value")
    }
    
    for input >= 1000 {
        res = append(res, 'M')
        input -= 1000
    } 
    for input >= 900 {
        res = append(res, 'C')
        res = append(res, 'M')
        input -= 900
    } 
    for input >= 500 {
            res = append(res, 'D')
            input -= 500
    }
    for input >= 400 {
        res = append(res, 'C')
        res = append(res, 'D')
        input -= 400
    }
    for input >= 100 {
            res = append(res, 'C')
            input -= 100
    }
    for input >= 90 {
        res = append(res, 'X')
        res = append(res, 'C')
        input -= 90
    }
    for input >= 50 {
            res = append(res, 'L')
            input -= 50
    }
    for input >= 40 {
        res = append(res, 'X')
        res = append(res, 'L')
        input -= 40
    }
    for input >= 10 {
            res = append(res, 'X')
            input -= 10
    }
    for input >= 9 {
        res = append(res, 'I')
        res = append(res, 'X')
        input -= 9
    }
    for input >= 5 {
            res = append(res, 'V')
            input -= 5
    }
    for input >= 4 {
        res = append(res, 'I')
        res = append(res, 'V')
        input -= 4
    }
    for input >= 1 {
            res = append(res, 'I')
            input -= 1
    }

    return string(res), nil
}
