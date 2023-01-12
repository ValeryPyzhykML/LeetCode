package main

type Stack struct {
	Last  *StackNode
	Count int
}

type StackNode struct {
	Val  byte
	Prev *StackNode
}

func (s *Stack) Push(char byte) {
	s.Last = &StackNode{Val: char, Prev: s.Last}
	s.Count++
}

func (s *Stack) Pop() byte {
	result := s.Last.Val
	s.Last = s.Last.Prev
	s.Count--
	return result
}

func isValid(s string) bool {
	stack := Stack{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack.Push(s[i])
		} else if stack.Count == 0 || s[i]-stack.Pop() > 2 { // 2 is bigest unicode diff betwen open and clode parentheses
			return false
		}
	}

	return stack.Count == 0
}