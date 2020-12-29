package day21

import (
	"regexp"
	"sort"
	"strings"
)

var (
	regexIngredients *regexp.Regexp
	regexAllergens   *regexp.Regexp
)

type Allergen struct {
	Name        string
	Ingredients map[string]bool
}

func NewAllergen(name string, ingredients []string) Allergen {
	a := Allergen{
		Name:        name,
		Ingredients: make(map[string]bool),
	}
	for _, i := range ingredients {
		a.Ingredients[i] = true
	}
	return a
}

func (a *Allergen) SetNewIngredients(ingredients []string) {
	var toRemove []string

	incoming := make(map[string]bool)
	for _, i := range ingredients {
		incoming[i] = true
	}

	for k := range a.Ingredients {
		if !incoming[k] {
			toRemove = append(toRemove, k)
		}
	}

	for _, remove := range toRemove {
		delete(a.Ingredients, remove)
	}
}

func (a *Allergen) GetIngredient() string {
	i := ""
	for k := range a.Ingredients {
		i = k
	}
	return i
}

func FirstPart(lines []string) int {
	allergens, ingredients := readLines(lines)
	notAllergens := 0
	for ingredient, times := range ingredients {
		if !isIngredientInAllergens(ingredient, allergens) {
			notAllergens += times
		}
	}
	return notAllergens
}

func SecondPart(lines []string) string {
	allergensByName, _ := readLines(lines)
	var allergens = make([][2]string, 0)
	for _, a := range allergensByName {
		allergens = append(allergens, [2]string{a.Name, a.GetIngredient()})
	}
	sort.Slice(allergens, func(i, j int) bool {
		return allergens[i][0] < allergens[j][0]
	})
	var result strings.Builder
	for i, data := range allergens {
		if i > 0 {
			result.WriteByte(',')
		}
		result.WriteString(data[1])
	}

	return result.String()
}

func readLines(lines []string) (map[string]Allergen, map[string]int) {
	ingredients := make(map[string]int)
	allergens := make(map[string]Allergen)
	for _, l := range lines {
		matches := regexIngredients.FindAllStringSubmatch(l, -1)
		rawIngredients := strings.Split(strings.TrimSpace(matches[0][1]), " ")
		matches = regexAllergens.FindAllStringSubmatch(l, -1)
		rawAllergens := strings.Split(strings.TrimSpace(matches[0][1]), ", ")
		for _, rawAllergen := range rawAllergens {
			a, ok := allergens[rawAllergen]
			if !ok {
				a = NewAllergen(rawAllergen, rawIngredients)
			} else {
				a.SetNewIngredients(rawIngredients)
			}
			allergens[rawAllergen] = a
		}
		for _, ingredient := range rawIngredients {
			ingredients[ingredient]++
		}
	}
	finalAllergens := map[string]Allergen{}
	for len(allergens) > 0 {
		for name := range allergens {
			allergen := allergens[name]
			if len(allergen.Ingredients) == 1 {
				finalAllergens[name] = allergen
				delete(allergens, name)
				for ingredient := range allergen.Ingredients {
					// delete this ingredient from the allergens
					for k, v := range allergens {
						delete(v.Ingredients, ingredient)
						allergens[k] = v
					}
				}
			}
		}
	}
	return finalAllergens, ingredients
}

func toList(m map[string]int) []string {
	result := make([]string, 0, len(m))
	for k := range m {
		result = append(result, k)
	}
	return result
}

func isIngredientInAllergens(ingredient string, allergens map[string]Allergen) bool {
	for _, a := range allergens {
		if a.Ingredients[ingredient] {
			return true
		}
	}
	return false
}

func init() {
	regexIngredients = regexp.MustCompile(`((\w+\s)+)\(`)
	regexAllergens = regexp.MustCompile(`contains ((\w+\,?\s?)+)`)
}
