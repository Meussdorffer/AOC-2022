package main

type Stack []string

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Pop() (string, bool) {
	if !s.isEmpty() {
		idx := len(*s) - 1
		str := (*s)[idx]
		*s = (*s)[:idx]
		return str, true
	}
	return "", false
}
