package ligaturizer_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.nhat.io/ligaturizer/internal/ligaturizer"
)

func TestChar_UnmarshalText(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		data           string
		expectedResult ligaturizer.Char
		expectedError  string
	}{
		{
			scenario:      "empty string",
			data:          "",
			expectedError: `char not supported: ""`,
		},
		{
			scenario: "rune",
			data:     "w",
			expectedResult: ligaturizer.Char{
				Name: "w",
				Rune: 'w',
			},
		},
		{
			scenario: "symbol",
			data:     "asciitilde",
			expectedResult: ligaturizer.Char{
				Name: "asciitilde",
				Rune: '~',
			},
		},
		{
			scenario: "symbol alias",
			data:     "tilde",
			expectedResult: ligaturizer.Char{
				Name: "asciitilde",
				Rune: '~',
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			var actual ligaturizer.Char

			err := actual.UnmarshalText([]byte(tc.data))

			assert.Equal(t, tc.expectedResult, actual)

			if tc.expectedError == "" {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestChar_MarshalText(t *testing.T) {
	t.Parallel()

	c := ligaturizer.Char{
		Name: "asciitilde",
		Rune: '~',
	}

	actual, err := c.MarshalText()
	require.NoError(t, err)

	assert.Equal(t, []byte("asciitilde"), actual)
}

func TestChar_String(t *testing.T) {
	t.Parallel()

	c := ligaturizer.Char{
		Name: "asciitilde",
		Rune: '~',
	}

	actual := c.String()
	expected := "~"

	assert.Equal(t, expected, actual)
}
