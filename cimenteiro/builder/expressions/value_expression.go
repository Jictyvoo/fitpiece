package expressions

import "fmt"

type ValueExpression[T any] struct {
	value T
}

func (expression ValueExpression[T]) Build() string {
	return fmt.Sprintf("%v", expression.value)
}

func (expression ValueExpression[T]) BuildPlaceholder(placeholder string) (string, []any) {
	return placeholder, []any{expression.value}
}

type RawClauseExpression = ValueExpression[string]

func RawExpression(expression string) RawClauseExpression {
	return RawClauseExpression{expression}
}
