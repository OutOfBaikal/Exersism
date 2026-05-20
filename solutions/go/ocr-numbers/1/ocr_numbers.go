package ocrnumbers

import (
    "fmt"
    "strings"
)

func recognizeDigit(grid string) string {
	switch grid {
	case "   " +
		 "  |" +
		 "  |":
		return "1"
	case " _ " +
		 " _|" +
		 "|_ ":
		return "2"
	case " _ " +
		 " _|" +
		 " _|":
		return "3"
	case "   " +
		 "|_|" +
		 "  |":
		return "4"
	case " _ " +
		 "|_ " +
		 " _|":
		return "5"
	case " _ " +
		 "|_ " +
		 "|_|":
		return "6"
	case " _ " +
		 "  |" +
		 "  |":
		return "7"
	case " _ " +
		 "|_|" +
		 "|_|":
		return "8"
	case " _ " +
		 "|_|" +
		 " _|":
		return "9"
	case " _ " +
		 "| |" +
		 "|_|":
		return "0"
	default:
		return "?" // Если знак поврежден или не распознан
	}
}

func Recognize(s string) ([]string, error) {
    lines := strings.Split(s[1:], "\n")
	if len(lines)%4 != 0 {
		return nil, fmt.Errorf("number of input lines is not a multiple of four")
	}

	var result []string

    for i := 0; i < len(lines); i += 4 {
        lineLen := len(lines[i])
        if lineLen % 3 != 0 {
            return nil, fmt.Errorf("number of input columns is not a multiple of three")
        }
        var currentNumber strings.Builder
        for j := 0; j < lineLen; j += 3 {
            digitKey := lines[i][j:j+3] + lines[i + 1][j:j+3] + lines[i + 2][j:j+3]
            currentNumber.WriteString(recognizeDigit(digitKey))
        }
    	result = append(result, currentNumber.String())
    }

    return result, nil
}
