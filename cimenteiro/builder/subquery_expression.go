package builder

import (
	"fmt"
)

type SubQueryExpression struct {
	subQuery *QueryBuilder
}

func SubQuery(query *QueryBuilder) SubQueryExpression {
	return SubQueryExpression{
		subQuery: query,
	}
}

func (expression SubQueryExpression) Build() string {
	if expression.subQuery != nil {
		rawGen := RawSqlGenerator{Query: *expression.subQuery}
		return fmt.Sprintf("(%s)", rawGen.Select())
	}
	return ""
}

func (expression SubQueryExpression) BuildPlaceholder(placeholder string) (string, []any) {
	if expression.subQuery != nil {
		placeholderGen := PlaceholderSqlGenerator{Query: expression.subQuery, Placeholder: placeholder}
		sqlStr, args := placeholderGen.Select()
		return fmt.Sprintf("(%s)", sqlStr), args
	}
	return "", nil
}
