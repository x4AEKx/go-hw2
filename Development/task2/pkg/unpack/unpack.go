package unpack

import (
	"errors"
	"strconv"
	"strings"
)

// UnpStr : unpacks the string
func UnpStr(str string) (string, error) {
	if len(str) == 0 {
		return "", nil
	}

	_, err := strconv.Atoi(str)
	if err == nil {
		return "", errors.New("некорректная строка")
	}

	runes := []rune(str)

	mainStr := ""
	prev := runes[0]

	for i := 0; i < len(runes); i++ {
		num, err := strconv.Atoi(string(runes[i]))

		if err == nil {
			mainStr += strings.Repeat(string(prev), num-1)
			prev = runes[i+1]
		} else {
			mainStr += string(runes[i])
			prev = runes[i]
		}
	}

	return mainStr, nil
}
