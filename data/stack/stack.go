package stack

import "errors"

type Data struct {
	Name string
}

type Stack struct {
	data []*Data
}

func Make(cap int) *Stack {
	return &Stack{
		data: make([]*Data, 0, cap),
	}
}

func (s *Stack) Count() int {
	return len(s.data)
}

func (s *Stack) Push(data *Data) {
	s.data = append(s.data, data)
}

func (s *Stack) Pop() (*Data, error) {
	if len(s.data) == 0 {
		return nil, errors.New("stack empty")
	}

	idx := len(s.data) - 1

	data := s.data[idx]

	s.data = s.data[:idx]

	return data, nil
}

func (s *Stack) Peek(level int) (*Data, error) {
	if level < 0 || level > (len(s.data)-1) {
		return nil, errors.New("invalid level position")
	}
	idx := (len(s.data) - 1) - level
	return s.data[idx], nil
}

func (s *Stack) Operate(f func(data *Data) error) error {
	for i := len(s.data) - 1; i > -1; i-- {
		if err := f(s.data[i]); err != nil {
			return err
		}
	}
	return nil
}
