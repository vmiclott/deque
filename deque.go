package deque

type Deque[T any] []T

func (q *Deque[T]) Prepend(x T) {
	*q = append([]T{x}, *q...)
}

func (q *Deque[T]) Append(x T) {
	*q = append(*q, x)
}

func (q *Deque[T]) PopLeft() T {
	old := *q
	n := len(*q)
	x := old[0]
	*q = old[1:n]
	return x
}

func (q *Deque[T]) PopRight() T {
	old := *q
	n := len(*q)
	x := old[n-1]
	*q = old[0 : n-1]
	return x
}

func (q Deque[T]) Len() int {
	return len(q)
}

func (q Deque[T]) PeekRight() T {
	return q[len(q)-1]
}

func (q Deque[T]) PeekLeft() T {
	return q[0]
}
