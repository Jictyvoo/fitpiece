package builder

import (
	"fmt"
	"github.com/wrapped-owls/fitpiece/cimenteiro/internal/utils"
	"strings"
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
func (expression SubQueryExpression) Build(writer utils.Writer) int {
	if expression.subQuery != nil {
		rawGen := RawSqlGenerator{Query: *expression.subQuery}
		length, _ := writer.WriteString(fmt.Sprintf("(%s)", rawGen.Select()))
		return length
	}
	return 0
}

// BuildPlaceholder generates a SQL with placeholders and a slice of values.
// The generated string and slice are both to be used together to prevent SQL-injection
func (expression SubQueryExpression) BuildPlaceholder(writer utils.Writer, placeholder string) (int, []any) {
	if expression.subQuery != nil {
		placeholderGen := PlaceholderSqlGenerator{Query: expression.subQuery, Placeholder: placeholder}
		sqlStr, args := placeholderGen.Select()

		length, _ := writer.WriteString(fmt.Sprintf("(%s)", sqlStr))
		return length, args
	}
	return 0, nil
}

func (expression SubQueryExpression) String() string {
	builder := strings.Builder{}
	expression.Build(&builder)
	return builder.String()
}

func (expression SubQueryExpression) StringPlaceholder(placeholder string) (string, []any) {
	builder := strings.Builder{}
	_, values := expression.BuildPlaceholder(&builder, placeholder)
	return builder.String(), values
}
