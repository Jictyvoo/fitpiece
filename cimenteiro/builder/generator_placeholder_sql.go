package builder

import "strings"

type PlaceholderSqlGenerator struct {
	Query       QueryBuilder
	Placeholder string
}

func (generator PlaceholderSqlGenerator) Update(values map[string]any) (string, []any) {
	sqlCommand := strings.Builder{}
	valuesList := make([]any, 0, len(values))

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

		sqlCommand.WriteString(generator.Placeholder)
		valuesList = append(valuesList, field)
	}
	if generator.Query.where != nil {
		sqlCommand.WriteString(" WHERE ")
		whereGen, args := generator.Query.where.BuildPlaceholder(generator.Placeholder)
		valuesList = append(valuesList, args...)
		sqlCommand.WriteString(whereGen)
	}
	return sqlCommand.String(), valuesList
}

func (generator PlaceholderSqlGenerator) Insert(values ...any) (string, []any) {
	sqlCommand := strings.Builder{}
	valuesList := make([]any, 0, len(values))

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
		sqlCommand.WriteString(generator.Placeholder)
		valuesList = append(valuesList, value)
	}
	sqlCommand.WriteRune(')')
	return sqlCommand.String(), valuesList
}

func (generator PlaceholderSqlGenerator) Select() (string, []any) {
	sqlCommand := strings.Builder{}
	valuesList := make([]any, 0, len(generator.Query.fields))

	sqlCommand.WriteString("SELECT ")
	buildSelectColumns(&sqlCommand, generator.Query)
	sqlCommand.WriteString(" FROM ")
	sqlCommand.WriteString(generator.Query.tableName.Name)

	buildJoinClauses(&sqlCommand, generator.Query)
	if generator.Query.where != nil {
		sqlCommand.WriteString(" WHERE ")
		whereGen, args := generator.Query.where.BuildPlaceholder(generator.Placeholder)
		valuesList = append(valuesList, args...)
		sqlCommand.WriteString(whereGen)
	}
	if len(generator.Query.organizers) > 0 {
		for _, item := range generator.Query.organizers {
			sqlCommand.WriteRune(' ')
			sqlCommand.WriteString(item.Build())
		}
	}

	return sqlCommand.String(), valuesList
}

func (generator PlaceholderSqlGenerator) Delete() (string, []any) {
	sqlCommand := strings.Builder{}
	var valuesList []any

	sqlCommand.WriteString("DELETE FROM ")
	sqlCommand.WriteString(generator.Query.tableName.Name)
	if generator.Query.where != nil {
		sqlCommand.WriteString(" WHERE ")
		strResult, args := generator.Query.where.BuildPlaceholder(generator.Placeholder)
		valuesList = args
		sqlCommand.WriteString(strResult)
	}
	return sqlCommand.String(), valuesList
}
