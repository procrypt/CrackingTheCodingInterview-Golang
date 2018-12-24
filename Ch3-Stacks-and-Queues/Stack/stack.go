package Stack

type Stack struct {
	Key interface{}
	Container []interface{}
}

func(s *Stack) Size() int {
	return len(s.Container)
}

func(s *Stack) Pop() interface{} {
	pop := s.Container[len(s.Container)-1]
	s.Container = append(s.Container[:len(s.Container)-1])
	return pop
}

func(s *Stack) Push(value interface{}) {
	s.Container = append(s.Container, value)
}

func(s *Stack) Peek() interface{} {
	return s.Container[len(s.Container)-1]
}

func(s *Stack) IsEmpty() bool {
	if len(s.Container) != 0 {
		return false
	}
	return true
}