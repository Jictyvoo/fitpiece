package elements

import "fmt"

type operator string

const (
	OperatorEqual        operator = "="
	OperatorDifference   operator = "<>"
	OperatorGreaterThan  operator = ">"
	OperatorLessThan     operator = "<"
	OperatorGreaterEqual operator = ">="
	OperatorLessEqual    operator = "<="
	OperatorNotIn        operator = "NOT IN"
	OperatorIn           operator = "IN"
)

type Clause struct {
	FirstHalf  Expression
	Operator   operator
	SecondHalf Expression
}

func (c Clause) Build() string {
	return fmt.Sprintf("%s %s %s", c.FirstHalf.Build(), c.Operator, c.SecondHalf.Build())
}
