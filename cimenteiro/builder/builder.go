package builder

import (
	"github.com/Wrapped-Owls/fitpiece/cimenteiro/internal/elements"
)

type QueryBuilder struct {
	tableName  elements.TableName
	fields     []elements.FieldExpression
	joins      []elements.JoinClause
	where      elements.Expression
	organizers []elements.Expression
}

func (query *QueryBuilder) Fields(fields ...string) *QueryBuilder {
	size := len(query.fields)
	// Extends the slice cap, to prevent reallocate inside the loop
	newCap := len(fields) + size
	query.fields = query.fields[:newCap]
	for index, field := range fields {
		query.fields[size+index] = elements.FieldExpression{Name: field, Alias: ""}
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
	query.fields = query.fields[:newCap]

	var previousField string
	var index uint = 1
	for _, field := range fields {
		if index == 2 {
			query.fields[size] = elements.FieldExpression{Name: previousField, Alias: field}
			size++
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
