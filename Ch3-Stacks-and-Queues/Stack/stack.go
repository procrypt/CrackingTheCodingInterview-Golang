package Stack

type Stack struct {
	Key interface{}
	Container []interface{}
}

func(s *Stack) Pop() {
	s.Container = append(s.Container[:len(s.Container)-1])
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