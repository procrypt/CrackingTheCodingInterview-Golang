package Queue

type Queue struct {
	Key interface{}
	Container []interface{}
}


func(q *Queue) Add(value interface{}) {
	q.Container = append(q.Container, value)
}

func(q *Queue) Remove() {
	q.Container = append(q.Container[1:])
}

func(q *Queue) Peek() interface{} {
	return q.Container[len(q.Container)-1]
}

func(q *Queue) IsEmpty() bool {
	if len(q.Container) != 0 {
		return false
	}
	return true
}

