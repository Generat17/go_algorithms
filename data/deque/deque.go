package deque

const minCapacity = 16

type Deque[T any] struct {
	buf    []T
	head   int
	tail   int
	count  int
	minCap int
}

func New[T any](size ...int) *Deque[T] {
	var capacity, minimum int
	if len(size) >= 1 {
		capacity = size[0]
		if len(size) >= 2 {
			minimum = size[1]
		}
	}

	minCap := minCapacity
	for minCap < minimum {
		minCap <<= 1
	}

	var buf []T
	if capacity != 0 {
		bufSize := minCap
		for bufSize < capacity {
			bufSize <<= 1
		}
		buf = make([]T, bufSize)
	}

	return &Deque[T]{
		buf:    buf,
		minCap: minCap,
	}
}

func (q *Deque[T]) Cap() int {
	if q == nil {
		return 0
	}
	return len(q.buf)
}

func (q *Deque[T]) Len() int {
	if q == nil {
		return 0
	}
	return q.count
}

func (q *Deque[T]) growIfFull() {
	if q.count != len(q.buf) {
		return
	}
	if len(q.buf) == 0 {
		if q.minCap == 0 {
			q.minCap = minCapacity
		}
		q.buf = make([]T, q.minCap)
		return
	}
	q.resize()
}

func (q *Deque[T]) shrinkIfExcess() {
	if len(q.buf) > q.minCap && (q.count<<2) == len(q.buf) {
		q.resize()
	}
}

func (q *Deque[T]) resize() {
	newBuf := make([]T, q.count<<1)
	if q.tail > q.head {
		copy(newBuf, q.buf[q.head:q.tail])
	} else {
		n := copy(newBuf, q.buf[q.head:])
		copy(newBuf[n:], q.buf[:q.tail])
	}

	q.head = 0
	q.tail = q.count
	q.buf = newBuf
}

func (q *Deque[T]) next(i int) int {
	return (i + 1) & (len(q.buf) - 1) // bitwise modulus
}

func (q *Deque[T]) prev(i int) int {
	return (i - 1) & (len(q.buf) - 1) // bitwise modulus
}

func (q *Deque[T]) SetMinCapacity(minCapacityExp uint) {
	if 1<<minCapacityExp > minCapacity {
		q.minCap = 1 << minCapacityExp
	} else {
		q.minCap = minCapacity
	}
}

func (q *Deque[T]) PushBack(elem T) {
	q.growIfFull()

	q.buf[q.tail] = elem
	// Calculate new tail position.
	q.tail = q.next(q.tail)
	q.count++
}

func (q *Deque[T]) PushFront(elem T) {
	q.growIfFull()

	// Calculate new head position.
	q.head = q.prev(q.head)
	q.buf[q.head] = elem
	q.count++
}

func (q *Deque[T]) PopFront() T {
	if q.count <= 0 {
		panic("deque: PopFront() called on empty deque")
	}
	ret := q.buf[q.head]
	var zero T
	q.buf[q.head] = zero
	// Calculate new head position.
	q.head = q.next(q.head)
	q.count--

	q.shrinkIfExcess()
	return ret
}

func (q *Deque[T]) PopBack() T {
	if q.count <= 0 {
		panic("deque: PopBack() called on empty deque")
	}

	// Calculate new tail position
	q.tail = q.prev(q.tail)

	// Remove value at tail.
	ret := q.buf[q.tail]
	var zero T
	q.buf[q.tail] = zero
	q.count--

	q.shrinkIfExcess()
	return ret
}

func (q *Deque[T]) Front() T {
	if q.count <= 0 {
		panic("deque: Front() called when empty")
	}
	return q.buf[q.head]
}

func (q *Deque[T]) Back() T {
	if q.count <= 0 {
		panic("deque: Back() called when empty")
	}
	return q.buf[q.prev(q.tail)]
}
