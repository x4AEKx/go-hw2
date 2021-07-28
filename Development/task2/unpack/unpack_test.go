package unpack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type expect struct {
	str    string
	errMsg string
}

type test struct {
	str string
	expect
}

var tests = []test{
	{"a4bc2d5e", expect{str: "aaaabccddddde"}},
	{"da4bc2d5e", expect{str: "daaaabccddddde"}},
	{"abcd", expect{str: "abcd"}},
	{"45", expect{str: "", errMsg: "некорректная строка"}},
	{"", expect{str: ""}},
	{" ", expect{str: " "}},
}

func TestUnpStr(t *testing.T) {
	for _, test := range tests {
		result, err := UnpStr(test.str)
		if result != test.expect.str {
			t.Error(
				"For", test.str,
				"expected", test.expect.str,
				"got", result,
			)
		}
		if err != nil {
			assert.EqualErrorf(t, err, test.expect.errMsg, "Error should be: %v, got: %v", test.expect.errMsg, err)
		}
	}
}
