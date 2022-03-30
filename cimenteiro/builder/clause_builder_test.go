package builder

import (
	"github.com/jictyvoo/fitpiece/cimenteiro/builder/expressions"
	"github.com/jictyvoo/fitpiece/cimenteiro/internal/elements"
	"github.com/jictyvoo/fitpiece/heartcore/failproof"
	"testing"
	"time"
)

func Test_ExpressionBuilder(t *testing.T) {
	tableZero := elements.TableName{Name: "table_0"}

	testClause := CreateClause().
		Equal(tableZero.Column("test"), "\"gopher\"").
		Clause()
	failproof.AssertEqual(t, testClause.String(), "`table_0`.`test` = \"gopher\"")
}

func Test_ExpressionBuilder_Not(t *testing.T) {
	testDate := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)

	testClause := CreateClause().
		GreaterThan("color", 0xFF7c458f).
		Or().
		Different("created_at", testDate).
		And().
		Not().
		LessThan("price", 5734.89).
		Clause()

	sqlStr, args := testClause.StringPlaceholder("?")
	failproof.AssertEqual(t, sqlStr, "((color > ? OR created_at <> ?) AND NOT price < ?)")
	failproof.AssertEqualCompare[[]any](
		t, compareAnySlice,
		args, []any{0xFF7c458f, testDate, 5734.89},
	)
}

func Test_ExpressionBuilder_AndOr(t *testing.T) {
	query := New(elements.TableName{Name: "table_0"})

	testClause := CreateClause().
		Different("type", 1).
		And().
		NotIn(expressions.NewValueExpression(2), 1, 2, 3, 4, 5).
		Clause()

	query.Where(testClause)
	failproof.AssertEqual(t, query.where.String(), "type <> 1 AND 2 NOT IN [1, 2, 3, 4, 5]")

	// Test deeper interaction
	testClause = CreateClause().
		LessEqual("price", 8000).
		And().
		InQuery(expressions.NewValueExpression("car"), &query).
		Or().
		Equal("size", "\"BIG\"").
		Clause()

	failproof.AssertEqual(
		t, testClause.String(),
		"price <= 8000 AND car IN (SELECT * FROM table_0 WHERE type <> 1 AND 2 NOT IN [1, 2, 3, 4, 5]) OR size = \"BIG\"",
	)
}
