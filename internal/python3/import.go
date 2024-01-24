package python3

import (
	python3 "github.com/nhatthm/cpy3"
	"go.nhat.io/once"
)

var modules once.ValuesMap[string, *Object, error]

// ImportModule is a wrapper around the C function PyImport_ImportModule.
func ImportModule(name string) (*Object, error) {
	module, err := modules.Do(name, func() (*Object, error) {
		module := python3.PyImport_ImportModule(name)

		if err := LastError(); err != nil {
			return nil, err
		}

		registerFinalizer(func() { module.DecRef() })

		return NewObject(module), nil
	})
	if err != nil {
		return nil, err
	}

	return module, nil
}

// MustImportModule imports a Python module and panics if it fails.
func MustImportModule(name string) *Object {
	module, err := ImportModule(name)
	if err != nil {
		panic(err)
	}

	return module
}
