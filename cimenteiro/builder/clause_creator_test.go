package builder

import (
	"github.com/jictyvoo/fitpiece/cimenteiro/builder/expressions"
	"github.com/jictyvoo/fitpiece/cimenteiro/internal/elements"
	"github.com/jictyvoo/fitpiece/heartcore/failproof"
	"testing"
)

func Test_ExpressionCreator(t *testing.T) {
	tableZero := elements.TableName{Name: "table_0"}

	testClause := ClauseCreator.Equal(tableZero.Column("test"), "\"gopher\"")
	failproof.AssertEqual(t, testClause.String(), "`table_0`.`test` = \"gopher\"")

	// Test more complex query
	query := New(tableZero)
	query.Where(
		ClauseCreator.And(
			ClauseCreator.Different("type", 1),
			ClauseCreator.NotIn(expressions.NewValueExpression(2), 1, 2, 3, 4, 5),
		),
	)
	failproof.AssertEqual(t, query.where.String(), "type <> 1 AND 2 NOT IN [1, 2, 3, 4, 5]")

	// Test deeper interaction
	testClause = ClauseCreator.Or(
		ClauseCreator.And(
			ClauseCreator.LessEqual("price", 8000),
			ClauseCreator.InQuery(expressions.NewValueExpression("car"), &query),
		),
		ClauseCreator.Equal("size", "\"BIG\""),
	)

	failproof.AssertEqual(
		t, testClause.String(),
		"price <= 8000 AND car IN (SELECT * FROM table_0 WHERE type <> 1 AND 2 NOT IN [1, 2, 3, 4, 5]) OR size = \"BIG\"",
	)
}
