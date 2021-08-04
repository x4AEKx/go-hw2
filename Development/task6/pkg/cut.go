package pkg

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Opts options command line
type Opts struct {
	fields    string
	delimiter string
	separated bool
}

// NewOpts constructor return opts instance
func NewOpts(fields string, delimiter string, separated bool) *Opts {
	return &Opts{
		fields:    fields,
		delimiter: delimiter,
		separated: separated,
	}
}

//Cut returns slice strings of cut strings in the stdin
func Cut(options Opts) ([]string, error) {
	matches := make([]string, 0)

	fields := strings.Split(options.fields, ",")

	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadBytes('\n')

		if err == io.EOF {
			break
		}

		line = line[:len(line)-1]

		lineSliceByDelimiter := bytes.Split(line, []byte(options.delimiter))

		newLine := make([]byte, 0)

		match, _ := regexp.Match(options.delimiter, line)
		if match {
			for _, v := range fields {
				field, err := strconv.Atoi(v)
				if err != nil {
					return nil, errors.New("invalid field value")
				}
				if field <= 0 {
					return nil, errors.New("fields are numbered from 1")
				}
				if field <= len(lineSliceByDelimiter) {
					newLine = append(newLine, lineSliceByDelimiter[field-1]...)
				}
			}

			matches = append(matches, string(newLine))
		}
	}

	return matches, nil
}
