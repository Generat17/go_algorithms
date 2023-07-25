package list

import "fmt"

// структура данных list
// каждый элемент list, то есть Node (узел) хранит указать на след и пред узел
// таким образом элементы списка могут распологаться где угодно в памяти и быть связанными
// в то время когда стандратный массив занимает непрерывное место в памяти и расширешние будет
// дорогостоящей операцией, так как придется искать новое место и копировать туда весь массив

type Node struct {
	Data string
	next *Node
	prev *Node
}

type List struct {
	Count int
	first *Node
	last  *Node
}

func (l *List) Add(data string) *Node {
	n := Node{
		Data: data,
		prev: l.last,
	}

	l.Count++

	if l.first == nil && l.last == nil {
		l.first = &n
		l.last = &n
		return &n
	}

	l.last.next = &n

	l.last = &n

	return &n
}

func (l *List) AddFront(data string) *Node {
	n := Node{
		Data: data,
		next: l.first,
	}

	l.Count++

	if l.first == nil && l.last == nil {
		l.first = &n
		l.last = &n
		return &n
	}

	l.first.prev = &n

	l.first = &n

	return &n
}

func (l *List) Find(data string) (*Node, error) {
	n := l.first

	if n != nil {
		if n.Data == data {
			return n, nil
		}
		n = n.next
	}
	return nil, fmt.Errorf("unable to locate %q in list", data)
}

func (l *List) FindReverse(data string) (*Node, error) {
	n := l.last
	for n != nil {
		if n.Data == data {
			return n, nil
		}
		n = n.prev
	}
	return nil, fmt.Errorf("unable to locate %q in list", data)
}

func (l *List) Remove(data string) (*Node, error) {
	n, err := l.Find(data)
	if err != nil {
		return nil, err
	}

	n.prev.next = n.next
	n.next.prev = n.prev

	l.Count--

	return n, nil

}

// Operate принимает в качестве аргумента функцию и применяет ее ко всем узлам list
func (l *List) Operate(f func(n *Node) error) error {
	n := l.first

	for n != nil {
		if err := f(n); err != nil {
			return err
		}
		n = n.next
	}
	return nil
}

func (l *List) OperateReverse(f func(n *Node) error) error {
	n := l.last
	for n != nil {
		if err := f(n); err != nil {
			return err
		}
		n = n.prev
	}
	return nil
}
