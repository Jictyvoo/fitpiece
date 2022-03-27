package expressions

import (
	"fmt"
	"strings"
)

type ArrayElementExpression[T any] struct {
	values []T
	wrapIn [2]rune
}

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
