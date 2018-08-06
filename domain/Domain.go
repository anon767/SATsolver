package domain

import (
	"fmt"
	"math/rand"
)

type Domain struct {
	terms [] Term
}

func NewDomain(terms []Term) Domain {
	return Domain{terms}
}

func (domain Domain) Permutate() []Term {
	var termsArray []Term
	for _, term := range domain.terms {
		termsArray = append(termsArray, term.GetGoodPermutations()[0])
	}
	return termsArray
}

func (domain Domain) PermutateParallel() []Term {
	var termsArray []Term
	out := make(chan []Term, len(domain.terms))

	for _, term := range domain.terms {
		go func(term Term) {
			var array []Term
			array = append(array, term.GetGoodPermutations()[0])
			out <- array
		}(term)
	}
	for i := 0; i < len(domain.terms); i++ {
		termsArray = append(termsArray, (<-out)[0])
	}
	return termsArray
}

func evaluateTerms(terms []Term) bool {
	result := true
	for _, term := range terms {
		result = result && term.EvaluateTerm()
	}
	return result
}
func (domain Domain) solve(keys map[string]int8) map[string]int8 {
	initiatedTerms := domain.PermutateParallel()
	for _, term := range initiatedTerms {
		for index, _ := range term.literals {
			if _, ok := keys[term.literals[index].Name]; ok {
				term.literals[index].evaluated = keys[term.literals[index].Name]
			} else {
				keys[term.literals[index].Name] = term.literals[index].evaluated
			}

		}
	}
	if !evaluateTerms(initiatedTerms) {
		fmt.Println("first try failed")
		for key, _ := range keys {
			if (rand.Intn(10) > 5) {
				keys[key] = -keys[key]
				return domain.solve(keys)
			}
		}
	}
	return keys
}
func (domain Domain) Solve() map[string]int8 {
	return domain.solve(make(map[string]int8))
}
