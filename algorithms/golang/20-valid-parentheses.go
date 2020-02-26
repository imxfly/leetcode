package main

// Stack Data structure implementation of the stack
type Stack struct {
	top int
	arr []rune
}

func (s *Stack) push(item rune) {
	s.arr[s.top] = item
	s.top++
}

func (s *Stack) pop() (rune, bool) {
	if s.top == 0 {
		return 0, false
	}
	s.top--
	ch := s.arr[s.top]
	return ch, true
}

// IsValid judge the validation of string
func IsValid(s string) bool {
	st := &Stack{
		top: 0,
		arr: make([]rune, len(s)),
	}

	for _, ch := range s {
		if ch == 40 || ch == 91 || ch == 123 {
			st.push(ch)
		} else {
			item, ok := st.pop()
			if !ok {
				return false
			}
			if item == 40 && ch != 41 || item == 91 && ch != 93 || item == 123 && ch != 125 {
				return false
			}
		}
	}

	return st.top == 0
}
