package builder

import (
	"fmt"
	"github.com/wrapped-owls/fitpiece/cimenteiro/internal/elements"
	"strings"
)

type ArrayElementExpression[T any] struct {
	subQuery *QueryBuilder
	values   []T
}

func SubQuery[T any](query *QueryBuilder) ArrayElementExpression[T] {
	return ArrayElementExpression[T]{
		subQuery: query,
	}
}

func ArrayExpression[T any](values ...T) ArrayElementExpression[T] {
	return ArrayElementExpression[T]{
		values: values,
	}
}

func NewFieldExpression(name string) elements.Expression {
	return elements.FieldExpression{Name: name}
}

func (expression ArrayElementExpression[T]) Build() string {
	if expression.subQuery != nil {
		rawGen := RawSqlGenerator{Query: *expression.subQuery}
		return fmt.Sprintf("(%s)", rawGen.Select())
	}
	builder := strings.Builder{}
	builder.WriteRune('[')
	for index, value := range expression.values {
		if index > 0 {
			builder.WriteRune(',')
		}
		builder.WriteString(fmt.Sprintf("%v", value))
	}
	builder.WriteRune(']')
	return builder.String()
}

func (expression ArrayElementExpression[T]) BuildPlaceholder(placeholder string) (string, []any) {
	if expression.subQuery != nil {
		placeholderGen := PlaceholderSqlGenerator{Query: expression.subQuery, Placeholder: placeholder}
		sqlStr, args := placeholderGen.Select()
		return fmt.Sprintf("(%s)", sqlStr), args
	}
	builder := strings.Builder{}
	valuesList := make([]any, 0, len(expression.values))
	builder.WriteRune('[')
	for index, value := range expression.values {
		if index > 0 {
			builder.WriteRune(',')
		}
		builder.WriteString(placeholder)
		valuesList = append(valuesList, value)
	}
	builder.WriteRune(']')
	return builder.String(), valuesList
}

type ValueExpression[T any] struct {
	value T
}

func (expression ValueExpression[T]) Build() string {
	return fmt.Sprintf("%v", expression.value)
}

func (expression ValueExpression[T]) BuildPlaceholder(placeholder string) (string, []any) {
	return placeholder, []any{expression.value}
}

type RawClauseExpression = ValueExpression[string]

func RawExpression(expression string) RawClauseExpression {
	return RawClauseExpression{expression}
}

type PrefixClauseExpression struct {
	prefix string
	value  elements.Expression
}

func NewPrefixExpression(prefix string, value elements.Expression) PrefixClauseExpression {
	return PrefixClauseExpression{prefix: prefix, value: value}
}

func (expression PrefixClauseExpression) Build() string {
	return fmt.Sprintf("%s %s", expression.prefix, expression.value.Build())
}

func (expression PrefixClauseExpression) BuildPlaceholder(placeholder string) (string, []any) {
	sqlStr, args := expression.value.BuildPlaceholder(placeholder)
	return fmt.Sprintf("%s %s", expression.prefix, sqlStr), args
}
