package builder

import (
	"github.com/wrapped-owls/fitpiece/cimenteiro/internal/utils"
	"strings"
)

// RawSqlGenerator structure that defines methods to create complete raw SQL queries
type RawSqlGenerator struct {
	Query QueryBuilder
}

func (generator RawSqlGenerator) buildWhereClause(writer utils.Writer) {
	if generator.Query.where != nil {
		_, _ = writer.WriteString(" WHERE ")
		generator.Query.where.Build(writer)
	}
}

// Update takes a map with values in string format and generates a raw SQL update query
func (generator RawSqlGenerator) Update(values map[string]string) string {
	sqlCommand := strings.Builder{}
	sqlCommand.WriteString("UPDATE ")
	sqlCommand.WriteString(generator.Query.tableName.String())
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
	generator.buildWhereClause(&sqlCommand)
	return sqlCommand.String()
}

// Insert takes multiple values in string format and generates a raw SQL 'insert' query.
//
// The values that define which columns will be inserted are defined by the QueryBuilder.Fields method.
func (generator RawSqlGenerator) Insert(values ...string) string {
	sqlCommand := strings.Builder{}

	sqlCommand.WriteString("INSERT INTO ")
	sqlCommand.WriteString(generator.Query.tableName.Name)
	sqlCommand.WriteRune('(')
	buildSelectColumns(&sqlCommand, generator.Query)
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

// Select generates a raw SQL 'select' query.
//
// The values that define which columns will be added to the query are defined by the QueryBuilder.Fields method.
// In case no fields are added, a select-all query will be generated
func (generator RawSqlGenerator) Select() string {
	sqlCommand := strings.Builder{}

	sqlCommand.WriteString("SELECT ")
	buildSelectColumns(&sqlCommand, generator.Query)
	sqlCommand.WriteString(" FROM ")
	sqlCommand.WriteString(generator.Query.tableName.Name)

	for _, joinClause := range generator.Query.joins {
		joinClause.Build(&sqlCommand)
		sqlCommand.WriteRune(' ')
	}

	generator.buildWhereClause(&sqlCommand)
	if len(generator.Query.organizers) > 0 {
		for _, item := range generator.Query.organizers {
			sqlCommand.WriteRune(' ')
			item.Build(&sqlCommand)
		}
	}

	return sqlCommand.String()
}

// Delete generates a raw SQL 'delete' query.
func (generator RawSqlGenerator) Delete() string {
	sqlCommand := strings.Builder{}
	sqlCommand.WriteString("DELETE FROM ")
	sqlCommand.WriteString(generator.Query.tableName.Name)

	generator.buildWhereClause(&sqlCommand)
	return sqlCommand.String()
}
