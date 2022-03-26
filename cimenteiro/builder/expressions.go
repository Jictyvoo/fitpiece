package builder

import (
	"fmt"
	"strings"
)

type ArrayElementExpression[T any] struct {
	subQuery *QueryBuilder
	values   []T
}

func SubQuery[T any](query *QueryBuilder) ArrayElementExpression[T] {
	return ArrayElementExpression[T]{
		subQuery: query,
	}
}

func ArrayExpression[T any](values ...T) ArrayElementExpression[T] {
	return ArrayElementExpression[T]{
		values: values,
	}
}

func (expression ArrayElementExpression[T]) Build() string {
	if expression.subQuery != nil {
		rawGen := RawSqlGenerator{Query: *expression.subQuery}
		return fmt.Sprintf("(%s)", rawGen.Select())
	}
	builder := strings.Builder{}
	builder.WriteRune('[')
	for index, value := range expression.values {
		if index > 0 {
			builder.WriteRune(',')
		}
		builder.WriteString(fmt.Sprintf("%v", value))
	}
	builder.WriteRune(']')
	return builder.String()
}

type ValueExpression[T any] struct {
	value T
}

func (expression ValueExpression[T]) Build() string {
	return fmt.Sprintf("%v", expression.value)
}

type RawClauseExpression = ValueExpression[string]

func RawExpression(expression string) RawClauseExpression {
	return RawClauseExpression{expression}
}
