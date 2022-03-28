package builder

import (
	"github.com/wrapped-owls/fitpiece/cimenteiro/internal/elements"
	"github.com/wrapped-owls/fitpiece/cimenteiro/internal/utils"
	"sort"
	"strings"
)

type PlaceholderSqlGenerator struct {
	Query       *QueryBuilder
	Placeholder string
}

func (generator PlaceholderSqlGenerator) buildWhere(sqlCommand utils.Writer, valuesList []any) []any {
	if generator.Query.where != nil {
		_, _ = sqlCommand.WriteString(" WHERE ")
		strResult, args := generator.Query.where.BuildPlaceholder(generator.Placeholder)
		valuesList = append(valuesList, args...)

		strResult = utils.RemoveBrackets(strResult)
		_, _ = sqlCommand.WriteString(strResult)
	}
	return valuesList
}

func (generator PlaceholderSqlGenerator) buildOrganizers(writer utils.Writer, valuesList []any) []any {
	for _, key := range elements.OrganizersSortOrder() {
		if item, ok := generator.Query.organizers[key]; ok {
			_, _ = writer.WriteRune(' ')
			strSql, args := item.BuildPlaceholder(generator.Placeholder)
			_, _ = writer.WriteString(strSql)
			valuesList = append(valuesList, args...)
		}
	}
	return valuesList
}

func (generator PlaceholderSqlGenerator) Update(values map[string]any) (string, []any) {
	sqlCommand := strings.Builder{}
	valuesList := make([]any, 0, len(values))

	sqlCommand.WriteString("UPDATE ")
	sqlCommand.WriteString(generator.Query.tableName.String())
	sqlCommand.WriteString(" SET ")
	counter := 0

	// take the keys and order them
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, column := range keys {
		if counter > 0 {
			sqlCommand.WriteRune(',')
			sqlCommand.WriteRune(' ')
		}
		counter++
		sqlCommand.WriteString(column)
		sqlCommand.WriteString(" = ")

		sqlCommand.WriteString(generator.Placeholder)
		field, _ := values[column]
		valuesList = append(valuesList, field)
	}
	valuesList = generator.buildWhere(&sqlCommand, valuesList)
	return sqlCommand.String(), valuesList
}

func (generator PlaceholderSqlGenerator) Insert(values ...any) (string, []any) {
	sqlCommand := strings.Builder{}
	valuesList := make([]any, 0, len(values))

	sqlCommand.WriteString("INSERT INTO ")
	sqlCommand.WriteString(generator.Query.tableName.Name)
	sqlCommand.WriteRune(' ')
	sqlCommand.WriteRune('(')
	buildSelectColumns(&sqlCommand, *generator.Query)
	sqlCommand.WriteRune(')')

	sqlCommand.WriteString(" VALUES (")
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
	buildSelectColumns(&sqlCommand, *generator.Query)
	sqlCommand.WriteString(" FROM ")
	sqlCommand.WriteString(generator.Query.tableName.Name)

	buildJoinClauses(&sqlCommand, *generator.Query)
	valuesList = generator.buildWhere(&sqlCommand, valuesList)
	if len(generator.Query.organizers) > 0 {
		valuesList = generator.buildOrganizers(&sqlCommand, valuesList)
	}
	return sqlCommand.String(), valuesList
}

func (generator PlaceholderSqlGenerator) Delete() (string, []any) {
	sqlCommand := strings.Builder{}
	var valuesList []any

	sqlCommand.WriteString("DELETE FROM ")
	sqlCommand.WriteString(generator.Query.tableName.Name)
	valuesList = generator.buildWhere(&sqlCommand, valuesList)
	return sqlCommand.String(), valuesList
}
