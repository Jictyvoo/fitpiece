package builder

import "strings"

type RawSqlGenerator struct {
	Query QueryBuilder
}

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
	if generator.Query.where != nil {
		sqlCommand.WriteString(" WHERE ")
		sqlCommand.WriteString(generator.Query.where.Build())
	}
	return sqlCommand.String()
}

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

func (generator RawSqlGenerator) Select() string {
	sqlCommand := strings.Builder{}

	sqlCommand.WriteString("SELECT ")
	buildSelectColumns(&sqlCommand, generator.Query)
	sqlCommand.WriteString(" FROM ")
	sqlCommand.WriteString(generator.Query.tableName.Name)

	for _, joinClause := range generator.Query.joins {
		sqlCommand.WriteString(joinClause.Build())
		sqlCommand.WriteRune(' ')
	}
	if generator.Query.where != nil {
		sqlCommand.WriteString(" WHERE ")
		sqlCommand.WriteString(generator.Query.where.Build())
	}
	if len(generator.Query.organizers) > 0 {
		for _, item := range generator.Query.organizers {
			sqlCommand.WriteRune(' ')
			sqlCommand.WriteString(item.Build())
		}
	}

	return sqlCommand.String()
}

func (generator RawSqlGenerator) Delete() string {
	sqlCommand := strings.Builder{}
	sqlCommand.WriteString("DELETE FROM ")
	sqlCommand.WriteString(generator.Query.tableName.Name)
	if generator.Query.where != nil {
		sqlCommand.WriteString(" WHERE ")
		sqlCommand.WriteString(generator.Query.where.Build())
	}
	return sqlCommand.String()
}
