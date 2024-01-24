package psmat

import "go.nhat.io/ligaturizer/internal/python3"

const moduleName = "psMat"

var module *python3.Object

func init() { //nolint: gochecknoinits
	module = python3.MustImportModule(moduleName)
}

// Scale returns a matrix which will scale by x in the horizontal direction and y in the vertical.
func Scale(x, y float64) []float64 {
	matrix := module.CallMethodArgs("scale", x, y)
	defer matrix.DecRef()

	result := make([]float64, matrix.Length())

	for i := range result {
		item := matrix.GetItem(i)
		result[i] = python3.AsFloat64(item)

		item.DecRef()
	}

	return result
}
