package tests

import (
	"../domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLiteral(t *testing.T) {
	literal := domain.NewLiteral(true, "A", -1)
	assert.Equal(t, false, literal.IsTrue())
	literal = literal.Change(true)
	assert.Equal(t, true, literal.IsTrue())
}

func TestLiteralEvaluationPanic(t *testing.T) {
	literal := domain.NewLiteral(true, "A", 0)
	assert.Panics(t, func() {
		literal.IsTrue()
	})
}
