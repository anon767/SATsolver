package tests

import (
	"../domain"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPermutate(t *testing.T) {
	terms := []domain.Term{}
	terms = append(terms, domain.NewTerm([]domain.Literal{domain.NewLiteral(true, "A", 0)}))
	terms = append(terms, domain.NewTerm([]domain.Literal{domain.NewLiteral(true, "B", 0)}))

	domain := domain.NewDomain(terms)
	assert.Equal(t, 2, len(domain.Permutate()))
}

func TestPermutateParallel(t *testing.T) {
	terms := []domain.Term{}
	terms = append(terms, domain.NewTerm([]domain.Literal{domain.NewLiteral(true, "A", 0)}))
	terms = append(terms, domain.NewTerm([]domain.Literal{domain.NewLiteral(true, "B", 0)}))
	domain := domain.NewDomain(terms)
	assert.Equal(t, len(domain.Permutate()), len(domain.PermutateParallel()))
}
func TestSolveSingleTerm(t *testing.T) {
	terms := []domain.Term{}
	terms = append(terms, domain.NewTerm([]domain.Literal{domain.NewLiteral(true, "A", 0)}))
	solvedTerm := map[string]int8{"A": 1}
	domain := domain.NewDomain(terms)
	assert.Equal(t, solvedTerm, domain.Solve())
}
func TestSolveMoreTwoTerm(t *testing.T) {
	terms := []domain.Term{}
	terms = append(terms, domain.NewTerm([]domain.Literal{domain.NewLiteral(true, "A", 0)}))
	terms = append(terms, domain.NewTerm([]domain.Literal{domain.NewLiteral(true, "B", 0)}))
	solvedTerm := map[string]int8{"A": 1, "B": 1}
	domain := domain.NewDomain(terms)
	assert.Equal(t, solvedTerm, domain.Solve())
}
func TestSolveMoreTwoTermMultipleLiterals(t *testing.T) {
	terms := []domain.Term{}
	terms = append(terms, domain.NewTerm([]domain.Literal{domain.NewLiteral(true, "A", 0), domain.NewLiteral(false, "B", 0)}))
	terms = append(terms, domain.NewTerm([]domain.Literal{domain.NewLiteral(true, "B", 0)}))
	solvedTerm := map[string]int8{"A": 1, "B": 1}
	domain := domain.NewDomain(terms)
	assert.Equal(t, solvedTerm, domain.Solve())
}
func TestSolveMoreMoreComplexStuff(t *testing.T) {
	fmt.Println("start")
	terms := []domain.Term{}
	terms = append(terms, domain.NewTerm([]domain.Literal{domain.NewLiteral(false, "A", 0), domain.NewLiteral(false, "B", 0), domain.NewLiteral(false, "C", 0)}))
	terms = append(terms, domain.NewTerm([]domain.Literal{domain.NewLiteral(true, "B", 0), domain.NewLiteral(false, "C", 0)}))
	terms = append(terms, domain.NewTerm([]domain.Literal{domain.NewLiteral(true, "A", 0), domain.NewLiteral(false, "C", 0)}))

	domainObject := domain.NewDomain(terms)
	keyMap := domainObject.Solve()
	validationTerms := []domain.Term{}
	validationTerms = append(validationTerms, domain.NewTerm([]domain.Literal{domain.NewLiteral(false, "A", keyMap["A"]), domain.NewLiteral(false, "B", keyMap["B"]), domain.NewLiteral(false, "C", keyMap["C"])}))
	validationTerms = append(validationTerms, domain.NewTerm([]domain.Literal{domain.NewLiteral(true, "B", keyMap["B"]), domain.NewLiteral(false, "C", keyMap["C"])}))
	validationTerms = append(validationTerms, domain.NewTerm([]domain.Literal{domain.NewLiteral(true, "A", keyMap["A"]), domain.NewLiteral(false, "C", keyMap["C"])}))

	assert.Equal(t, true, domain.EvaluateTerms(validationTerms))

}
func TestSolveMoreMoreComplexStuffWithFixed(t *testing.T) {
	fmt.Println("start")
	terms := []domain.Term{}
	terms = append(terms, domain.NewTerm([]domain.Literal{domain.NewLiteral(false, "A", -1), domain.NewLiteral(false, "B", -1), domain.NewLiteral(false, "C", 0)}))
	terms = append(terms, domain.NewTerm([]domain.Literal{domain.NewLiteral(true, "B", -1), domain.NewLiteral(false, "C", 0)}))
	terms = append(terms, domain.NewTerm([]domain.Literal{domain.NewLiteral(true, "A", -1), domain.NewLiteral(false, "C", 0)}))

	domainObject := domain.NewDomain(terms)
	keyMap := domainObject.Solve()
	validationTerms := []domain.Term{}
	validationTerms = append(validationTerms, domain.NewTerm([]domain.Literal{domain.NewLiteral(false, "A", keyMap["A"]), domain.NewLiteral(false, "B", keyMap["B"]), domain.NewLiteral(false, "C", keyMap["C"])}))
	validationTerms = append(validationTerms, domain.NewTerm([]domain.Literal{domain.NewLiteral(true, "B", keyMap["B"]), domain.NewLiteral(false, "C", keyMap["C"])}))
	validationTerms = append(validationTerms, domain.NewTerm([]domain.Literal{domain.NewLiteral(true, "A", keyMap["A"]), domain.NewLiteral(false, "C", keyMap["C"])}))

	assert.Equal(t, true, domain.EvaluateTerms(validationTerms))
	assert.Equal(t, int8(-1), keyMap["A"])
	assert.Equal(t, int8(-1), keyMap["B"])
}
func TestNoSolution(t *testing.T) {
	fmt.Println("start")
	terms := []domain.Term{}
	terms = append(terms, domain.NewTerm([]domain.Literal{domain.NewLiteral(false, "A", 1), domain.NewLiteral(false, "B", 1), domain.NewLiteral(false, "C", 1)}))
	domainObject := domain.NewDomain(terms)
	keyMap := domainObject.Solve()
	assert.Equal(t, keyMap, make(map[string]int8))
}
