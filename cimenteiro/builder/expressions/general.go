package expressions

import (
	"github.com/wrapped-owls/fitpiece/cimenteiro/internal/elements"
)

func ArrayExpression[T any](values ...T) ArrayElementExpression[T] {
	return ArrayElementExpression[T]{
		values: values,
	}
}

func RawExpression(expression string) RawClauseExpression {
	return RawClauseExpression{expression}
}

func NewPrefixExpression(prefix string, value elements.Expression) PrefixClauseExpression {
	return PrefixClauseExpression{prefix: prefix, value: value}
}

func NewFieldExpression(name string) elements.Expression {
	return elements.FieldExpression{Name: name}
}

func NewValueExpression[T any](value T) elements.Expression {
	return ValueExpression[T]{value: value}
}
