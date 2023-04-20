package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Define a test case struct
type testCase struct {
	name           string
	input          string
	expectedResult float64
	hasError       bool
}

func TestFindUnique(t *testing.T) {
	// Define a slice of testCase as test table
	testTable := []testCase{
		{
			name:           `Regular test with "+" and "-"`,
			input:          "1+2-3",
			expectedResult: 0,
			hasError:       false,
		}, {
			name:           `Test with "*" and "/" to reveal the operations priority`,
			input:          "2+2/2+3*4",
			expectedResult: 15,
			hasError:       false,
		}, {
			name:           `Test with "()" to reveal brackets handling`,
			input:          "3*(2+1)",
			expectedResult: 9,
			hasError:       false,
		}, {
			name:           `Test with negative numbers`,
			input:          "-3*(-15)/2+(-17*2)/2",
			expectedResult: 5.5,
			hasError:       false,
		}, {
			name:           `Test with the only number`,
			input:          "4",
			expectedResult: 4,
			hasError:       false,
		}, {
			name:           `Test with axcess signs anywhere`,
			input:          "1+-3*",
			expectedResult: -2,
			hasError:       false,
		}, {
			name:           `Test with floats`,
			input:          "-1*(-1*17.73)-(-11*(14+1.3)-3/(-1-2.79))+3",
			expectedResult: 188.238443,
			hasError:       false,
		}, {
			name:     `Test with a wrong input`,
			input:    "1=3",
			hasError: true,
		}, {
			name:     `Test with a wrong input`,
			input:    "1abcd7",
			hasError: true,
		}, {
			name:     `Test with a wrong input`,
			input:    "16/1..7+(46-53)",
			hasError: true,
		},
	}
	// Begin test
	for _, test := range testTable {
		actual, err := GetResult(test.input)

		res := assert.Equal(t, test.expectedResult, actual, test.name)
		if test.hasError {
			assert.NotNil(t, err, test.name)
		} else {
			assert.Nil(t, err, test.name)
		}
		assert.True(t, res)
	}
}
