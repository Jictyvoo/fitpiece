package builder

import (
	"github.com/jictyvoo/fitpiece/cimenteiro/internal/elements"
	"github.com/jictyvoo/fitpiece/cimenteiro/internal/utils"
	"sort"
	"strings"
)

// PlaceholderSqlGenerator structure that defines methods to create complete SQL queries
// with a placeholder and a slice with its values
type PlaceholderSqlGenerator struct {
	Query       *QueryBuilder
	Placeholder string
}

func (generator PlaceholderSqlGenerator) buildWhere(sqlCommand utils.Writer, valuesList []any) []any {
	if generator.Query.where != nil {
		_, _ = sqlCommand.WriteString(" WHERE ")

		whereWriter := strings.Builder{}
		_, args := generator.Query.where.BuildPlaceholder(&whereWriter, generator.Placeholder)
		valuesList = append(valuesList, args...)

		strResult := utils.RemoveBrackets(whereWriter.String())
		_, _ = sqlCommand.WriteString(strResult)
	}
	return valuesList
}

func (generator PlaceholderSqlGenerator) buildOrganizers(writer utils.Writer, valuesList []any) []any {
	for _, key := range elements.OrganizersSortOrder() {
		if item, ok := generator.Query.organizers[key]; ok {
			_, _ = writer.WriteRune(' ')
			_, args := item.BuildPlaceholder(writer, generator.Placeholder)
			valuesList = append(valuesList, args...)
		}
	}
	return valuesList
}

// Update takes a map with values in string format and generates a SQL update query
// with a slice of its arguments
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

// Insert takes multiple values in string format and generates a SQL 'insert' query.
//
// The values that define which columns will be inserted are defined by the QueryBuilder.Fields method.
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

// Select generates a SQL 'select' query.
//
// The values that define which columns will be added to the query are defined by the QueryBuilder.Fields method.
// In case no fields are added, a select-all query will be generated
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

// Delete generates a SQL 'delete' query.
func (generator PlaceholderSqlGenerator) Delete() (string, []any) {
	sqlCommand := strings.Builder{}
	var valuesList []any

	sqlCommand.WriteString("DELETE FROM ")
	sqlCommand.WriteString(generator.Query.tableName.Name)
	valuesList = generator.buildWhere(&sqlCommand, valuesList)
	return sqlCommand.String(), valuesList
}
