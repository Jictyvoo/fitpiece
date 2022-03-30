package elements

type Expression interface {

	// Build generates the string for the Expression as a raw SQL
	Build() string

	// BuildPlaceholder generates a SQL with placeholders and a slice of values.
	// The generated string and slice are both to be used together to prevent SQL-injection
	BuildPlaceholder(placeholder string) (string, []any)
}
