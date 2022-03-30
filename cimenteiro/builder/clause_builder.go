package builder

import (
	"github.com/jictyvoo/fitpiece/cimenteiro/builder/expressions"
	"github.com/jictyvoo/fitpiece/cimenteiro/internal/elements"
)

type expressionProcessor = func(expression elements.Expression) elements.Expression

// ClauseBuilder is the object that helps to create clauses, that can be used in the
// "where" method of the QueryBuilder
type ClauseBuilder struct {
	clause      elements.Expression
	processNext expressionProcessor
}

// CreateClause creates a new empty builder and returns it as a pointer, to help in chaining
func CreateClause() *ClauseBuilder {
	return &ClauseBuilder{}
}

// Clause generates and return the final clause expression for all the clause builder
func (builder ClauseBuilder) Clause() elements.Expression {
	return builder.clause
}

func (builder *ClauseBuilder) registerClause(expression elements.Expression) {
	if builder.processNext != nil {
		builder.clause = builder.processNext(expression)
		builder.processNext = nil
	} else {
		builder.clause = expression
	}
}

// Not register a callback to add the keyword "NOT" as a prefix to the next expression that will be added.
// If in case it was the last method called, nothing will be done
func (builder *ClauseBuilder) Not() *ClauseBuilder {
	currentExpression := builder.clause
	currentProcessor := builder.processNext
	builder.processNext = func(expression elements.Expression) elements.Expression {
		notExpression := ClauseCreator.Not(expression)
		if currentProcessor != nil {
			return currentProcessor(notExpression)
		}
		return expressions.NewPairsClauseExpression(currentExpression, notExpression)
	}
	return builder
}

// In takes an elements.Expression for the left side of the in (it represents the value to be compared),
// and a list of values that will be added on the right side of the expression
func (builder *ClauseBuilder) In(clause elements.Expression, values ...any) *ClauseBuilder {
	builder.registerClause(ClauseCreator.In(clause, values...))
	return builder
}

// InQuery takes an elements.Expression for the left side of the in (it represents the value to be compared),
// and a QueryBuilder that will generate an 'select' SQL-string and be added on the right side of the expression
func (builder *ClauseBuilder) InQuery(clause elements.Expression, subQuery *QueryBuilder) *ClauseBuilder {
	builder.registerClause(ClauseCreator.InQuery(clause, subQuery))
	return builder
}

// NotIn takes an elements.Expression for the left side of the in (it represents the value to be compared),
// and a list of values that will be added on the right side of the expression
func (builder *ClauseBuilder) NotIn(clause elements.Expression, values ...any) *ClauseBuilder {
	builder.registerClause(ClauseCreator.NotIn(clause, values...))
	return builder
}

// NotInQuery takes an elements.Expression for the left side of the in (it represents the value to be compared),
// and a QueryBuilder that will generate an 'select' SQL-string and be added on the right side of the expression
func (builder *ClauseBuilder) NotInQuery(clause elements.Expression, subQuery *QueryBuilder) *ClauseBuilder {
	builder.registerClause(ClauseCreator.NotInQuery(clause, subQuery))
	return builder
}

// Equal generates a clause that checks if a column is equal to a given value
//
// The generated expression will be like: `columnName = "value"` | `columnName = 42`
func (builder *ClauseBuilder) Equal(column string, value any) *ClauseBuilder {
	builder.registerClause(ClauseCreator.Equal(column, value))
	return builder
}

// Different generates a clause that checks if a column is different to a given value
//
// The generated expression will be like: `columnName <> "value"` | `columnName <> 42`
func (builder *ClauseBuilder) Different(column string, value any) *ClauseBuilder {
	builder.registerClause(ClauseCreator.Different(column, value))
	return builder
}

// GreaterThan generates a clause that checks if a column is greater than a given value
//
// The generated expression will be like: `columnName > "value"` | `columnName > 42`
func (builder *ClauseBuilder) GreaterThan(column string, value any) *ClauseBuilder {
	builder.registerClause(ClauseCreator.GreaterThan(column, value))
	return builder
}

// LessThan generates a clause that checks if a column is less than a given value
//
// The generated expression will be like: `columnName < "value"` | `columnName < 42`
func (builder *ClauseBuilder) LessThan(column string, value any) *ClauseBuilder {
	builder.registerClause(ClauseCreator.LessThan(column, value))
	return builder
}

// GreaterEqual generates a clause that checks if a column is grater-equal than a given value
//
// The generated expression will be like: `columnName >= "value"` | `columnName >= 42`
func (builder *ClauseBuilder) GreaterEqual(column string, value any) *ClauseBuilder {
	builder.registerClause(ClauseCreator.GreaterEqual(column, value))
	return builder
}

// LessEqual generates a clause that checks if a column is less-equal than a given value
//
// The generated expression will be like: `columnName <= "value"` | `columnName <= 42`
func (builder *ClauseBuilder) LessEqual(column string, value any) *ClauseBuilder {
	builder.registerClause(ClauseCreator.LessEqual(column, value))
	return builder
}

// And register a callback that will be called when another method is called after this.
// This callback will join the current elements.Expression built with the new elements.Expression
// using the elements.KeywordAnd to join both expressions
func (builder *ClauseBuilder) And() *ClauseBuilder {
	currentExpression := builder.clause
	builder.processNext = func(expression elements.Expression) elements.Expression {
		return ClauseCreator.And(currentExpression, expression)
	}

	return builder
}

// Or register a callback that will be called when another method is called after this.
// This callback will join the current elements.Expression built with the new elements.Expression
// using the elements.KeywordOr to join both expressions
func (builder *ClauseBuilder) Or() *ClauseBuilder {
	currentExpression := builder.clause
	builder.processNext = func(expression elements.Expression) elements.Expression {
		return ClauseCreator.Or(currentExpression, expression)
	}

	return builder
}
