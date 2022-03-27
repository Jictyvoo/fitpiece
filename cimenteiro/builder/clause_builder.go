package builder

import (
	"github.com/wrapped-owls/fitpiece/cimenteiro/builder/expressions"
	"github.com/wrapped-owls/fitpiece/cimenteiro/internal/elements"
)

/**********************************************
* ClauseBuilder section of the QueryBuilder
**********************************************/

func (query QueryBuilder) Not(expression elements.Expression) elements.Expression {
	return expressions.NewPrefixExpression("NOT", expression)
}

// RawClause is a clause that is not a field expression
func (query QueryBuilder) RawClause(expression string) elements.Expression {
	return expressions.RawExpression(expression)
}

func (query QueryBuilder) In(clause elements.Expression, values ...any) elements.Clause {
	return elements.Clause{
		FirstHalf:  clause,
		Operator:   elements.OperatorIn,
		SecondHalf: expressions.ArrayExpression(values...),
	}
}

func (query QueryBuilder) InQuery(clause elements.Expression, subQuery *QueryBuilder) elements.Clause {
	return elements.Clause{
		FirstHalf:  clause,
		Operator:   elements.OperatorIn,
		SecondHalf: SubQuery(subQuery),
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
		SecondHalf: expressions.NewValueExpression(value),
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
