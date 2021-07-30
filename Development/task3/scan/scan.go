package scan

import (
	"bufio"
	"errors"
	"os"
)

func ReadFile(fn string) ([]string, error) {
	file, err := os.Open(fn)
	if err != nil {
		// log.Fatalf("Error when opening file: %s", err)
		return nil, errors.New("error when opening file")
	}
	defer file.Close()

	var strs []string

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		strs = append(strs, fileScanner.Text())
	}
	if err := fileScanner.Err(); err != nil {
		// log.Fatalf("Error while reading file: %s", err)
		return nil, errors.New("error while reading file")
	}

	return strs, nil
}
