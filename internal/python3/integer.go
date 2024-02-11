package python3

import "go.nhat.io/cpy3"

// PyInt creates a new Python int object.
func PyInt(v int) *PyObject {
	return cpy3.PyLong_FromLong(v)
}

// AsInt converts a Python object to an int.
func AsInt(o *Object) int {
	return cpy3.PyLong_AsLong(o.PyObject())
}
