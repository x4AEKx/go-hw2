package scan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type expect struct {
	str    []string
	errMsg string
}

type test struct {
	str string
	expect
}

var tests = []test{
	{"", expect{errMsg: "error when opening file"}},
	{"../file.txt", expect{str: []string{"Qwerty", "Asdfg", "Zxcvb"}}},
}

func TestReadFile(t *testing.T) {
	for _, test := range tests {
		result, err := ReadFile(test.str)

		if result != nil && result[0] != test.expect.str[0] {
			t.Error(
				"For opened file", test.str,
				"expected", test.expect.str,
				"got", result,
			)
		}

		if err != nil {
			assert.EqualErrorf(t, err, test.expect.errMsg, "Error should be: %v, got: %v", test.expect.errMsg, err)
		}
	}

}
