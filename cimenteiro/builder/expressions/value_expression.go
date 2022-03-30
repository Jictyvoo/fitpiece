package expressions

import (
	"fmt"
	"github.com/wrapped-owls/fitpiece/cimenteiro/internal/utils"
	"strings"
)

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
func (expression ValueExpression[T]) Build(writer utils.Writer) int {
	length, _ := writer.WriteString(fmt.Sprintf("%v", expression.value))
	return length
}

// BuildPlaceholder generates a SQL with placeholders and a slice of values.
// The generated string and slice are both to be used together to prevent SQL-injection
func (expression ValueExpression[T]) BuildPlaceholder(writer utils.Writer, placeholder string) (int, []any) {
	length, _ := writer.WriteString(placeholder)
	return length, []any{expression.value}
}

func (expression ValueExpression[T]) String() string {
	builder := strings.Builder{}
	expression.Build(&builder)
	return builder.String()
}

func (expression ValueExpression[T]) StringPlaceholder(placeholder string) (string, []any) {
	builder := strings.Builder{}
	_, values := expression.BuildPlaceholder(&builder, placeholder)
	return builder.String(), values
}
