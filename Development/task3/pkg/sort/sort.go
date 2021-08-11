package sort

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type args struct {
	fields    string
	number    bool
	reverse   bool
	duplicate bool
}

// Grep options command line
type Sort struct {
	args
	files  []string
	result []string
}

// New constructor return Grep instance
func New(fields string, number bool, reverse bool, duplicate bool, files []string) *Sort {
	return &Sort{
		args: args{
			fields:    fields,
			number:    number,
			reverse:   reverse,
			duplicate: duplicate,
		},

		files: files,
	}
}

//Execute writes the result in struct and returns an error
func (s *Sort) Execute() error {
	matches := make([]string, 0)

	for _, file := range s.files {
		fileMatches, err := searchFile(s.args, file)
		if err != nil {
			return err
		}

		matches = append(matches, fileMatches...)

		if s.args.duplicate {
			newMatches := make([]string, 0)

			set := map[string]struct{}{}

			for _, match := range matches {
				if _, ok := set[match]; !ok {
					newMatches = append(newMatches, match)
					set[match] = struct{}{}
				}
			}
			matches = newMatches
		}

		if s.args.reverse {
			sort.Sort(sort.Reverse(sort.StringSlice(matches)))
		} else {
			sort.Strings(matches)
		}

	}

	s.result = matches

	return nil
}

// Output writes the result in Stdout and returns an error
func (s *Sort) Output() error {
	_, err := fmt.Fprintln(os.Stdout, strings.Join(s.result, "\n"))
	return err
}

func searchFile(options args, path string) ([]string, error) {
	matches := make([]string, 0)
	reader, err := open(path)
	if err != nil {
		return nil, err
	}

	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}

		line = line[:len(line)-1]

		matches = append(matches, string(line))
	}

	return matches, nil
}

func open(path string) (*bufio.Reader, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New("no such file or directory")
	}

	return bufio.NewReader(file), nil
}
