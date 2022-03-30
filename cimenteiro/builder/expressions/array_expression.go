package expressions

import (
	"fmt"
	"github.com/wrapped-owls/fitpiece/cimenteiro/internal/utils"
	"strings"
)

// ArrayElementExpression defines how an array can be written in a SQL.
// To work well, a wrapIn value array can be set, to wrap the values inside
type ArrayElementExpression[T any] struct {
	values []T
	wrapIn [2]rune
}

// Build generates the string for the ArrayElementExpression as a raw SQL
func (expression ArrayElementExpression[T]) Build(writer utils.Writer) int {
	totalLength := 0
	if expression.wrapIn[0] > 0 {
		_, _ = writer.WriteRune(expression.wrapIn[0])
		totalLength += 1
	}
	for index, value := range expression.values {
		if index > 0 {
			_, _ = writer.WriteRune(',')
			_, _ = writer.WriteRune(' ')
			totalLength += 2
		}
		length, _ := writer.WriteString(fmt.Sprintf("%v", value))
		totalLength += length
	}
	if expression.wrapIn[1] > 0 {
		_, _ = writer.WriteRune(expression.wrapIn[1])
		totalLength += 1
	}
	return totalLength
}

// BuildPlaceholder generates a SQL with placeholders and a slice of values.
// The generated string and slice are both to be used together to prevent SQL-injection
func (expression ArrayElementExpression[T]) BuildPlaceholder(writer utils.Writer, placeholder string) (int, []any) {
	valuesList := make([]any, 0, len(expression.values))
	totalLength := 0

	if expression.wrapIn[0] > 0 {
		_, _ = writer.WriteRune(expression.wrapIn[0])
		totalLength += 1
	}
	for index, value := range expression.values {
		if index > 0 {
			_, _ = writer.WriteRune(',')
			_, _ = writer.WriteRune(' ')
			totalLength += 2
		}
		length, _ := writer.WriteString(placeholder)
		totalLength += length
		valuesList = append(valuesList, value)
	}
	if expression.wrapIn[1] > 0 {
		_, _ = writer.WriteRune(expression.wrapIn[1])
		totalLength += 1
	}
	return totalLength, valuesList
}

func (expression ArrayElementExpression[T]) String() string {
	builder := strings.Builder{}
	expression.Build(&builder)
	return builder.String()
}

func (expression ArrayElementExpression[T]) StringPlaceholder(placeholder string) (string, []any) {
	builder := strings.Builder{}
	_, values := expression.BuildPlaceholder(&builder, placeholder)
	return builder.String(), values
}
