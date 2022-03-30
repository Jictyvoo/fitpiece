package expressions

import "fmt"

type (
	// ValueExpression generates an expression using a single value of any type for it
	ValueExpression[T any] struct {
		value T
	}

	// RawClauseExpression specify ValueExpression for a string value.
	// TODO: Remove it from ValueExpression, and create a new struct
	RawClauseExpression = ValueExpression[string]
)

// Build generates the string for the ValueExpression as a raw SQL
func (expression ValueExpression[T]) Build() string {
	return fmt.Sprintf("%v", expression.value)
}

// BuildPlaceholder generates a SQL with placeholders and a slice of values.
// The generated string and slice are both to be used together to prevent SQL-injection
func (expression ValueExpression[T]) BuildPlaceholder(placeholder string) (string, []any) {
	return placeholder, []any{expression.value}
}
