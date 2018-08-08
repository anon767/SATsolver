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
			seperator := strings.Split(word, ":")
			modifier := !strings.HasPrefix(seperator[0], "-")
			name := seperator[0]
			if !modifier {
				name = strings.Replace(seperator[0], "-", "", -1)
			}
			var evaluation int8 = 0
			if seperator[1] == "-1" {
				evaluation = -1
			} else if seperator[1] == "1" {
				evaluation = 1
			}
			literals = append(literals, domain.NewLiteral(modifier, name, evaluation))
		}
		terms = append(terms, domain.NewTerm(literals))
	}
	return domain.NewDomain(terms)
}
