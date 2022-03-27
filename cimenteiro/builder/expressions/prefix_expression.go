package expressions

import (
	"fmt"
	"github.com/wrapped-owls/fitpiece/cimenteiro/internal/elements"
)

type PrefixClauseExpression struct {
	prefix string
	value  elements.Expression
}

func (expression PrefixClauseExpression) Build() string {
	return fmt.Sprintf("%s %s", expression.prefix, expression.value.Build())
}

func (expression PrefixClauseExpression) BuildPlaceholder(placeholder string) (string, []any) {
	sqlStr, args := expression.value.BuildPlaceholder(placeholder)
	return fmt.Sprintf("%s %s", expression.prefix, sqlStr), args
}
