package expressions

import (
	"fmt"
	"github.com/jictyvoo/fitpiece/cimenteiro/internal/elements"
)

type PairsClauseExpression struct {
	first  elements.Expression
	second elements.Expression
}

func (expression PairsClauseExpression) Build() string {
	return fmt.Sprintf("%s %s", expression.first.Build(), expression.second.Build())
}

func (expression PairsClauseExpression) BuildPlaceholder(placeholder string) (string, []any) {
	firstSql, fArgs := expression.first.BuildPlaceholder(placeholder)
	secondSql, sArgs := expression.second.BuildPlaceholder(placeholder)
	return fmt.Sprintf("%s %s", firstSql, secondSql), append(fArgs, sArgs...)
}
