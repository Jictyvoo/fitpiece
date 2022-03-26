package elements

import "fmt"

type operator string

const (
	OperatorPlus         operator = "+"
	OperatorMinus        operator = "-"
	OperatorEqual        operator = "="
	OperatorDifference   operator = "<>"
	OperatorGreaterThan  operator = ">"
	OperatorLessThan     operator = "<"
	OperatorGreaterEqual operator = ">="
	OperatorLessEqual    operator = "<="
)

type Clause struct {
	FirstHalf  Expression
	Operator   operator
	SecondHalf Expression
}

func (c Clause) Build() string {
	return fmt.Sprintf("%s %s %s", c.FirstHalf.Build(), c.Operator, c.SecondHalf.Build())
}
