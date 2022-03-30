package expressions

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
func (expression DirectValueArrayExpression[T]) Build() string {
	return ArrayElementExpression[T]{
		values: expression.values,
		wrapIn: expression.wrapIn,
	}.Build()
}

// BuildPlaceholder generates a SQL with placeholders and a slice of values.
// The generated string and slice are both to be used together to prevent SQL-injection
func (expression DirectValueArrayExpression[T]) BuildPlaceholder(placeholder string) (string, []any) {
	return expression.Build(), []any{}
}
