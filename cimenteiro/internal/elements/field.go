package elements

import (
	"github.com/wrapped-owls/fitpiece/cimenteiro/internal/utils"
	"strings"
)

type FieldExpression struct {
	Name  string
	Alias string
}

func (expression FieldExpression) Build(writer utils.Writer) int {
	if len(expression.Alias) <= 0 {
		length, _ := writer.WriteString(expression.Name)
		return length
	}

	totalLength := 0
	length, _ := writer.WriteString(expression.Name)
	totalLength += length

	length, _ = writer.WriteString(" AS ")
	totalLength += length

	length, _ = writer.WriteString(expression.Alias)

	return totalLength + length
}

func (expression FieldExpression) BuildPlaceholder(writer utils.Writer, placeholder string) (int, []any) {
	return expression.Build(writer), []any{}
}

func (expression FieldExpression) String() string {
	builder := strings.Builder{}
	expression.Build(&builder)
	return builder.String()
}

func (expression FieldExpression) StringPlaceholder(placeholder string) (string, []any) {
	builder := strings.Builder{}
	_, values := expression.BuildPlaceholder(&builder, placeholder)
	return builder.String(), values
}
