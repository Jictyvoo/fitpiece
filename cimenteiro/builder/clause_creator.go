package builder

import (
	"github.com/jictyvoo/fitpiece/cimenteiro/builder/expressions"
	"github.com/jictyvoo/fitpiece/cimenteiro/internal/elements"
)

// __clauseCreator syntax sugar for create a const namespace that
// is useful to attach methods and functions
type __clauseCreator bool

func (query __clauseCreator) Not(expression elements.Expression) elements.Expression {
	return expressions.PrefixExpression("NOT", expression)
}

// RawClause is a clause that is not a field expression
func (query __clauseCreator) RawClause(expression string) elements.Expression {
	return expressions.RawExpression(expression)
}

func (query __clauseCreator) In(clause elements.Expression, values ...any) elements.Clause {
	return elements.Clause{
		FirstHalf:  clause,
		Operator:   elements.OperatorIn,
		SecondHalf: expressions.ArrayExpression(values...),
	}
}

func (query __clauseCreator) InQuery(clause elements.Expression, subQuery *QueryBuilder) elements.Clause {
	return elements.Clause{
		FirstHalf:  clause,
		Operator:   elements.OperatorIn,
		SecondHalf: SubQuery(subQuery),
	}
}

func (query __clauseCreator) NotIn(clause elements.Expression, values ...any) elements.Clause {
	inExpression := query.In(clause, values...)
	inExpression.Operator = elements.OperatorNotIn
	return inExpression
}

func (query __clauseCreator) NotInQuery(clause elements.Expression, subQuery *QueryBuilder) elements.Clause {
	inExpression := query.InQuery(clause, subQuery)
	inExpression.Operator = elements.OperatorNotIn
	return inExpression
}

func (query __clauseCreator) Equal(column string, value any) elements.Clause {
	return elements.Clause{
		FirstHalf: elements.FieldExpression{
			Name: column,
		},
		Operator:   elements.OperatorEqual,
		SecondHalf: expressions.NewValueExpression(value),
	}
}

func (query __clauseCreator) Different(column string, value any) elements.Clause {
	clause := query.Equal(column, value)
	clause.Operator = elements.OperatorDifference
	return clause
}

func (query __clauseCreator) GreaterThan(column string, value any) elements.Clause {
	clause := query.Equal(column, value)
	clause.Operator = elements.OperatorGreaterThan
	return clause
}

func (query __clauseCreator) LessThan(column string, value any) elements.Clause {
	clause := query.Equal(column, value)
	clause.Operator = elements.OperatorLessThan
	return clause
}

func (query __clauseCreator) GreaterEqual(column string, value any) elements.Clause {
	clause := query.Equal(column, value)
	clause.Operator = elements.OperatorGreaterEqual
	return clause
}

func (query __clauseCreator) LessEqual(column string, value any) elements.Clause {
	clause := query.Equal(column, value)
	clause.Operator = elements.OperatorLessEqual
	return clause
}

func (query __clauseCreator) And(first, second elements.Expression) elements.Clause {
	return elements.Clause{
		FirstHalf:  first,
		Operator:   "AND",
		SecondHalf: second,
	}
}

func (query __clauseCreator) Or(first, second elements.Expression) elements.Clause {
	return elements.Clause{
		FirstHalf:  first,
		Operator:   "OR",
		SecondHalf: second,
	}
}
