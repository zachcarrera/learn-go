package brackets

func Bracket(input string) bool {
	bracketPair := map[rune]rune{
		']': '[',
		'}': '{',
		')': '(',
	}
	stack := make([]rune, 0)
	for _, char := range input {
		switch char {
		case '{', '[', '(':
			stack = append(stack, char)
		case '}', ']', ')':
			if len(stack) == 0 || stack[len(stack)-1] != bracketPair[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
