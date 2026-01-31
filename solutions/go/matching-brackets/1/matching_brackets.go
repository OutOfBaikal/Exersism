package brackets

func Bracket(input string) bool {
	stack := []rune{}
    pairs := map[rune]rune {
        ']': '[',
        '}': '{',
        ')': '(',
    }

    for _, ch := range input {
        switch ch {
            case '[', '{', '(':
            	stack = append(stack, ch)
            case ']', '}', ')':
            	if len(stack) == 0 {
                    return false
                }
                top := stack[len(stack) - 1]
            	if pairs[ch] != top {
                    return false
                }
            	stack = stack[:len(stack) - 1]
        }
    }
    return len(stack) == 0
}
