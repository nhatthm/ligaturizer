package python3

import (
	"go.nhat.io/cpy3"
)

// Error is a Python error.
type Error string

// Error returns a string representation of the Error.
func (e Error) Error() string {
	return string(e)
}

// ImportError is returned when a Python module cannot be imported.
type ImportError struct {
	Module  string
	Path    string
	Message string
}

// Error returns a string representation of the ImportError.
func (e ImportError) Error() string {
	return e.Message
}

// ModuleNotFoundError is returned when a Python module cannot be found.
type ModuleNotFoundError struct {
	Module  string
	Path    string
	Message string
}

// Error returns a string representation of the ModuleNotFoundError.
func (e ModuleNotFoundError) Error() string {
	return e.Message
}

// MustSuccess panics if the last Python operation failed.
func MustSuccess() {
	if err := LastError(); err != nil {
		panic(err)
	}
}

// LastError returns the last error that occurred in the Python interpreter.
func LastError() error {
	err, traceback := fetchError()
	if err == nil {
		return nil
	}

	defer clearError()
	defer err.DecRef()
	defer traceback.DecRef()

	switch {
	case ErrorIs(err, cpy3.PyExc_ModuleNotFoundError):
		return newErrModuleNotFound(err)

	case ErrorIs(err, cpy3.PyExc_ImportError):
		return newErrImport(err)
	}

	return Error(err.String())
}

// ErrorIs returns true if err is an instance of ex.
func ErrorIs(err *Object, target *PyObject) bool {
	if target == nil {
		return err.Type().PyObject() == target
	}

	return cpy3.PyErr_GivenExceptionMatches(err.Type().PyObject(), target)
}

// clearError clears the last error that occurred in the Python interpreter.
func clearError() {
	cpy3.PyErr_Clear()
}

// fetchError returns the last error that occurred in the Python interpreter.
func fetchError() (*Object, *Object) {
	ex, val, traceback := cpy3.PyErr_Fetch()
	if ex == nil {
		return nil, nil
	}

	if val == nil {
		val = ex
	} else {
		ex.DecRef()
	}

	return NewObject(val), NewObject(traceback)
}

func newErrModuleNotFound(err *Object) error {
	name := err.GetAttr("name")
	path := err.GetAttr("path")
	msg := err.GetAttr("msg")

	defer name.DecRef()
	defer path.DecRef()
	defer msg.DecRef()

	return ModuleNotFoundError{
		Module:  AsString(name),
		Path:    AsString(path),
		Message: AsString(msg),
	}
}

func newErrImport(err *Object) error {
	name := err.GetAttr("name")
	path := err.GetAttr("path")
	msg := err.GetAttr("msg")

	defer name.DecRef()
	defer path.DecRef()
	defer msg.DecRef()

	return ImportError{
		Module:  AsString(name),
		Path:    AsString(path),
		Message: AsString(msg),
	}
}
