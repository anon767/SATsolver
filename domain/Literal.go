package domain

import "fmt"

type Literal struct {
	modifier  bool
	Name      string
	evaluated int8
}

func NewLiteral(modifier bool, name string, evaluated int8) Literal {
	return Literal{modifier, name, evaluated}
}

func (literal Literal) IsTrue() bool {
	if (literal.evaluated == 0) {
		panic(fmt.Sprintf("%s was not evaluated yet", literal.Name))
	}
	return literal.modifier == (literal.evaluated > 0)
}
func (literal Literal) IsVariable() bool {
	return literal.evaluated == 0
}

func (l Literal) Change(value bool) Literal {
	if value {
		return Literal{l.modifier, l.Name, 1}
	} else {
		return Literal{l.modifier, l.Name, -1}
	}
}
