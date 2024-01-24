package python3

import python3 "github.com/nhatthm/cpy3"

// Tuple is a Python tuple.
type Tuple struct {
	obj *PyObject
}

// DecRef decreases the reference count of the object.
func (t *Tuple) DecRef() {
	t.obj.DecRef()
}

// PyObject returns the underlying PyObject.
func (t *Tuple) PyObject() *PyObject {
	return t.obj
}

// Length returns the length of the tuple.
func (t *Tuple) Length() int {
	return t.obj.Length()
}

// Set sets the item at index to value.
func (t *Tuple) Set(index int, value any) {
	defer MustSuccess()

	python3.PyTuple_SetItem(t.obj, index, ToPyObject(value))
}

// Get returns the item at index.
func (t *Tuple) Get(index int) *Object {
	defer MustSuccess()

	return NewObject(python3.PyTuple_GetItem(t.obj, index))
}

// NewTuple creates a new tuple.
func NewTuple(length int) *Tuple {
	return &Tuple{
		obj: python3.PyTuple_New(length),
	}
}

// NewTupleFromValues converts a slice of any to a tuple.
func NewTupleFromValues[T any](values ...T) *Tuple {
	tuple := NewTuple(len(values))

	for i, v := range values {
		tuple.Set(i, v)
	}

	return tuple
}

// NewTupleFromAnything converts a slice of any to a tuple.
func NewTupleFromAnything(values ...any) *Tuple {
	return NewTupleFromValues(values...)
}
