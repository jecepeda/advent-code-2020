package day7

import (
	"regexp"
	"strconv"
)

// BagGraph is the graph that contains
// all bags
type BagGraph map[string]*Bag

// NewGraph creates a new bag graph
func NewGraph() BagGraph {
	return make(BagGraph)
}

// AddBag adds a new bag to the system
func (g BagGraph) AddBag(bagID string) *Bag {
	b := NewBag(bagID)
	g[bagID] = b
	return b
}

// Link links two bags given a quantity
func (g BagGraph) Link(from, to string, qty int) {
	bFrom, ok := g[from]
	if !ok {
		bFrom = g.AddBag(from)
	}
	bTo, ok := g[to]
	if !ok {
		bTo = g.AddBag(to)
	}
	bFrom.Contains = append(bFrom.Contains, BagRelationShip{
		Qty: qty,
		Bag: bTo,
	})
	bTo.ContainedIn = append(bTo.ContainedIn, bFrom)
}

// Bag is the bag which is going to be used
// by the rules. It has an id e.g. "shiny gold"
// and tells which bags can contain and which ones
// contains the current bag
type Bag struct {
	ID          string
	Contains    []BagRelationShip
	ContainedIn []*Bag
}

// NewBag creates a new bag
func NewBag(id string) *Bag {
	return &Bag{
		ID:          id,
		Contains:    make([]BagRelationShip, 0),
		ContainedIn: make([]*Bag, 0),
	}
}

// BagsWhichContainMe returns the bags that contains me
func (b *Bag) BagsWhichContainMe() map[string]*Bag {
	contained := make(map[string]*Bag)
	for _, c := range b.ContainedIn {
		contained[c.ID] = c
		bContained := c.BagsWhichContainMe()
		for k, v := range bContained {
			contained[k] = v
		}
	}
	return contained
}

// BagsThatINeedToHold returns the bags that I need to hold in order to carry my own bag
func (b *Bag) BagsThatINeedToHold() int {
	var holded int
	for _, c := range b.Contains {
		hold := c.Bag.BagsThatINeedToHold()
		holded += c.Qty
		holded += c.Qty * hold
	}
	return holded
}

// BagRelationShip is the relationship between
// bag A and bag B, in the way of
// "bag A contains Qty bag B"
type BagRelationShip struct {
	Qty int
	Bag *Bag
}

// BagContained is a simple way of representing
// the amount of bags you can handle
type BagContained struct {
	BagID string
	Qty   int
}

// BagMatcher is a set of regular expressions meant
// to ease the work of finding bags, colors, and amounts
type BagMatcher struct {
	BagColorMatcher      *regexp.Regexp
	NoBagMatcher         *regexp.Regexp
	ContainedBagsMatcher *regexp.Regexp
}

// GetBagColor gets the color of the main bag
func (b *BagMatcher) GetBagColor(str string) string {
	matches := b.BagColorMatcher.FindStringSubmatch(str)
	return matches[1]
}

// HasBags tells if the bag contains bags or not
func (b *BagMatcher) HasBags(str string) bool {
	return !b.NoBagMatcher.MatchString(str)
}

// GetBagsContained gets the number of bags contained in the main bag
func (b *BagMatcher) GetBagsContained(str string) ([]BagContained, error) {
	matches := b.ContainedBagsMatcher.FindAllStringSubmatch(str, -1)
	bags := make([]BagContained, 0, len(matches))
	for _, match := range matches {
		qty, err := strconv.Atoi(match[1])
		if err != nil {
			return nil, err
		}
		bags = append(bags, BagContained{
			BagID: match[2],
			Qty:   qty,
		})

	}
	return bags, nil
}

// NewBagMatcher creates a new bag matcher
func NewBagMatcher() (*BagMatcher, error) {
	bagID := `(\w+\s\w+)`
	numberOfBags := `(\d+)`
	bagColorMatcher, err := regexp.Compile("^" + bagID + " bags contain")
	if err != nil {
		return nil, err
	}
	noBagMatcher, err := regexp.Compile(`no other bags`)
	if err != nil {
		return nil, err
	}
	containedBagMatcher, err := regexp.Compile(numberOfBags + " " + bagID + ` bags?,?`)
	if err != nil {
		return nil, err
	}
	return &BagMatcher{
		BagColorMatcher:      bagColorMatcher,
		NoBagMatcher:         noBagMatcher,
		ContainedBagsMatcher: containedBagMatcher,
	}, nil
}

func buildGraphOfBags(lines []string) (BagGraph, error) {
	matcher, err := NewBagMatcher()
	if err != nil {
		return nil, err
	}
	g := NewGraph()
	for _, line := range lines {
		from := matcher.GetBagColor(line)
		if matcher.HasBags(line) {
			bagsContained, err := matcher.GetBagsContained(line)
			if err != nil {
				return nil, err
			}
			for _, contained := range bagsContained {
				g.Link(from, contained.BagID, contained.Qty)
			}
		}
	}
	return g, nil
}

// FirstPart asks how many bags contains the shiny gold bag
func FirstPart(lines []string) (int, error) {
	g, err := buildGraphOfBags(lines)
	if err != nil {
		return 0, err
	}
	shinyGold, ok := g["shiny gold"]
	if !ok {
		return 0, nil
	}
	bContained := shinyGold.BagsWhichContainMe()
	return len(bContained), nil
}

// SecondPart asks how many bags holds the shiny gold bag
func SecondPart(lines []string) (int, error) {
	g, err := buildGraphOfBags(lines)
	if err != nil {
		return 0, err
	}
	shinyGold, ok := g["shiny gold"]
	if !ok {
		return 0, nil
	}
	return shinyGold.BagsThatINeedToHold(), nil
}
