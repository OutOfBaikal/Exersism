package transpose

// import (
//     "fmt"
// )

func Transpose(input []string) []string {
	if len(input) == 0 {
		return []string{}
	}

	maxLen := 0
	for _, s := range input {
		if len(s) > maxLen {
			maxLen = len(s)
		}
	}

	res := make([]string, maxLen)

	for i := 0; i < maxLen; i++ {
		var row []byte
		// Идем по строкам оригинальной матрицы
		for j := 0; j < len(input); j++ {
			if i < len(input[j]) {
				// Если в текущей строке есть символ на этой позиции — добавляем его
				// НО сначала нужно проверить, не короче ли были ПРЕДЫДУЩИЕ строки
				// Если короче, их нужно дополнить пробелами (Padding to the left)
				for len(row) < j {
					row = append(row, ' ')
				}
				row = append(row, input[j][i])
			}
		}
		res[i] = string(row)
	}

	return res
}

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
