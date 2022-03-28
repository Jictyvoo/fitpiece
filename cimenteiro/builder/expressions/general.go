package expressions

import (
	"github.com/jictyvoo/fitpiece/cimenteiro/internal/elements"
)

func ArrayExpression[T any](values ...T) ArrayElementExpression[T] {
	return ArrayElementExpression[T]{
		values: values,
		wrapIn: [2]rune{'[', ']'},
	}
}

func MultiDirectValueExpression[T ValueFieldType](wrappers [2]rune, values ...T) DirectValueArrayExpression[T] {
	return DirectValueArrayExpression[T]{values: values, wrapIn: wrappers}
}

func RawExpression(expression string) RawClauseExpression {
	return RawClauseExpression{expression}
}
func NewFieldExpression(name string) elements.Expression {
	return elements.FieldExpression{Name: name}
}

func NewValueExpression[T any](value T) elements.Expression {
	return ValueExpression[T]{value: value}
}

func NewPairsClauseExpression(first, second elements.Expression) PairsClauseExpression {
	return PairsClauseExpression{first: first, second: second}
}

func PrefixExpression(prefix string, value elements.Expression) PairsClauseExpression {
	return PairsClauseExpression{
		first:  NewFieldExpression(prefix),
		second: value,
	}
}

func SuffixExpression(prefix elements.Expression, suffix string) PairsClauseExpression {
	return PairsClauseExpression{
		first:  prefix,
		second: NewFieldExpression(suffix),
	}
}
