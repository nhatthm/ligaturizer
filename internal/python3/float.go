package python3

import "go.nhat.io/cpy3"

// PyFloat64 creates a new Python float object.
func PyFloat64(v float64) *PyObject {
	return cpy3.PyFloat_FromDouble(v)
}

// AsFloat64 converts a Python object to a float64.
func AsFloat64(o *Object) float64 {
	return cpy3.PyFloat_AsDouble(o.PyObject())
}
