package python3

import python3 "github.com/nhatthm/cpy3"

// PyFloat64 creates a new Python float object.
func PyFloat64(v float64) *PyObject {
	return python3.PyFloat_FromDouble(v)
}

// AsFloat64 converts a Python object to a float64.
func AsFloat64(o *Object) float64 {
	return python3.PyFloat_AsDouble(o.PyObject())
}
