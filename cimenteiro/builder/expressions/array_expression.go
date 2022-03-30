package expressions

import (
	"fmt"
	"strings"
)

// ArrayElementExpression defines how an array can be written in a SQL.
// To work well, a wrapIn value array can be set, to wrap the values inside
type ArrayElementExpression[T any] struct {
	values []T
	wrapIn [2]rune
}

// Build generates the string for the ArrayElementExpression as a raw SQL
func (expression ArrayElementExpression[T]) Build() string {
	builder := strings.Builder{}
	if expression.wrapIn[0] > 0 {
		builder.WriteRune(expression.wrapIn[0])
	}
	for index, value := range expression.values {
		if index > 0 {
			builder.WriteRune(',')
			builder.WriteRune(' ')
		}
		builder.WriteString(fmt.Sprintf("%v", value))
	}
	if expression.wrapIn[1] > 0 {
		builder.WriteRune(expression.wrapIn[1])
	}
	return builder.String()
}

// BuildPlaceholder generates a SQL with placeholders and a slice of values.
// The generated string and slice are both to be used together to prevent SQL-injection
func (expression ArrayElementExpression[T]) BuildPlaceholder(placeholder string) (string, []any) {
	builder := strings.Builder{}
	valuesList := make([]any, 0, len(expression.values))
	if expression.wrapIn[0] > 0 {
		builder.WriteRune(expression.wrapIn[0])
	}
	for index, value := range expression.values {
		if index > 0 {
			builder.WriteRune(',')
			builder.WriteRune(' ')
		}
		builder.WriteString(placeholder)
		valuesList = append(valuesList, value)
	}
	if expression.wrapIn[1] > 0 {
		builder.WriteRune(expression.wrapIn[1])
	}
	return builder.String(), valuesList
}
