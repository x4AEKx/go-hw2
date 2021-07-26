package unpack

import "testing"

type test struct {
	str    string
	expect string
}

var tests = []test{
	{"a4bc2d5e", "aaaabccddddde"},
	{"abcd", "abcd"},
	{"45", ""},
	{"", ""},
}

func TestUnpStr(t *testing.T) {
	for _, test := range tests {
		v := UnpStr(test.str)
		if v != test.expect {
			t.Error(
				"For", test.str,
				"expected", test.expect,
				"got", v,
			)
		}
	}
}
