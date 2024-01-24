package python3

import (
	python3 "github.com/nhatthm/cpy3"
)

var finializers = make([]func(), 0)

func init() { // nolint: gochecknoinits
	python3.Py_Initialize()

	if !python3.Py_IsInitialized() {
		panic("could not initializing the python interpreter")
	}
}

// Finalize finializes the python interpreter.
func Finalize() {
	for _, f := range finializers {
		f()
	}

	python3.Py_Finalize()
}

func registerFinalizer(f func()) {
	finializers = append(finializers, f)
}
