package expressions

import (
	"fmt"
	"github.com/jictyvoo/fitpiece/cimenteiro/internal/elements"
)

// PairsClauseExpression defines an easy way to write expressions that are not divided by an operator
type PairsClauseExpression struct {
	first  elements.Expression
	second elements.Expression
}

// Build generates the string for the PairsClauseExpression as a raw SQL
func (expression PairsClauseExpression) Build() string {
	return fmt.Sprintf("%s %s", expression.first.Build(), expression.second.Build())
}

// BuildPlaceholder generates a SQL with placeholders and a slice of values.
// The generated string and slice are both to be used together to prevent SQL-injection
func (expression PairsClauseExpression) BuildPlaceholder(placeholder string) (string, []any) {
	firstSql, fArgs := expression.first.BuildPlaceholder(placeholder)
	secondSql, sArgs := expression.second.BuildPlaceholder(placeholder)
	return fmt.Sprintf("%s %s", firstSql, secondSql), append(fArgs, sArgs...)
}
