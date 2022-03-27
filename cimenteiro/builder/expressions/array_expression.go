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
	builder.WriteRune(expression.wrapIn[0])
	for index, value := range expression.values {
		if index > 0 {
			builder.WriteRune(',')
		}
		builder.WriteString(fmt.Sprintf("%v", value))
	}
	builder.WriteRune(expression.wrapIn[1])
	return builder.String()
}

func (expression ArrayElementExpression[T]) BuildPlaceholder(placeholder string) (string, []any) {
	builder := strings.Builder{}
	valuesList := make([]any, 0, len(expression.values))
	builder.WriteRune(expression.wrapIn[0])
	for index, value := range expression.values {
		if index > 0 {
			builder.WriteRune(',')
		}
		builder.WriteString(placeholder)
		valuesList = append(valuesList, value)
	}
	builder.WriteRune(expression.wrapIn[1])
	return builder.String(), valuesList
}
