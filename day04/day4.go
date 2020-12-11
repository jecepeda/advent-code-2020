package day04

import (
	"regexp"
	"strconv"
	"strings"
)

type Passport map[string]string

// NewPassport creates a new passport from the line
func NewPassport(line string) Passport {
	passport := make(Passport)
	fragments := strings.Split(line, " ")
	for _, f := range fragments {
		parts := strings.Split(f, ":")
		passport[parts[0]] = parts[1]
	}
	return passport
}

// Contains checks the validity of the password
func (p Passport) Contains(fields []string) bool {
	for _, f := range fields {
		if _, ok := p[f]; !ok {
			return false
		}
	}
	return true
}

// Validate validates the password
func (p Passport) Validate(validations ...ValidationFunction) bool {
	for _, v := range validations {
		if !v(p) {
			return false
		}
	}
	return true
}

// ValidationFunction is a function used to validate passports
type ValidationFunction func(p Passport) bool

func validYear(field string, min, max int) ValidationFunction {
	return func(p Passport) bool {
		strField, ok := p[field]
		if !ok {
			return false
		}
		intField, err := strconv.Atoi(strField)
		if err != nil {
			return false
		}
		return intField >= min && intField <= max
	}
}
func validHeight(p Passport) bool {
	strHeight, ok := p["hgt"]
	if !ok {
		return false
	}
	if strings.Contains(strHeight, "cm") {
		strHeight = strings.ReplaceAll(strHeight, "cm", "")
		height, err := strconv.Atoi(strHeight)
		if err != nil {
			return false
		}
		return height >= 150 && height <= 193
	} else if strings.Contains(strHeight, "in") {
		strHeight = strings.ReplaceAll(strHeight, "in", "")
		height, err := strconv.Atoi(strHeight)
		if err != nil {
			return false
		}
		return height >= 59 && height <= 76
	}
	return false
}

func validHairColor(p Passport) bool {
	hair, ok := p["hcl"]
	if !ok {
		return false
	}
	match, err := regexp.Match(`^\#([0-9a-f]{6})$`, []byte(hair))
	if err != nil {
		return false
	}
	return match
}

func validColors(colors ...string) ValidationFunction {
	return func(p Passport) bool {
		color, ok := p["ecl"]
		if !ok {
			return false
		}
		for _, c := range colors {
			if c == color {
				return true
			}
		}
		return false
	}
}

func validPassportID(p Passport) bool {
	passportID, ok := p["pid"]
	if !ok {
		return false
	}
	match, err := regexp.Match(`^([0-9]{9})$`, []byte(passportID))
	if err != nil {
		return false
	}
	return match
}

// FirstPart checks the passports contained on the lines
// and returns the ones that are valid
func FirstPart(lines []string) int {
	passports := cleanPassports(lines)
	var valid int

	for _, p := range passports {
		if p.Contains([]string{"ecl", "eyr", "hcl", "byr", "iyr", "pid", "hgt"}) {
			valid++
		}
	}

	return valid
}

func SecondPart(lines []string) int {
	passports := cleanPassports(lines)
	var valid int

	for _, p := range passports {
		if p.Validate(
			validYear("byr", 1920, 2002),
			validYear("iyr", 2010, 2020),
			validYear("eyr", 2020, 2030),
			validHeight,
			validHairColor,
			validColors("amb", "blu", "brn", "gry", "grn", "hzl", "oth"),
			validPassportID,
		) {
			valid++
		}
	}

	return valid
}

func cleanPassports(lines []string) []Passport {
	var result []Passport
	for i := 0; i < len(lines); i++ {
		var line string
		for i < len(lines) && lines[i] != "" {
			if line == "" {
				line = lines[i]
			} else {
				line += " " + lines[i]
			}
			i++
		}
		result = append(result, NewPassport(line))
	}
	return result
}
