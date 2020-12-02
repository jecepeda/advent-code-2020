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
	rawNumbers, err := ReadFile(path)
	if err != nil {
		return nil, err
	}
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

// ReadFile reads a file and returns its content split
func ReadFile(path string) ([]string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	return lines, nil
}
