package day19

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Kind represents the kind of rule we're using
type Kind int

// values Kind can get
const (
	Terminal Kind = iota
	NonTerminal
)

// Operation represents the rule numbers that correspond to an operation
type Operation []int

// Rule defines the rules to match a set of characters
type Rule struct {
	Kind       Kind
	Number     int
	Operations []Operation
	StrRule    string
	Raw        string
}

const elevenRulePlaceHolder = "THIS_WILL_NEVER_MATCH"

// Build buils the given rule
func (r *Rule) Build(rules map[int]*Rule) {
	// already built
	if r.StrRule != "" {
		return
	}
	switch r.Kind {
	case Terminal:
		r.StrRule = r.Raw
	case NonTerminal:
		operations := []string{}
		for _, operation := range r.Operations {
			strOp := ""
			for _, o := range operation {
				if r.Number == o {
					switch r.Number {
					case 8:
						strOp += "+"
					case 11:
						strOp += elevenRulePlaceHolder
					}
					continue
				}
				v := rules[o]
				if v.StrRule == "" {
					v.Build(rules)
				}
				strOp += v.StrRule
			}
			operations = append(operations, strOp)
		}
		r.StrRule = "(" + strings.Join(operations, "|") + ")"
	}
}

func ValidateSentences(lines []string) (int, error) {
	rules, sentences, err := parseLines(lines)
	if err != nil {
		return 0, err
	}
	strZeroRule := rules[0].StrRule
	strZeroRule = fmt.Sprintf("^%s$", strZeroRule)
	elevenRule, ok := rules[11]
	if ok {
		for i := 0; i < 10; i++ {
			strZeroRule = strings.ReplaceAll(strZeroRule, elevenRulePlaceHolder, elevenRule.StrRule)
		}
	}

	var zeroRegex = regexp.MustCompile(strZeroRule)
	var matches int
	for _, s := range sentences {
		if zeroRegex.MatchString(s) {
			matches++
		}
	}
	return matches, nil
}

func parseLines(lines []string) (map[int]*Rule, []string, error) {
	var (
		messages     = make([]string, 0)
		rules        = make(map[int]*Rule)
		phase    int = 0
	)
	for _, l := range lines {
		if l == "" {
			phase++
			continue
		}
		if phase == 0 {
			fragments := strings.Split(l, ":")
			ruleNr, err := strconv.Atoi(fragments[0])
			if err != nil {
				return nil, nil, err
			}
			rule := &Rule{
				Number: ruleNr,
			}
			if strings.Contains(fragments[1], "\"") {
				rule.Kind = Terminal
				character := strings.TrimSpace(strings.ReplaceAll(fragments[1], "\"", ""))
				rule.Raw = character
			} else {
				rule.Kind = NonTerminal
				rawOperations := strings.Split(fragments[1], "|")
				for _, rawOp := range rawOperations {
					var ops []int
					fragments = strings.Split(strings.TrimSpace(rawOp), " ")
					for _, f := range fragments {
						v, err := strconv.Atoi(f)
						if err != nil {
							return nil, nil, err
						}
						ops = append(ops, v)
					}
					rule.Operations = append(rule.Operations, ops)
				}
			}

			rules[ruleNr] = rule
		} else {
			messages = append(messages, l)
		}
	}

	for k := range rules {
		rules[k].Build(rules)
	}

	return rules, messages, nil
}
