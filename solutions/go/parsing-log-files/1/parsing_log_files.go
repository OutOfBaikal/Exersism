package parsinglogfiles

import (
	"regexp"
	"strings"
    "fmt"
)

func IsValidLine(text string) bool {
	validPrefixes := []string{"[TRC]", "[DBG]", "[INF]", "[WRN]", "[ERR]", "[FTL]"}
	for _, prefix := range validPrefixes {
		if strings.HasPrefix(text, prefix) {
			return true
		}
	}
	return false
}

func SplitLogLine(text string) []string {
	// Regex to match the custom separator
	re := regexp.MustCompile(`<[~*=-]*>`)
	// Split the text using the regex
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	// Regex to find quoted text containing "password"
	re := regexp.MustCompile(`"([^"]*(?i)password[^"]*)"`)

	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	// Replace the matched pattern with an empty string
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+([a-zA-Z0-9_]+)`)

 // Создаем срез для хранения обработанных строк
 taggedLines := make([]string, len(lines))

 // Итерируемся по каждой строке
 for i, line := range lines {
  // Находим имя пользователя
  matches := re.FindStringSubmatch(line)
     for _, x := range matches {
         fmt.Printf(x)
     }
  if len(matches) > 1 {
   // Если найдено имя пользователя, добавляем тег
   userName := matches[1]
   taggedLines[i] = fmt.Sprintf("[USR] %s %s", userName, line)
  } else {
   // Если имя пользователя не найдено, оставляем строку без изменений
   taggedLines[i] = line
  }
 }

 return taggedLines
}
