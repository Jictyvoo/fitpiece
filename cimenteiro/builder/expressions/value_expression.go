package expressions

import "fmt"

type (
	ValueExpression[T any] struct {
		value T
	}
	RawClauseExpression = ValueExpression[string]
)

func (expression ValueExpression[T]) Build() string {
	return fmt.Sprintf("%v", expression.value)
}

func (expression ValueExpression[T]) BuildPlaceholder(placeholder string) (string, []any) {
	return placeholder, []any{expression.value}
}
