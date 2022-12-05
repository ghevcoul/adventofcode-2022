package utils

import (
	"bufio"
	"os"
)

// An iterator that returns one line of a file at a time
// Based on/stolen from https://bbengfort.github.io/2016/12/yielding-functions-for-iteration-golang/
func Readlines(path string) (<-chan string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	chnl := make(chan string)
	go func() {
		for scanner.Scan() {
			chnl <- scanner.Text()
		}
		close(chnl)
	}()

	return chnl, nil
}
