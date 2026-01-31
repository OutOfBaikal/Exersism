package reverse

func Reverse(input string) string {
    data := []rune(input)
    for i := 0; i < len(data) / 2; i++ {
        data[i], data[len(data) - i - 1] = data[len(data) - i - 1], data[i]
    }
	return string(data)
}
