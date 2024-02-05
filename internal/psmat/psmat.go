package psmat

import (
	"sync"

	"go.nhat.io/ligaturizer/internal/python3"
)

const moduleName = "psMat"

var getModule = sync.OnceValue(func() *python3.Object {
	module, err := python3.ImportModule(moduleName)
	if err != nil {
		panic(err)
	}

	return module
})

// Scale returns a matrix which will scale by x in the horizontal direction and y in the vertical.
func Scale(x, y float64) []float64 {
	matrix := getModule().CallMethodArgs("scale", x, y)
	defer matrix.DecRef()

	result := make([]float64, matrix.Length())

	for i := range result {
		item := matrix.GetItem(i)
		result[i] = python3.AsFloat64(item)

		item.DecRef()
	}

	return result
}
