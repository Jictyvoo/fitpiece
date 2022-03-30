package elements

import (
	"fmt"
	"strings"
)

type (
	operator string
	keyword  = operator
)

const (
	// Operators

	OperatorEqual        operator = "="
	OperatorDifference   operator = "<>"
	OperatorGreaterThan  operator = ">"
	OperatorLessThan     operator = "<"
	OperatorGreaterEqual operator = ">="
	OperatorLessEqual    operator = "<="

	// Keywords

	KeywordNotIn keyword = "NOT IN"
	KeywordIn    keyword = "IN"
	KeywordAnd   keyword = "AND"
	KeywordOr    keyword = "OR"
)

type Clause struct {
	FirstHalf  Expression
	Operator   operator
	SecondHalf Expression
}

func (c Clause) Build() string {
	return fmt.Sprintf("%s %s %s", c.FirstHalf.Build(), c.Operator, c.SecondHalf.Build())
}

func (c Clause) BuildPlaceholder(placeholder string) (string, []any) {
	valueList := make([]any, 0, 2)
	stringBuilder := strings.Builder{}

	// Checker for brackets
	writeBrackets := c.Operator == KeywordAnd || c.Operator == KeywordOr
	if writeBrackets {
		stringBuilder.WriteString("(")
	}

	// Write first half
	strResult, argsResult := c.FirstHalf.BuildPlaceholder(placeholder)
	stringBuilder.WriteString(strResult)
	valueList = append(valueList, argsResult...)

	// Write operator
	if len(c.Operator) > 0 {
		stringBuilder.WriteRune(' ')
		stringBuilder.WriteString(string(c.Operator))
	}
	stringBuilder.WriteRune(' ')

	// Write second half
	strResult, argsResult = c.SecondHalf.BuildPlaceholder(placeholder)
	stringBuilder.WriteString(strResult)
	valueList = append(valueList, argsResult...)

	// Close brackets
	if writeBrackets {
		stringBuilder.WriteRune(')')
	}
	return stringBuilder.String(), valueList
}
