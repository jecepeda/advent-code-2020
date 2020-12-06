package day4

import (
	"strings"
	"testing"

	"github.com/jecepeda/advent-code-2020/util"
	"github.com/stretchr/testify/require"
)

func TestFirstPart(t *testing.T) {
	tests := []struct {
		name string
		path string
		want int
	}{
		{
			name: "test case",
			path: "./test.txt",
			want: 2,
		},
		{
			name: "real part 1",
			path: "./input.txt",
			want: 192,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := util.ReadFile(tt.path)
			require.NoError(t, err)
			if got := FirstPart(lines); got != tt.want {
				t.Errorf("FirstPart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSecondPart(t *testing.T) {
	tests := []struct {
		name string
		path string
		want int
	}{
		{
			name: "test case",
			path: "./test.txt",
			want: 2,
		},
		{
			name: "real part 1",
			path: "./input.txt",
			want: 101,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := util.ReadFile(tt.path)
			require.NoError(t, err)
			if got := SecondPart(lines); got != tt.want {
				t.Errorf("SecondPart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInvalidPassports(t *testing.T) {
	tests := []struct {
		name    string
		rawLine string
		want    int
	}{
		{
			name: "test case",
			rawLine: `eyr:1972 cid:100
hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926

iyr:2019
hcl:#602927 eyr:1967 hgt:170cm
ecl:grn pid:012533040 byr:1946

hcl:dab227 iyr:2012
ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277

hgt:59cm ecl:zzz
eyr:2038 hcl:74454a iyr:2023
pid:3556412378 byr:2007`,
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines := strings.Split(tt.rawLine, "\n")
			if got := SecondPart(lines); got != tt.want {
				t.Errorf("InvalidPassports() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidPassports(t *testing.T) {
	tests := []struct {
		name    string
		rawLine string
		want    int
	}{
		{
			name: "test case",
			rawLine: `pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f

eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022

iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719`,
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines := strings.Split(tt.rawLine, "\n")
			if got := SecondPart(lines); got != tt.want {
				t.Errorf("InvalidPassports() = %v, want %v", got, tt.want)
			}
		})
	}
}
