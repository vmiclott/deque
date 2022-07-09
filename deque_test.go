package deque

import (
	"reflect"
	"testing"
)

func TestLen(t *testing.T) {
	q := Deque[int]{1, 2, 3}
	want := 3
	got := q.Len()

	if got != want {
		t.Errorf("expected: %v, got: %v", want, got)
	}
}

func TestPeekLeft(t *testing.T) {
	type test struct {
		q    Deque[int]
		want int
	}

	tests := []test{
		{q: Deque[int]{1}, want: 1},
		{q: Deque[int]{-1, 1}, want: -1},
		{q: Deque[int]{4, 3, 2}, want: 4},
	}

	for _, tc := range tests {
		got := tc.q.PeekLeft()
		if got != tc.want {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestPeekRight(t *testing.T) {
	type test struct {
		q    Deque[int]
		want int
	}

	tests := []test{
		{q: Deque[int]{1}, want: 1},
		{q: Deque[int]{-1, 1}, want: 1},
		{q: Deque[int]{4, 3, 2}, want: 2},
	}

	for _, tc := range tests {
		got := tc.q.PeekRight()
		if got != tc.want {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestPrepend(t *testing.T) {
	q := Deque[int]{}
	prepend := []int{1, -1, 2, -2}
	want := Deque[int]{-2, 2, -1, 1}
	for _, x := range prepend {
		q.Prepend(x)
	}

	if !reflect.DeepEqual(q, want) {
		t.Errorf("expected: %v, got: %v", want, q)
	}
}

func TestAppend(t *testing.T) {
	q := Deque[int]{}
	append := []int{1, -1, 2, -2}
	want := Deque[int]{1, -1, 2, -2}
	for _, x := range append {
		q.Append(x)
	}

	if !reflect.DeepEqual(want, q) {
		t.Errorf("expected: %v, got: %v", want, q)
	}
}

func TestPopLeft(t *testing.T) {
	q := Deque[int]{2, -1, 1}
	want := 2
	got := q.PopLeft()

	if got != want {
		t.Errorf("wrong popped element value; expected: %v, got: %v", want, got)
	}

	if !reflect.DeepEqual(Deque[int]{-1, 1}, q) {
		t.Errorf("wrong queue after popping; expected: %v, got: %v", Deque[int]{-1, 1}, q)
	}
}

func TestPopRight(t *testing.T) {
	q := Deque[int]{1, -1, 2}
	want := 2
	got := q.PopRight()

	if got != want {
		t.Errorf("wrong popped element value; expected: %v, got: %v", want, got)
	}

	if !reflect.DeepEqual(Deque[int]{1, -1}, q) {
		t.Errorf("wrong queue after popping; expected: %v, got: %v", Deque[int]{1, -1}, q)
	}
}
