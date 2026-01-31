package luhn

import (
    "unicode"
    "strconv"
)

func Valid(id string) bool {
	cleaned := ""
	for _, c := range id {
		if unicode.IsDigit(c) {
			cleaned += string(c)
		} else if !unicode.IsSpace(c) {
			// Если есть символы, кроме цифр и пробелов — невалидно
			return false
		}
	}

	if len(cleaned) <= 1 {
		return false
	}

	sum := 0
	shouldDouble := false

	for i := len(cleaned) - 1; i >= 0; i-- {
		digit, _ := strconv.Atoi(string(cleaned[i]))
		value := digit

		if shouldDouble {
			doubled := digit * 2
			if doubled > 9 {
				doubled -= 9
			}
			value = doubled
		}

		sum += value
		shouldDouble = !shouldDouble
	}

	return sum%10 == 0
}
