package day02

import (
	"regexp"
	"strconv"
)

type PasswordValidator struct {
	Char     rune
	Min      int
	Max      int
	Password string
}

// ValidPart1 checks that the password is valid
// by checking the ocurrences of the given character
func (pv PasswordValidator) ValidPart1() bool {
	occurences := make(map[rune]int)
	for _, c := range pv.Password {
		occurences[c]++
	}
	return occurences[pv.Char] >= pv.Min && occurences[pv.Char] <= pv.Max
}

// ValidPart2 checks that the password is valid by checking that
// either the lowest position or the highest position contains the given character
func (pv PasswordValidator) ValidPart2() bool {
	runes := []rune(pv.Password)
	return (runes[pv.Min-1] == pv.Char) != (runes[pv.Max-1] == pv.Char)
}

// CheckPasswordsPart1 checks the passwords introduced by
// the North Pole Toboggan Rental Shop
// and returns how many passwords are correct
func CheckPasswordsPart1(lines []string) int {
	var correct int
	for _, l := range lines {
		pv, err := ConvertToPassword(l)
		if err != nil {
			panic(err)
		}
		if pv.ValidPart1() {
			correct++
		}
	}
	return correct
}

// CheckPasswordsPart2 checks the passwords introduced by
// the North Pole Toboggan Rental Shop
// and returns how many passwords are correct
func CheckPasswordsPart2(lines []string) int {
	var correct int
	for _, l := range lines {
		pv, err := ConvertToPassword(l)
		if err != nil {
			panic(err)
		}
		if pv.ValidPart2() {
			correct++
		}
	}
	return correct
}

// ConvertToPassword converts a line into a password validator
func ConvertToPassword(line string) (PasswordValidator, error) {
	var (
		pv  PasswordValidator
		err error
	)
	r, err := regexp.Compile(`(\d+)-(\d+)\s(\w): (\w+)`)
	if err != nil {
		return PasswordValidator{}, err
	}
	res := r.FindStringSubmatch(line)
	pv.Min, err = strconv.Atoi(res[1])
	if err != nil {
		return PasswordValidator{}, err
	}
	pv.Max, err = strconv.Atoi(res[2])
	if err != nil {
		return PasswordValidator{}, err
	}
	pv.Char = rune(res[3][0])
	pv.Password = res[4]
	return pv, nil
}
