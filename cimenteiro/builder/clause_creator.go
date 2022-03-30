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

// In takes an elements.Expression for the left side of the in (it represents the value to be compared),
// and a list of values that will be added on the right side of the expression
func (query __clauseCreator) In(clause elements.Expression, values ...any) elements.Clause {
	return elements.Clause{
		FirstHalf:  clause,
		Operator:   elements.OperatorIn,
		SecondHalf: expressions.ArrayExpression(values...),
	}
}

// InQuery takes an elements.Expression for the left side of the in (it represents the value to be compared),
// and a QueryBuilder that will generate an 'select' SQL-string and be added on the right side of the expression
func (query __clauseCreator) InQuery(clause elements.Expression, subQuery *QueryBuilder) elements.Clause {
	return elements.Clause{
		FirstHalf:  clause,
		Operator:   elements.OperatorIn,
		SecondHalf: SubQuery(subQuery),
	}
}

// NotIn takes an elements.Expression for the left side of the in (it represents the value to be compared),
// and a list of values that will be added on the right side of the expression
func (query __clauseCreator) NotIn(clause elements.Expression, values ...any) elements.Clause {
	inExpression := query.In(clause, values...)
	inExpression.Operator = elements.OperatorNotIn
	return inExpression
}

// NotInQuery takes an elements.Expression for the left side of the in (it represents the value to be compared),
// and a QueryBuilder that will generate an 'select' SQL-string and be added on the right side of the expression
func (query __clauseCreator) NotInQuery(clause elements.Expression, subQuery *QueryBuilder) elements.Clause {
	inExpression := query.InQuery(clause, subQuery)
	inExpression.Operator = elements.OperatorNotIn
	return inExpression
}

// Equal generates a clause that checks if a column is equal to a given value
//
// The generated expression will be like: `columnName = "value"` | `columnName = 42`
func (query __clauseCreator) Equal(column string, value any) elements.Clause {
	return elements.Clause{
		FirstHalf: elements.FieldExpression{
			Name: column,
		},
		Operator:   elements.OperatorEqual,
		SecondHalf: expressions.NewValueExpression(value),
	}
}

// Different generates a clause that checks if a column is different to a given value
//
// The generated expression will be like: `columnName <> "value"` | `columnName <> 42`
func (query __clauseCreator) Different(column string, value any) elements.Clause {
	clause := query.Equal(column, value)
	clause.Operator = elements.OperatorDifference
	return clause
}

// GreaterThan generates a clause that checks if a column is greater than a given value
//
// The generated expression will be like: `columnName > "value"` | `columnName > 42`
func (query __clauseCreator) GreaterThan(column string, value any) elements.Clause {
	clause := query.Equal(column, value)
	clause.Operator = elements.OperatorGreaterThan
	return clause
}

// LessThan generates a clause that checks if a column is less than a given value
//
// The generated expression will be like: `columnName < "value"` | `columnName < 42`
func (query __clauseCreator) LessThan(column string, value any) elements.Clause {
	clause := query.Equal(column, value)
	clause.Operator = elements.OperatorLessThan
	return clause
}

// GreaterEqual generates a clause that checks if a column is grater-equal than a given value
//
// The generated expression will be like: `columnName >= "value"` | `columnName >= 42`
func (query __clauseCreator) GreaterEqual(column string, value any) elements.Clause {
	clause := query.Equal(column, value)
	clause.Operator = elements.OperatorGreaterEqual
	return clause
}

// LessEqual generates a clause that checks if a column is less-equal than a given value
//
// The generated expression will be like: `columnName <= "value"` | `columnName <= 42`
func (query __clauseCreator) LessEqual(column string, value any) elements.Clause {
	clause := query.Equal(column, value)
	clause.Operator = elements.OperatorLessEqual
	return clause
}

// And receives two expressions and returns an elements.Clause that have an 'AND' operator between the expressions
func (query __clauseCreator) And(first, second elements.Expression) elements.Clause {
	return elements.Clause{
		FirstHalf:  first,
		Operator:   elements.OperatorAnd,
		SecondHalf: second,
	}
}

// Or receives two expressions and returns an elements.Clause that have an 'OR' operator between the expressions
func (query __clauseCreator) Or(first, second elements.Expression) elements.Clause {
	return elements.Clause{
		FirstHalf:  first,
		Operator:   elements.OperatorOr,
		SecondHalf: second,
	}
}
