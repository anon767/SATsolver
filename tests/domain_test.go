package tests

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"satsolver/domain"
	"fmt"
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
	solvedTerm := domain.NewLiteral(true, "A", 1)
	domain := domain.NewDomain(terms)
	assert.Equal(t, solvedTerm, domain.Solve()[0])
}
func TestSolveMoreTwoTerm(t *testing.T) {
	terms := []domain.Term{}
	terms = append(terms, domain.NewTerm([]domain.Literal{domain.NewLiteral(true, "A", 0)}))
	terms = append(terms, domain.NewTerm([]domain.Literal{domain.NewLiteral(true, "B", 0)}))
	solvedTerm := []domain.Literal{}
	solvedTerm = append(solvedTerm, domain.NewLiteral(true, "A", 1))
	solvedTerm = append(solvedTerm, domain.NewLiteral(true, "B", 1))
	domain := domain.NewDomain(terms)
	assert.ElementsMatch(t, solvedTerm, domain.Solve())
}
func TestSolveMoreTwoTermMultipleLiterals(t *testing.T) {
	terms := []domain.Term{}
	terms = append(terms, domain.NewTerm([]domain.Literal{domain.NewLiteral(true, "A", 0), domain.NewLiteral(false, "B", 0)}))
	terms = append(terms, domain.NewTerm([]domain.Literal{domain.NewLiteral(true, "B", 0)}))
	solvedTerm := []domain.Literal{}
	solvedTerm = append(solvedTerm, domain.NewLiteral(true, "A", 1))
	solvedTerm = append(solvedTerm, domain.NewLiteral(true, "B", 1))
	domain := domain.NewDomain(terms)
	assert.ElementsMatch(t, solvedTerm, domain.Solve())
}
func TestSolveMoreMoreComplexStuff(t *testing.T) {
	fmt.Println("start")
	terms := []domain.Term{}
	terms = append(terms, domain.NewTerm([]domain.Literal{domain.NewLiteral(false, "A", 0), domain.NewLiteral(false, "B", 0), domain.NewLiteral(false, "C", 0)}))
	terms = append(terms, domain.NewTerm([]domain.Literal{domain.NewLiteral(true, "B", 0), domain.NewLiteral(false, "C", 0)}))
	terms = append(terms, domain.NewTerm([]domain.Literal{domain.NewLiteral(true, "A", 0), domain.NewLiteral(false, "C", 0)}))
	solvedTerm := []domain.Literal{}
	solvedTerm = append(solvedTerm, domain.NewLiteral(true, "A", 1))
	solvedTerm = append(solvedTerm, domain.NewLiteral(true, "B", 1))
	solvedTerm = append(solvedTerm, domain.NewLiteral(true, "C", -1))
	domain := domain.NewDomain(terms)

	fmt.Println(domain.Solve())
	assert.ElementsMatch(t, solvedTerm, domain.Solve())
}
