package day10

type Stack struct {
	items []rune
}

func (stack *Stack) Len() int {
	return len(stack.items)
}

func (stack *Stack) Push(item rune) *Stack {
	stack.items = append(stack.items, item)
	return stack
}

func (stack *Stack) Pop() rune {
	item := stack.items[len(stack.items) - 1]
	stack.items = stack.items[:len(stack.items) - 1]
	return item
}

func isMatch(open rune, close rune) bool {
	return (open == '{' && close == '}') ||
		(open == '(' && close == ')') ||
		(open == '<' && close == '>') ||
		(open == '[' && close == ']')
}

func isOpeningChar(ch rune) bool {
	if ch == '{' || ch == '(' || ch == '[' || ch == '<' {
		return true
	}
	return false
}