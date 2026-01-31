package wordy

import (
    "strconv"
    "strings"
    "regexp"
)

func Answer(question string) (int, bool) {
	re := regexp.MustCompile(`^What is (-?\d+)( (plus|minus|multiplied by|divided by) (-?\d+))*\?$`)

    if !re.MatchString(question) {
        return 0, false
    }

    question = strings.TrimPrefix(question, "What is ")
    question = strings.TrimSuffix(question, "?")

    parts := strings.Fields(question)
    result, err := strconv.Atoi(parts[0])
    if err != nil {
        return 0, false
    }
    i := 1

    for i < len(parts) {
        operator := parts[i]
        i++
        if parts[i] == "by" {
            i++
        }
        nextNum, err := strconv.Atoi(parts[i])
        if err != nil {
        	return 0, false
    	}
        i++

        switch operator {
            case "plus":
            	result += nextNum
        	case "minus":
            	result -= nextNum
            case "multiplied":
            	result *= nextNum
            case "divided":
            	if nextNum == 0 {
                    return 0, false
                }
            	result /= nextNum
        	default:
            	return 0, false
        }
    }

    return result, true
}
