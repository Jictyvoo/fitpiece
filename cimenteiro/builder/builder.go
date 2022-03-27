package builder

import (
	"github.com/wrapped-owls/fitpiece/cimenteiro/internal/elements"
)

type QueryBuilder struct {
	tableName  elements.TableName
	fields     []elements.FieldExpression
	joins      []elements.JoinClause
	where      elements.Expression
	organizers []elements.Expression
}

func New(table elements.TableName) QueryBuilder {
	return QueryBuilder{tableName: table}
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

/**********************************************
* ClauseBuilder section of the QueryBuilder
**********************************************/

func (query QueryBuilder) Not(expression elements.Expression) elements.Expression {
	return NewPrefixExpression("NOT", expression)
}

// RawClause is a clause that is not a field expression
func (query QueryBuilder) RawClause(expression string) elements.Expression {
	return RawExpression(expression)
}

func (query QueryBuilder) In(clause elements.Expression, values ...any) elements.Clause {
	return elements.Clause{
		FirstHalf:  clause,
		Operator:   elements.OperatorIn,
		SecondHalf: ArrayExpression(values...),
	}
}

func (query QueryBuilder) InQuery(clause elements.Expression, subQuery *QueryBuilder) elements.Clause {
	return elements.Clause{
		FirstHalf:  clause,
		Operator:   elements.OperatorIn,
		SecondHalf: SubQuery[any](subQuery),
	}
}

func (query QueryBuilder) NotIn(clause elements.Expression, values ...any) elements.Clause {
	inExpression := query.In(clause, values...)
	inExpression.Operator = elements.OperatorNotIn
	return inExpression
}

func (query QueryBuilder) NotInQuery(clause elements.Expression, subQuery *QueryBuilder) elements.Clause {
	inExpression := query.InQuery(clause, subQuery)
	inExpression.Operator = elements.OperatorNotIn
	return inExpression
}

func (query QueryBuilder) Equal(column string, value any) elements.Clause {
	return elements.Clause{
		FirstHalf: elements.FieldExpression{
			Name: column,
		},
		Operator:   elements.OperatorEqual,
		SecondHalf: ValueExpression[any]{value: value},
	}
}

func (query QueryBuilder) Different(column string, value any) elements.Clause {
	clause := query.Equal(column, value)
	clause.Operator = elements.OperatorDifference
	return clause
}

func (query QueryBuilder) GreaterThan(column string, value any) elements.Clause {
	clause := query.Equal(column, value)
	clause.Operator = elements.OperatorGreaterThan
	return clause
}

func (query QueryBuilder) LessThan(column string, value any) elements.Clause {
	clause := query.Equal(column, value)
	clause.Operator = elements.OperatorLessThan
	return clause
}

func (query QueryBuilder) GreaterEqual(column string, value any) elements.Clause {
	clause := query.Equal(column, value)
	clause.Operator = elements.OperatorGreaterEqual
	return clause
}

func (query QueryBuilder) LessEqual(column string, value any) elements.Clause {
	clause := query.Equal(column, value)
	clause.Operator = elements.OperatorLessEqual
	return clause
}

func (query QueryBuilder) And(first, second elements.Expression) elements.Clause {
	return elements.Clause{
		FirstHalf:  first,
		Operator:   "AND",
		SecondHalf: second,
	}
}

func (query QueryBuilder) Or(first, second elements.Expression) elements.Clause {
	return elements.Clause{
		FirstHalf:  first,
		Operator:   "OR",
		SecondHalf: second,
	}
}
