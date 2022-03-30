package expressions

import (
	"github.com/jictyvoo/fitpiece/cimenteiro/internal/elements"
)

// ArrayExpression create a new ArrayElementExpression object with values that will be wrapped in '[' ']'
func ArrayExpression[T any](values ...T) ArrayElementExpression[T] {
	return ArrayElementExpression[T]{
		values: values,
		wrapIn: [2]rune{'[', ']'},
	}
}

// MultiDirectValueExpression create a new DirectValueArrayExpression with given wrappers and values.
// If given wrappers are equal to 0, then the values will not be wrapped
func MultiDirectValueExpression[T ValueFieldType](wrappers [2]rune, values ...T) DirectValueArrayExpression[T] {
	return DirectValueArrayExpression[T]{values: values, wrapIn: wrappers}
}

// RawExpression create a new RawClauseExpression using the given SQL string
func RawExpression(expression string) RawClauseExpression {
	return RawClauseExpression{expression}
}

// NewFieldExpression create a Field expression, that field can receive an Alias, and when generated
// will put the value of field directly, without any placeholder
func NewFieldExpression(name string) elements.Expression {
	return elements.FieldExpression{Name: name}
}

// NewValueExpression create a ValueExpression that has the same type and value given as parameter
func NewValueExpression[T any](value T) elements.Expression {
	return ValueExpression[T]{value: value}
}

// NewPairsClauseExpression create a PairsClauseExpression using both the elements.Expression given
func NewPairsClauseExpression(first, second elements.Expression) PairsClauseExpression {
	return PairsClauseExpression{first: first, second: second}
}

// PrefixExpression create a new PairsClauseExpression with a string prefix
func PrefixExpression(prefix string, value elements.Expression) PairsClauseExpression {
	return PairsClauseExpression{
		first:  NewFieldExpression(prefix),
		second: value,
	}
}

// SuffixExpression create a new PairsClauseExpression with a string suffix
func SuffixExpression(prefix elements.Expression, suffix string) PairsClauseExpression {
	return PairsClauseExpression{
		first:  prefix,
		second: NewFieldExpression(suffix),
	}
}
