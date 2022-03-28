package builder

import (
	"github.com/jictyvoo/fitpiece/cimenteiro/builder/expressions"
	"github.com/jictyvoo/fitpiece/cimenteiro/internal/elements"
)

type expressionProcessor = func(expression elements.Expression) elements.Expression

type ClauseBuilder struct {
	clause      elements.Expression
	processNext expressionProcessor
}

func CreateClause() *ClauseBuilder {
	return &ClauseBuilder{}
}

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

func (builder *ClauseBuilder) In(clause elements.Expression, values ...any) *ClauseBuilder {
	builder.registerClause(ClauseCreator.In(clause, values...))
	return builder
}

func (builder *ClauseBuilder) InQuery(clause elements.Expression, subQuery *QueryBuilder) *ClauseBuilder {
	builder.registerClause(ClauseCreator.InQuery(clause, subQuery))
	return builder
}

func (builder *ClauseBuilder) NotIn(clause elements.Expression, values ...any) *ClauseBuilder {
	builder.registerClause(ClauseCreator.NotIn(clause, values...))
	return builder
}

func (builder *ClauseBuilder) NotInQuery(clause elements.Expression, subQuery *QueryBuilder) *ClauseBuilder {
	builder.registerClause(ClauseCreator.NotInQuery(clause, subQuery))
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
