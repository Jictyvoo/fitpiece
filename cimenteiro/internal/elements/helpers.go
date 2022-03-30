package elements

import "github.com/wrapped-owls/fitpiece/cimenteiro/internal/utils"

type Expression interface {

	// Build generates the string for the Expression as a raw SQL
	Build(writer utils.Writer) int

	// BuildPlaceholder generates a SQL with placeholders and a slice of values.
	// The generated string and slice are both to be used together to prevent SQL-injection
	BuildPlaceholder(writer utils.Writer, placeholder string) (int, []any)

	// String runs the Build method and returns the string directly.
	// Useful for tests purposes
	String() string

	// StringPlaceholder runs the BuildPlaceholder and returns both string and slice values
	StringPlaceholder(placeholder string) (string, []any)
}
