package expressions

import (
	"github.com/wrapped-owls/fitpiece/cimenteiro/internal/utils"
	"strings"
)

type ValueFieldType interface {
	~string | ~int | ~uint
}

// DirectValueArrayExpression defines how an array can be written in a SQL.
// The unique difference from ArrayElementExpression is that the BuildPlaceholder method will generate an empty slice,
// and not put the placeholder in the SQL string
type DirectValueArrayExpression[T ValueFieldType] struct {
	values []T
	wrapIn [2]rune
}

// Build generates the string for the DirectValueArrayExpression as a raw SQL
func (expression DirectValueArrayExpression[T]) Build(writer utils.Writer) int {
	return ArrayElementExpression[T]{
		values: expression.values,
		wrapIn: expression.wrapIn,
	}.Build(writer)
}

// BuildPlaceholder generates a SQL with placeholders and a slice of values.
// The generated string and slice are both to be used together to prevent SQL-injection
func (expression DirectValueArrayExpression[T]) BuildPlaceholder(writer utils.Writer, placeholder string) (int, []any) {
	return expression.Build(writer), []any{}
}

func (expression DirectValueArrayExpression[T]) String() string {
	builder := strings.Builder{}
	expression.Build(&builder)
	return builder.String()
}

func (expression DirectValueArrayExpression[T]) StringPlaceholder(placeholder string) (string, []any) {
	builder := strings.Builder{}
	_, values := expression.BuildPlaceholder(&builder, placeholder)
	return builder.String(), values
}
