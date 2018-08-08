package tests

import (
	"../domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTermTruthy(t *testing.T) {
	term := domain.Term{}
	assert.Equal(t, 0, term.GetSize())
	term = domain.NewTerm([]domain.Literal{domain.NewLiteral(true, "A", 1), domain.NewLiteral(true, "B", 1), domain.NewLiteral(false, "C", 1)})
	assert.Equal(t, 3, term.GetSize())
	assert.Equal(t, true, term.EvaluateTerm())
}
func TestTermFalsy(t *testing.T) {
	term := domain.Term{}
	assert.Equal(t, 0, term.GetSize())
	term = domain.NewTerm([]domain.Literal{domain.NewLiteral(true, "A", -1), domain.NewLiteral(true, "B", -1), domain.NewLiteral(false, "C", -1)})
	assert.Equal(t, 3, term.GetSize())
	assert.Equal(t, true, term.EvaluateTerm())
}

func TestPermutation(t *testing.T) {
	term := domain.NewTerm([]domain.Literal{domain.NewLiteral(true, "A", 0), domain.NewLiteral(true, "B", 0), domain.NewLiteral(true, "C", 0)})
	assert.Equal(t, len(term.Permutate()), 8)
}

func TestFixedPermutation(t *testing.T) {
	term := domain.NewTerm([]domain.Literal{domain.NewLiteral(true, "A", 1), domain.NewLiteral(true, "B", -1), domain.NewLiteral(true, "C", 0)})
	assert.Equal(t, len(term.Permutate()), 2)
}

func TestWorkingPermutation(t *testing.T) {
	term := domain.NewTerm([]domain.Literal{domain.NewLiteral(true, "A", 0), domain.NewLiteral(true, "B", 0), domain.NewLiteral(true, "C", 0)})
	assert.Equal(t, len(term.GetGoodPermutations()), 7)
}
