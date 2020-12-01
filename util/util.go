// Package util contains the utilities
// that are going to be used during the advent of code
package util

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// FileToIntList reads a file and parses its content into
// an integer list, throwing an error if anything goes wrong
func FileToIntList(path string) ([]int, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	rawNumbers := strings.Split(string(data), "\n")
	numbers := make([]int, 0, len(rawNumbers))
	for _, raw := range rawNumbers {
		if raw == "" {
			continue
		}
		v, err := strconv.Atoi(raw)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, v)
	}
	return numbers, nil
}
