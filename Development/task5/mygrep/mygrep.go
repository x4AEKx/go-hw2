package mygrep

import (
	"bufio"
	"fmt"
	"io"
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
	reader := open(path)

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
			matches = append(matches, string(line))
		}
	}

	if options.count {
		i := fmt.Sprint(len(matches))
		sliceLetters := strings.Split(i, "")
		matches = sliceLetters
	}

	return matches
}

func open(path string) *bufio.Reader {
	file, err := os.Open(path)
	if err != nil {
		panic("cannot open file")
	}

	return bufio.NewReader(file)
}
