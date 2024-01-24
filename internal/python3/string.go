package python3

import python3 "github.com/nhatthm/cpy3"

// GoString converts a Python object to a string.
func GoString(o *PyObject) string {
	str := o.Str()
	defer str.DecRef()

	return python3.PyUnicode_AsUTF8(str)
}

// PyString creates a new Python string object.
func PyString(s string) *PyObject {
	return python3.PyUnicode_FromString(s)
}

// AsString converts a Python object to a string.
func AsString(o *Object) string {
	return GoString(o.PyObject())
}
