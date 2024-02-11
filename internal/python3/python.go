package python3

import (
	"go.nhat.io/cpy3"
)

var finializers = make([]func(), 0)

func init() { // nolint: gochecknoinits
	cpy3.Py_Initialize()

	if !cpy3.Py_IsInitialized() {
		panic("could not initializing the python interpreter")
	}
}

// Finalize finializes the python interpreter.
func Finalize() {
	for _, f := range finializers {
		f()
	}

	cpy3.Py_Finalize()
}

func registerFinalizer(f func()) {
	finializers = append(finializers, f)
}
