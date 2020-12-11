package day06

// CustomDeclarationForm is the survey made to the passengers' groups
// when arriving to the destination
type CustomDeclarationForm struct {
	Results      string
	Answers      map[rune]int
	Participants int
}

// GetTruthyAnswers returns the number
// of answers which have been truth at most
// by 1 person
func (s *CustomDeclarationForm) GetTruthyAnswers() int {
	return len(s.Answers)
}

// GetCommonTruthyAnswers returns the number
// of answers which all participants have answered yes
func (s *CustomDeclarationForm) GetCommonTruthyAnswers() int {
	count := 0
	for _, numOfYesAnswers := range s.Answers {
		if numOfYesAnswers == s.Participants {
			count++
		}
	}
	return count
}

// NewCustomDeclarationForm creates a new survey
func NewCustomDeclarationForm(results string) CustomDeclarationForm {
	participants := 1
	answers := map[rune]int{}
	for _, a := range results {
		if a == ' ' {
			participants++
		} else {
			answers[a]++
		}
	}
	return CustomDeclarationForm{
		Results:      results,
		Answers:      answers,
		Participants: participants,
	}
}

func cleanCustomDeclarationForm(lines []string) []CustomDeclarationForm {
	var result []CustomDeclarationForm
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
		result = append(result, NewCustomDeclarationForm(line))
	}
	return result
}

// FirstPart checks the number of truthy answers
// that the passengers' groups have answered
// yes at least once
func FirstPart(lines []string) int {
	count := 0
	form := cleanCustomDeclarationForm(lines)
	for _, s := range form {
		count += s.GetTruthyAnswers()
	}
	return count
}

// SecondPart checks the number of truthy answers
// that all passengers' groups have answered yes
func SecondPart(lines []string) int {
	count := 0
	forms := cleanCustomDeclarationForm(lines)
	for _, s := range forms {
		count += s.GetCommonTruthyAnswers()
	}
	return count
}
