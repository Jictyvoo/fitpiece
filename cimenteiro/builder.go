package cimenteiro

import (
	"github.com/Wrapped-Owls/fitpiece/cimenteiro/internal/elements"
	"strings"
)

type QueryBuilder struct {
	tableName  elements.TableName
	fields     []elements.FieldExpression
	joins      []elements.JoinClause
	where      elements.Expression
	organizers []elements.Expression
}

func (builder *QueryBuilder) Fields(fields ...string) *QueryBuilder {
	index := len(builder.fields)
	// Extends the slice cap, to prevent reallocate inside the loop
	newCap := len(fields) + index
	builder.fields = builder.fields[:newCap]
	for _, field := range fields {
		builder.fields[index] = elements.FieldExpression{Name: field, Alias: ""}
		index++
	}
	return builder
}

// FIXME: Add parameter for first tableName (to not use builder.tableName)
func (builder *QueryBuilder) buildJoin(
	tableName elements.TableName, firstColumn string, secondColumn string,
) elements.JoinClause {
	return elements.JoinClause{
		JoinType: elements.JoinAll,
		Table:    tableName,
		On: elements.Clause{
			FirstHalf:  elements.FieldExpression{Name: builder.tableName.Column(firstColumn)},
			Operator:   elements.OperatorEqual,
			SecondHalf: elements.FieldExpression{Name: tableName.Column(secondColumn)},
		},
	}
}

func (builder *QueryBuilder) LeftJoin(tableName elements.TableName, firstColumn string, secondColumn string) *QueryBuilder {
	leftJoin := builder.buildJoin(tableName, firstColumn, secondColumn)
	leftJoin.JoinType = elements.JoinLeft
	builder.joins = append(builder.joins, leftJoin)
	return builder
}

func (builder *QueryBuilder) RightJoin(tableName elements.TableName, firstColumn string, secondColumn string) *QueryBuilder {
	rightJoin := builder.buildJoin(tableName, firstColumn, secondColumn)
	rightJoin.JoinType = elements.JoinRight
	builder.joins = append(builder.joins, rightJoin)
	return builder
}

func (builder *QueryBuilder) InnerJoin(tableName elements.TableName, firstColumn string, secondColumn string) *QueryBuilder {
	innerJoin := builder.buildJoin(tableName, firstColumn, secondColumn)
	innerJoin.JoinType = elements.JoinInner
	builder.joins = append(builder.joins, innerJoin)
	return builder
}

func (builder *QueryBuilder) OuterJoin(tableName elements.TableName, firstColumn string, secondColumn string) *QueryBuilder {
	outerJoin := builder.buildJoin(tableName, firstColumn, secondColumn)
	outerJoin.JoinType = elements.JoinOuter
	builder.joins = append(builder.joins, outerJoin)
	return builder
}

func (builder QueryBuilder) Update(values map[string]string) string {
	sqlCommand := strings.Builder{}
	sqlCommand.WriteString("UPDATE ")
	sqlCommand.WriteString(builder.tableName.String())
	sqlCommand.WriteString(" SET ")
	counter := 0
	for column, field := range values {
		if counter > 0 {
			sqlCommand.WriteRune(',')
		}
		counter++
		sqlCommand.WriteString(column)
		sqlCommand.WriteString(" = ")
		sqlCommand.WriteRune('\'')
		sqlCommand.WriteString(field)
		sqlCommand.WriteRune('\'')
	}
	if builder.where != nil {
		sqlCommand.WriteString(" WHERE ")
		sqlCommand.WriteString(builder.where.Build())
	}
	return sqlCommand.String()
}

func (builder QueryBuilder) Insert(values ...string) string {
	sqlCommand := strings.Builder{}
	sqlCommand.WriteString("INSERT INTO ")
	sqlCommand.WriteString(builder.tableName.Name)
	sqlCommand.WriteRune('(')

	for index, fieldName := range builder.fields {
		if index > 0 {
			sqlCommand.WriteRune(',')
		}
		sqlCommand.WriteString(fieldName.Name)
	}
	sqlCommand.WriteRune(')')
	sqlCommand.WriteString(" VALUES(")
	for index, value := range values {
		if index > 0 {
			sqlCommand.WriteRune(',')
		}
		sqlCommand.WriteString(value)
	}
	sqlCommand.WriteRune(')')
	return sqlCommand.String()
}

func (builder QueryBuilder) Select() string {
	sqlCommand := strings.Builder{}
	sqlCommand.WriteString("SELECT ")
	if len(builder.fields) < 0 {
		sqlCommand.WriteRune('*')
	} else {
		for index, fieldName := range builder.fields {
			if index > 0 {
				sqlCommand.WriteRune(',')
			}
			sqlCommand.WriteString(fieldName.Build())
		}
	}
	sqlCommand.WriteString(" FROM ")
	sqlCommand.WriteString(builder.tableName.Name)
	if builder.where != nil {
		sqlCommand.WriteString(" WHERE ")
		sqlCommand.WriteString(builder.where.Build())
	}
	if len(builder.organizers) > 0 {
		for _, item := range builder.organizers {
			sqlCommand.WriteRune(' ')
			sqlCommand.WriteString(item.Build())
		}
	}

	return sqlCommand.String()
}

func (builder QueryBuilder) Delete() string {
	sqlCommand := strings.Builder{}
	sqlCommand.WriteString("DELETE FROM ")
	sqlCommand.WriteString(builder.tableName.Name)
	if builder.where != nil {
		sqlCommand.WriteString(" WHERE ")
		sqlCommand.WriteString(builder.where.Build())
	}
	return sqlCommand.String()
}
