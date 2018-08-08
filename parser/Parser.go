package parser

import (
	"../domain"
	"io/ioutil"
	"strings"
)

type Parser struct {
	FilePath string
}

func NewLineSplit(input string) []string {
	return strings.Split(input, "\n")
}

func (parser Parser) ParseFile() domain.Domain {
	contents, err := ioutil.ReadFile(parser.FilePath)
	if err != nil {
		panic(err)
	}
	lines := NewLineSplit(string(contents))
	var terms []domain.Term
	for _, line := range lines {
		wordsPerLine := strings.Split(line, " ")
		var literals []domain.Literal
		for _, word := range wordsPerLine {
			modifier := !strings.HasPrefix(word, "-")
			name := word
			if !modifier {
				name = strings.Replace(word, "-", "", -1)
			}
			literals = append(literals, domain.NewLiteral(modifier, name, 0))
		}
		terms = append(terms, domain.NewTerm(literals))
	}
	return domain.NewDomain(terms)
}
