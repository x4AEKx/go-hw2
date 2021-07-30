package find

import (
	"reflect"
	"testing"
)

type expect struct {
	result map[string][]string
}

type test struct {
	str []string
	expect
}

var tests = []test{
	{
		[]string{"пятак", "пятак", "пятка", "Яптка", "апятк", "тяпка", "листок", "ислток", "слиток", "столик"},
		expect{
			result: map[string][]string{
				"листок": {"ислток", "листок", "слиток", "столик"},
				"пятак":  {"апятк", "пятак", "пятка", "тяпка", "яптка"},
			},
		},
	},
}

func TestSetsOfAnagram(t *testing.T) {
	for _, test := range tests {
		result := SetsOfAnagram(test.str)

		if !reflect.DeepEqual(result, test.expect.result) {
			t.Error(
				"\n expected:", test.expect.result, "\n",
				"got:", result,
			)
		}
	}

}
