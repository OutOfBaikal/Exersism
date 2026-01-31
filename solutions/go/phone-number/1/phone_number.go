package phonenumber

import (
    "regexp"
    "fmt"
)

func Number(phoneNumber string) (string, error) {
    re := regexp.MustCompile(`^\s*(?:\+?1[-.\s]?)?(1?(\d{3})|\(?\d{3}\)?)[-.\s]*(\d{3})[-.\s]*(\d{4})\s*$`)
    if !re.MatchString(phoneNumber) {
        return "", fmt.Errorf("Invalid input")
    }
	pattern := regexp.MustCompile(`\d`)
    matches := pattern.FindAllString(phoneNumber, -1)
    result := ""
    for _, match := range matches {
        result += match
    }
    if len(result) == 11 {
        result = result[1:]
    } 
    if result[0] == '1' || result[0] == '0' || result[3] == '1' || result[3] == '0' {
        return "", fmt.Errorf("Invalid input")
    }
    return result, nil
}

func AreaCode(phoneNumber string) (string, error) {
	data, err := Number(phoneNumber)
    if err != nil {
        return "", err
    }
    return data[0:3], nil
}

func Format(phoneNumber string) (string, error) {
	data, err := Number(phoneNumber)
    if err != nil {
        return "", err
    }
    return "(" + data[0:3] + ") " + data[3:6] + "-" + data[6:10], nil
}
