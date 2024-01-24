package psmat_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/ligaturizer/internal/psmat"
)

func TestScale(t *testing.T) {
	t.Parallel()

	actual := psmat.Scale(1.15, 1.65)
	expected := []float64{1.15, 0, 0, 1.65, 0, 0}

	assert.Equal(t, expected, actual)
}
