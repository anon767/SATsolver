package tests

import (
	"../parser"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLineSplit(t *testing.T) {
	assert.Equal(t, []string{"a", "b"}, parser.NewLineSplit("a\nb"))
}
func TestParser(t *testing.T) {
	p := parser.Parser{FilePath: "./cnfTestFile.cnf"}
	parsedDomain := p.ParseFile()
	solvedTerm := map[string]int8{"C": 1, "B": -1, "A": 1}
	assert.Equal(t, solvedTerm, parsedDomain.Solve())
}
func TestSudoku(t *testing.T) {
	p := parser.Parser{FilePath: "./sudoku.cnf"}
	parsedDomain := p.ParseFile()
	fmt.Println(parsedDomain.Solve())
}
