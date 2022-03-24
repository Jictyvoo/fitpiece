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

func (query *QueryBuilder) Fields(fields ...string) *QueryBuilder {
	index := len(query.fields)
	// Extends the slice cap, to prevent reallocate inside the loop
	newCap := len(fields) + index
	query.fields = query.fields[:newCap]
	for _, field := range fields {
		query.fields[index] = elements.FieldExpression{Name: field, Alias: ""}
		index++
	}
	return query
}

// FIXME: Add parameter for first tableName (to not use builder.tableName)
func (query *QueryBuilder) buildJoin(
	tableName elements.TableName, firstColumn string, secondColumn string,
) elements.JoinClause {
	return elements.JoinClause{
		JoinType: elements.JoinAll,
		Table:    tableName,
		On: elements.Clause{
			FirstHalf:  elements.FieldExpression{Name: query.tableName.Column(firstColumn)},
			Operator:   elements.OperatorEqual,
			SecondHalf: elements.FieldExpression{Name: tableName.Column(secondColumn)},
		},
	}
}

func (query *QueryBuilder) LeftJoin(tableName elements.TableName, firstColumn string, secondColumn string) *QueryBuilder {
	leftJoin := query.buildJoin(tableName, firstColumn, secondColumn)
	leftJoin.JoinType = elements.JoinLeft
	query.joins = append(query.joins, leftJoin)
	return query
}

func (query *QueryBuilder) RightJoin(tableName elements.TableName, firstColumn string, secondColumn string) *QueryBuilder {
	rightJoin := query.buildJoin(tableName, firstColumn, secondColumn)
	rightJoin.JoinType = elements.JoinRight
	query.joins = append(query.joins, rightJoin)
	return query
}

func (query *QueryBuilder) InnerJoin(tableName elements.TableName, firstColumn string, secondColumn string) *QueryBuilder {
	innerJoin := query.buildJoin(tableName, firstColumn, secondColumn)
	innerJoin.JoinType = elements.JoinInner
	query.joins = append(query.joins, innerJoin)
	return query
}

func (query *QueryBuilder) OuterJoin(tableName elements.TableName, firstColumn string, secondColumn string) *QueryBuilder {
	outerJoin := query.buildJoin(tableName, firstColumn, secondColumn)
	outerJoin.JoinType = elements.JoinOuter
	query.joins = append(query.joins, outerJoin)
	return query
}

func (query QueryBuilder) Update(values map[string]string) string {
	sqlCommand := strings.Builder{}
	sqlCommand.WriteString("UPDATE ")
	sqlCommand.WriteString(query.tableName.String())
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
	if query.where != nil {
		sqlCommand.WriteString(" WHERE ")
		sqlCommand.WriteString(query.where.Build())
	}
	return sqlCommand.String()
}

func (query QueryBuilder) Insert(values ...string) string {
	sqlCommand := strings.Builder{}
	sqlCommand.WriteString("INSERT INTO ")
	sqlCommand.WriteString(query.tableName.Name)
	sqlCommand.WriteRune('(')

	for index, fieldName := range query.fields {
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

func (query QueryBuilder) Select() string {
	sqlCommand := strings.Builder{}
	sqlCommand.WriteString("SELECT ")
	if len(query.fields) < 0 {
		sqlCommand.WriteRune('*')
	} else {
		for index, fieldName := range query.fields {
			if index > 0 {
				sqlCommand.WriteRune(',')
			}
			sqlCommand.WriteString(fieldName.Build())
		}
	}
	sqlCommand.WriteString(" FROM ")
	sqlCommand.WriteString(query.tableName.Name)
	if query.where != nil {
		sqlCommand.WriteString(" WHERE ")
		sqlCommand.WriteString(query.where.Build())
	}
	if len(query.organizers) > 0 {
		for _, item := range query.organizers {
			sqlCommand.WriteRune(' ')
			sqlCommand.WriteString(item.Build())
		}
	}

	return sqlCommand.String()
}

func (query QueryBuilder) Delete() string {
	sqlCommand := strings.Builder{}
	sqlCommand.WriteString("DELETE FROM ")
	sqlCommand.WriteString(query.tableName.Name)
	if query.where != nil {
		sqlCommand.WriteString(" WHERE ")
		sqlCommand.WriteString(query.where.Build())
	}
	return sqlCommand.String()
}
