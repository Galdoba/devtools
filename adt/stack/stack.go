package stack

type Stack interface {
	Push(string)
	Pop() string
	Top() string
	Size() int
	IsEmpty() bool
}

type stack[T any] struct {
	Push func(T)
	Pop  func() T
	Len  func() int
}

func NewStack[T any]() stack[T] {
	data := make(map[int]T)
	s := stack[T]{
		Push: func(i T) {
			data[len(data)+1] = i
		},
		Pop: func() T {
			if len(data) < 1 {
				return *new(T)
			}
			res := data[len(data)]
			delete(data, len(data))
			return res
		},
		Len: func() int {
			return len(data)
		},
	}

	return s
}
