package elements

import (
	"strings"

	"github.com/jictyvoo/fitpiece/cimenteiro/internal/utils"
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

func (c Clause) Build(writer utils.Writer) int {
	if firstLength := c.FirstHalf.Build(writer); firstLength > 0 {
		_, _ = writer.WriteRune(' ')
		_, _ = writer.WriteString(string(c.Operator))
		_, _ = writer.WriteRune(' ')
		secondLength := c.SecondHalf.Build(writer)
		return firstLength + 2 + len(c.Operator) + secondLength
	}
	return 0
}

func (c Clause) BuildPlaceholder(writer utils.Writer, placeholder string) (int, []any) {
	valueList := make([]any, 0, 2)
	totalWriteLength := 0

	// Checker for brackets
	writeBrackets := c.Operator == KeywordAnd || c.Operator == KeywordOr
	if writeBrackets {
		_, _ = writer.WriteString("(")
		totalWriteLength += 1
	}

	// Write first half
	lengthResult, argsResult := c.FirstHalf.BuildPlaceholder(writer, placeholder)
	totalWriteLength += lengthResult
	valueList = append(valueList, argsResult...)

	// Write operator
	if len(c.Operator) > 0 {
		_, _ = writer.WriteRune(' ')
		lengthResult, _ = writer.WriteString(string(c.Operator))
		totalWriteLength += lengthResult + 1
	}
	_, _ = writer.WriteRune(' ')
	totalWriteLength += 1

	// Write second half
	lengthResult, argsResult = c.SecondHalf.BuildPlaceholder(writer, placeholder)
	totalWriteLength += lengthResult
	valueList = append(valueList, argsResult...)

	// Close brackets
	if writeBrackets {
		_, _ = writer.WriteRune(')')
		totalWriteLength += 1
	}
	return totalWriteLength, valueList
}

func (c Clause) String() string {
	builder := strings.Builder{}
	c.Build(&builder)
	return builder.String()
}

func (c Clause) StringPlaceholder(placeholder string) (string, []any) {
	builder := strings.Builder{}
	_, values := c.BuildPlaceholder(&builder, placeholder)
	return builder.String(), values
}
