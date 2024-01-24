package python3

import python3 "github.com/nhatthm/cpy3"

// PyInt creates a new Python int object.
func PyInt(v int) *PyObject {
	return python3.PyLong_FromLong(v)
}

// AsInt converts a Python object to an int.
func AsInt(o *Object) int {
	return python3.PyLong_AsLong(o.PyObject())
}
