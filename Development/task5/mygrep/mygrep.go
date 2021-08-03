package mygrep

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

// Opts options command line
type Opts struct {
	after      int
	before     int
	context    int
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
}

// NewOpts constructor return opts instance
func NewOpts(after int, before int, context int, count bool, ignoreCase bool, invert bool, fixed bool, lineNum bool) *Opts {
	return &Opts{
		after:      after,
		before:     before,
		context:    context,
		count:      count,
		ignoreCase: ignoreCase,
		invert:     invert,
		fixed:      fixed,
		lineNum:    lineNum,
	}
}

//Search returns matches of a regular expression in the file's content
func Search(pattern string, options Opts, files []string) []string {
	matches := make([]string, 0, 15)

	if options.ignoreCase {
		pattern = "(?i)" + pattern
	}

	for _, file := range files {
		fileMatches := searchFile(pattern, options, file)
		if len(files) > 1 {
			for _, m := range fileMatches {
				matches = append(matches, file+":"+m)
			}
		} else {
			matches = append(matches, fileMatches...)
		}
	}
	return matches
}

func searchFile(pattern string, options Opts, path string) []string {
	matches := make([]string, 0)
	reader, err := open(path)
	if err != nil {
		log.Fatalf("Error: %s", err)
		return nil
	}

	inverted := options.invert
	var i int
	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}

		i++
		line = line[:len(line)-1]
		if match, _ := regexp.Match(pattern, line); match != inverted {
			if options.lineNum {
				line = append([]byte(fmt.Sprintf("%d:", i)), line...)
			}
			if options.fixed {
				regex := regexp.MustCompile(pattern)
				out := regex.ReplaceAll(line, []byte(fmt.Sprintf("\033[1;34m%s\033[0m", pattern)))
				line = out
			}

			matches = append(matches, string(line))
		}
	}

	if options.count {
		strLenOfMatch := fmt.Sprint(len(matches))
		sliceLenOfMatch := strings.Split(strLenOfMatch, "")
		matches = sliceLenOfMatch
	}

	return matches
}

func open(path string) (*bufio.Reader, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New("cannot open file")
	}

	return bufio.NewReader(file), nil
}
