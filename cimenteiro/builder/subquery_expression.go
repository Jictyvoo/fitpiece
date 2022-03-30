package builder

import (
	"fmt"
)

// SubQueryExpression is the structure expression for sub-queries that implement elements.Expression
type SubQueryExpression struct {
	subQuery *QueryBuilder
}

// SubQuery creates a new SubQueryExpression with the given QueryBuilder
func SubQuery(query *QueryBuilder) SubQueryExpression {
	return SubQueryExpression{
		subQuery: query,
	}
}

// Build generates the string for the subQuery as a raw SQL
func (expression SubQueryExpression) Build() string {
	if expression.subQuery != nil {
		rawGen := RawSqlGenerator{Query: *expression.subQuery}
		return fmt.Sprintf("(%s)", rawGen.Select())
	}
	return ""
}

// BuildPlaceholder generates a SQL with placeholders and a slice of values.
// The generated string and slice are both to be used together to prevent SQL-injection
func (expression SubQueryExpression) BuildPlaceholder(placeholder string) (string, []any) {
	if expression.subQuery != nil {
		placeholderGen := PlaceholderSqlGenerator{Query: expression.subQuery, Placeholder: placeholder}
		sqlStr, args := placeholderGen.Select()
		return fmt.Sprintf("(%s)", sqlStr), args
	}
	return "", nil
}
