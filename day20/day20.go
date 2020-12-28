package day20

import (
	"strconv"
	"strings"
)

type PartialImage struct {
	ID             int
	Image          [][]rune
	MatchingImages []int
	ComputedSides  []string
	Visited        bool
}

func NewPartialImage(id int, image [][]rune) PartialImage {
	pi := PartialImage{
		ID:    id,
		Image: image,
	}
	pi.BuildSides()
	return pi
}

func (p PartialImage) HasCoincidences(p2 PartialImage) (bool, string) {
	for _, p1s := range p.ComputedSides {
		for _, p2s := range p2.ComputedSides {
			if p1s == p2s {
				return true, p1s
			}
		}
	}
	return false, ""
}

func (p *PartialImage) BuildSides() {
	sides := make([]string, 0, 4)
	sides = append(sides, string(p.Image[0]))
	sides = append(sides, reverse(string(p.Image[0])))
	var right []rune
	var left []rune
	for i := range p.Image {
		left = append(left, p.Image[i][0])
		right = append(right, p.Image[i][len(p.Image)-1])
	}
	sides = append(sides, string(left))
	sides = append(sides, reverse(string(left)))
	sides = append(sides, string(right))
	sides = append(sides, reverse(string(right)))
	sides = append(sides, string(p.Image[len(p.Image)-1]))
	sides = append(sides, reverse(string(p.Image[len(p.Image)-1])))
	p.ComputedSides = sides
}

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

func findMatches(images []PartialImage) ([]PartialImage, map[[2]int]string) {
	matches := map[[2]int]string{}
	for i := range images {
		first := images[i]
		for _, second := range images {
			if first.ID == second.ID {
				continue
			}
			if v, match := first.HasCoincidences(second); v {
				images[i].MatchingImages = append(images[i].MatchingImages, second.ID)
				matches[[2]int{first.ID, second.ID}] = match
			}
		}
	}
	return images, matches
}

func readLines(lines []string) ([]PartialImage, error) {
	var images []PartialImage
	var (
		tileID       int
		partialImage [][]rune
		err          error
	)
	for _, line := range lines {
		if strings.HasPrefix(line, "Tile") {
			replaced := strings.Trim(strings.ReplaceAll(line, "Tile ", ""), ":")
			tileID, err = strconv.Atoi(replaced)
			if err != nil {
				return nil, err
			}
		} else if line == "" {
			images = append(images, NewPartialImage(tileID, partialImage))
			tileID = 0
			partialImage = nil
		} else {
			partialImage = append(partialImage, []rune(line))
		}
	}
	return images, nil
}

func FirstPart(lines []string) (int, error) {
	images, err := readLines(lines)
	if err != nil {
		return 0, err
	}
	images, _ = findMatches(images)
	result := 1
	for _, im := range images {
		if len(im.MatchingImages) == 2 {
			result *= im.ID
		}
	}
	return result, nil
}
