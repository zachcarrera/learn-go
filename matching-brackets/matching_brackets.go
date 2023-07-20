package brackets

func Bracket(input string) bool {
	stack := make([]rune, 0)
	for _, char := range input {
		if char == '{' || char == '[' || char == '(' {
			stack = append(stack, char)
		}

		if len(stack) > 0 && ((char == '}' && stack[len(stack)-1] == '{') || (char == ']' && stack[len(stack)-1] == '[') || (char == ')' && stack[len(stack)-1] == '(')) {
			stack = stack[:len(stack)-1]
		} else if char == '}' || char == ']' || char == ')' {
			return false
		}
	}
	return len(stack) == 0
}
