package builder

import (
	"github.com/wrapped-owls/fitpiece/cimenteiro/internal/elements"
	"github.com/wrapped-owls/fitpiece/heartcore/failproof"
	"testing"
)

func Test_ExpressionBuilder(t *testing.T) {
	tableZero := elements.TableName{Name: "table_0"}
	query := New(tableZero)

	query.Where(query.Equal(tableZero.Column("test"), "\"gopher\""))
	failproof.AssertEqual(t, query.where.Build(), "`table_0`.`test` = \"gopher\"")

	// Test more complex query
	query.Where(
		query.And(
			query.Different("type", 1),
			query.NotIn(ValueExpression[int]{value: 2}, 1, 2, 3, 4, 5),
		),
	)
	failproof.AssertEqual(t, query.where.Build(), "type <> 1 AND 2 NOT IN [1,2,3,4,5]")

	// Test deeper interaction
	secondQuery := New(tableZero)
	secondQuery.Where(
		secondQuery.Or(
			secondQuery.And(
				secondQuery.LessEqual("price", 8000),
				secondQuery.InQuery(ValueExpression[string]{"car"}, &query),
			),
			secondQuery.Equal("size", "\"BIG\""),
		),
	)
	failproof.AssertEqual(
		t, secondQuery.where.Build(),
		"price <= 8000 AND car IN (SELECT * FROM table_0 WHERE type <> 1 AND 2 NOT IN [1,2,3,4,5]) OR size = \"BIG\"",
	)
}
