package builder

import (
	"github.com/wrapped-owls/fitpiece/cimenteiro/builder/expressions"
	"github.com/wrapped-owls/fitpiece/cimenteiro/internal/elements"
)

type QueryBuilder struct {
	tableName  elements.TableName
	fields     []elements.FieldExpression
	joins      []elements.JoinClause
	where      elements.Expression
	organizers map[elements.Organizer]elements.Expression
}

func New(table elements.TableName) QueryBuilder {
	return QueryBuilder{tableName: table, organizers: map[elements.Organizer]elements.Expression{}}
}

func expandSlice[T any](originalSlice []T, newCap uint) []T {
	expandedSlice := make([]T, len(originalSlice), newCap)
	copy(expandedSlice, originalSlice)
	return expandedSlice
}

func (query *QueryBuilder) Fields(fields ...string) *QueryBuilder {
	size := len(query.fields)
	// Extends the slice cap, to prevent reallocate inside the loop
	newCap := len(fields) + size
	query.fields = expandSlice(query.fields, uint(newCap))
	// query.fields = query.fields[:newCap]
	for _, field := range fields {
		query.fields = append(query.fields, elements.FieldExpression{Name: field, Alias: ""})
	}
	return query
}

func (query *QueryBuilder) FieldsAs(fields ...string) *QueryBuilder {
	size := len(query.fields)
	newLength := len(fields)
	// Divide by 2, moving bit to the right
	newLength = newLength >> 1
	// Extends the slice cap, to prevent reallocate inside the loop
	newCap := newLength + size
	query.fields = expandSlice(query.fields, uint(newCap))

	var previousField string
	var index uint = 1
	for _, field := range fields {
		if index == 2 {
			query.fields = append(query.fields, elements.FieldExpression{Name: previousField, Alias: field})
			index = 0
		}
		index++
		previousField = field
	}
	return query
}

func (query *QueryBuilder) buildJoin(
	originTable, tableName elements.TableName, firstColumn, secondColumn string,
) elements.JoinClause {
	return elements.JoinClause{
		JoinType: elements.JoinAll,
		Table:    tableName,
		On: elements.Clause{
			FirstHalf:  elements.FieldExpression{Name: originTable.Column(firstColumn)},
			Operator:   elements.OperatorEqual,
			SecondHalf: elements.FieldExpression{Name: tableName.Column(secondColumn)},
		},
	}
}

func (query *QueryBuilder) CrossJoin(from, with elements.TableName, firstColumn, secondColumn string) *QueryBuilder {
	leftJoin := query.buildJoin(from, with, firstColumn, secondColumn)
	query.joins = append(query.joins, leftJoin)
	return query
}

func (query *QueryBuilder) CrossJoinOrigin(with elements.TableName, firstColumn, secondColumn string) *QueryBuilder {
	return query.CrossJoin(query.tableName, with, firstColumn, secondColumn)
}

func (query *QueryBuilder) LeftJoin(from, with elements.TableName, firstColumn, secondColumn string) *QueryBuilder {
	leftJoin := query.buildJoin(from, with, firstColumn, secondColumn)
	leftJoin.JoinType = elements.JoinLeft
	query.joins = append(query.joins, leftJoin)
	return query
}

func (query *QueryBuilder) LeftJoinOrigin(with elements.TableName, firstColumn, secondColumn string) *QueryBuilder {
	return query.LeftJoin(query.tableName, with, firstColumn, secondColumn)
}

func (query *QueryBuilder) RightJoin(from, with elements.TableName, firstColumn, secondColumn string) *QueryBuilder {
	rightJoin := query.buildJoin(from, with, firstColumn, secondColumn)
	rightJoin.JoinType = elements.JoinRight
	query.joins = append(query.joins, rightJoin)
	return query
}

func (query *QueryBuilder) RightJoinOrigin(with elements.TableName, firstColumn, secondColumn string) *QueryBuilder {
	return query.RightJoin(query.tableName, with, firstColumn, secondColumn)
}

func (query *QueryBuilder) InnerJoin(from, with elements.TableName, firstColumn, secondColumn string) *QueryBuilder {
	innerJoin := query.buildJoin(from, with, firstColumn, secondColumn)
	innerJoin.JoinType = elements.JoinInner
	query.joins = append(query.joins, innerJoin)
	return query
}

func (query *QueryBuilder) InnerJoinOrigin(with elements.TableName, firstColumn, secondColumn string) *QueryBuilder {
	return query.InnerJoin(query.tableName, with, firstColumn, secondColumn)
}

func (query *QueryBuilder) OuterJoin(from, with elements.TableName, firstColumn, secondColumn string) *QueryBuilder {
	outerJoin := query.buildJoin(from, with, firstColumn, secondColumn)
	outerJoin.JoinType = elements.JoinOuter
	query.joins = append(query.joins, outerJoin)
	return query
}

func (query *QueryBuilder) OuterJoinOrigin(with elements.TableName, firstColumn, secondColumn string) *QueryBuilder {
	return query.OuterJoin(query.tableName, with, firstColumn, secondColumn)
}

func (query *QueryBuilder) Where(whereClause elements.Expression) *QueryBuilder {
	query.where = whereClause
	return query
}

// GroupBy Creates the group expression and add it to the query
func (query *QueryBuilder) GroupBy(fields ...string) *QueryBuilder {
	query.organizers[elements.OrganizerGroup] = expressions.PrefixExpression(
		"GROUP BY", expressions.MultiDirectValueExpression([2]rune{0, 0}, fields...),
	)
	return query
}

// Having specify HAVING conditions for GROUP BY
func (query *QueryBuilder) Having(expression elements.Expression) *QueryBuilder {
	query.organizers[elements.OrganizerHaving] = expressions.PrefixExpression(
		"HAVING", expression,
	)
	return query
}

func (query QueryBuilder) buildOrderBy(desc bool, columns ...string) elements.Expression {
	sortType := "ASC"
	if desc {
		sortType = "DESC"
	}
	return expressions.SuffixExpression(
		expressions.PrefixExpression(
			"ORDER BY", expressions.MultiDirectValueExpression([2]rune{0, 0}, columns...),
		),
		sortType,
	)
}

// OrderBy specify order when retrieve records from database
func (query *QueryBuilder) OrderBy(columns ...string) *QueryBuilder {
	query.organizers[elements.OrganizerOrder] = query.buildOrderBy(false, columns...)
	return query
}

// OrderByDesc specify order when retrieve records from database
func (query *QueryBuilder) OrderByDesc(columns ...string) *QueryBuilder {
	query.organizers[elements.OrganizerGroup] = query.buildOrderBy(true, columns...)
	return query
}

// Limit specify the number of records to be retrieved
// TODO: Create dialects to know how to create the statements
func (query *QueryBuilder) Limit(limit int) *QueryBuilder {
	query.organizers[elements.OrganizerLimit] = expressions.PrefixExpression(
		"LIMIT",
		expressions.NewValueExpression(limit),
	)
	return query
}

// Offset specify the number of records to skip before starting to return the records
func (query *QueryBuilder) Offset(offset int) *QueryBuilder {
	query.organizers[elements.OrganizerOffset] = expressions.PrefixExpression(
		"OFFSET",
		expressions.NewValueExpression(offset),
	)
	return query
}
