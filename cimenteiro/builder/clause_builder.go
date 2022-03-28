package builder

import (
	"github.com/wrapped-owls/fitpiece/cimenteiro/internal/elements"
)

type expressionProcessor = func(expression elements.Expression) elements.Expression

type ClauseBuilder struct {
	clause      elements.Expression
	processNext expressionProcessor
}

func (builder *ClauseBuilder) registerClause(expression elements.Expression) {
	if builder.processNext != nil {
		builder.clause = builder.processNext(expression)
	} else {
		builder.clause = expression
	}
}

/**********************************************
* ClauseBuilder section of the QueryBuilder
**********************************************/

func (builder *ClauseBuilder) Not() *ClauseBuilder {
	builder.processNext = ClauseCreator.Not
	return builder
}

func (builder *ClauseBuilder) In(values ...any) *ClauseBuilder {
	builder.registerClause(ClauseCreator.In(builder.clause, values...))
	return builder
}

func (builder *ClauseBuilder) InQuery(subQuery *QueryBuilder) *ClauseBuilder {
	builder.registerClause(ClauseCreator.InQuery(builder.clause, subQuery))
	return builder
}

func (builder *ClauseBuilder) NotIn(values ...any) *ClauseBuilder {
	builder.registerClause(ClauseCreator.NotIn(builder.clause, values...))
	return builder
}

func (builder *ClauseBuilder) NotInQuery(subQuery *QueryBuilder) *ClauseBuilder {
	builder.registerClause(ClauseCreator.NotInQuery(builder.clause, subQuery))
	return builder
}

func (builder *ClauseBuilder) Equal(column string, value any) *ClauseBuilder {
	builder.registerClause(ClauseCreator.Equal(column, value))
	return builder
}

func (builder *ClauseBuilder) Different(column string, value any) *ClauseBuilder {
	builder.registerClause(ClauseCreator.Different(column, value))
	return builder
}

func (builder *ClauseBuilder) GreaterThan(column string, value any) *ClauseBuilder {
	builder.registerClause(ClauseCreator.GreaterThan(column, value))
	return builder
}

func (builder *ClauseBuilder) LessThan(column string, value any) *ClauseBuilder {
	builder.registerClause(ClauseCreator.LessThan(column, value))
	return builder
}

func (builder *ClauseBuilder) GreaterEqual(column string, value any) *ClauseBuilder {
	builder.registerClause(ClauseCreator.GreaterEqual(column, value))
	return builder
}

func (builder *ClauseBuilder) LessEqual(column string, value any) *ClauseBuilder {
	builder.registerClause(ClauseCreator.LessEqual(column, value))
	return builder
}

func (builder *ClauseBuilder) And() *ClauseBuilder {
	currentExpression := builder.clause
	builder.processNext = func(expression elements.Expression) elements.Expression {
		return ClauseCreator.And(currentExpression, expression)
	}

	return builder
}

func (builder *ClauseBuilder) Or() *ClauseBuilder {
	currentExpression := builder.clause
	builder.processNext = func(expression elements.Expression) elements.Expression {
		return ClauseCreator.Or(currentExpression, expression)
	}

	return builder
}
