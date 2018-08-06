package domain

type Term struct {
	literals []Literal
}

func (term Term) GetSize() int {
	return len(term.literals)
}

func NewTerm(literals []Literal) Term {
	return Term{literals}
}

func (term Term) EvaluateTerm() bool {
	acc := false
	for _, element := range term.literals {
		acc = acc || (element.IsTrue())
	}
	return acc
}

func permutate(literals []Literal) [][]Literal {
	if len(literals) == 1 {
		if (literals[0].IsVariable()) {
			return [][]Literal{
				{NewLiteral(literals[0].modifier, literals[0].Name, 1)},
				{NewLiteral(literals[0].modifier, literals[0].Name, -1)}}
		} else {
			return [][]Literal{{literals[0]}}
		}
	} else if len(literals) > 1 {
		var temp = [][]Literal{}
		for _, literalsArray := range permutate(literals[1:]) {
			if (literals[0].IsVariable()) {
				temp = append(temp, append(literalsArray, NewLiteral(literals[0].modifier, literals[0].Name, 1)))
				temp = append(temp, append(literalsArray, NewLiteral(literals[0].modifier, literals[0].Name, -1)))
			} else {
				temp = append(temp, append(literalsArray, literals[0]))
			}
		}
		return temp
	} else {
		return make([][]Literal, 0)
	}
}

func (term Term) Permutate() [][]Literal {
	return permutate(term.literals)
}

func (term Term) GetGoodPermutations() []Term {
	workingTerms := []Term{}
	for _, literalArray := range permutate(term.literals) {
		term2 := NewTerm(literalArray)
		if (term2.EvaluateTerm()) {

			workingTerms = append(workingTerms, term2)
		}
	}
	return workingTerms
}
