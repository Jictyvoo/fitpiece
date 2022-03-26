package builder

import (
	"github.com/Wrapped-Owls/fitpiece/cimenteiro/internal/elements"
	"github.com/wrapped-owls/fitpiece/heartcore/failproof"
	"testing"
)

func compareJoins(a, b elements.JoinClause) int {
	return a.Compare(b)
}

func Test_QueryFields(t *testing.T) {
	tableObj := elements.TableName{Name: "Test"}
	query := New(tableObj)

	testFields := [...]elements.FieldExpression{
		{Name: tableObj.Column("avocado")},
		{Name: "lemon"},
		{Name: "juice"},
		{Name: elements.TableName{Name: "banana"}.Column("fruit")},
	}

	query.Fields(testFields[0].Name, testFields[1].Name, testFields[2].Name, testFields[3].Name)
	failproof.AssertEqualSlice(t, query.fields, testFields[:])
}

func Test_QueryFieldsAs(t *testing.T) {
	tableOne := elements.TableName{Name: "table_1"}
	query := New(tableOne)

	testFields := [...]elements.FieldExpression{
		{Name: tableOne.Column("panda"), Alias: "bear"},
		{Name: "count(*)", Alias: "total"},
		{Name: elements.TableName{Name: "table_2"}.Column("name"), Alias: "cats"},
	}

	query.FieldsAs(
		testFields[0].Name, testFields[0].Alias,
		testFields[1].Name, testFields[1].Alias,
		testFields[2].Name, testFields[2].Alias,
	)
	failproof.AssertEqualSlice(t, query.fields, testFields[:])
}

func Test_QueryFieldsAs_OddParameters(t *testing.T) {
	query := New(elements.TableName{Name: "table_0"})

	testFields := [...]elements.FieldExpression{
		{Name: elements.TableName{Name: "table_1"}.Column("panda"), Alias: "bear"},
		{Name: "count(*)", Alias: "total"},
		{Name: elements.TableName{Name: "table_2"}.Column("name"), Alias: "cats"},
	}

	query.FieldsAs(
		testFields[0].Name, testFields[0].Alias,
		testFields[1].Name, testFields[1].Alias,
		testFields[2].Name,
	)
	failproof.AssertEqualSlice(
		t, query.fields,
		[]elements.FieldExpression{
			testFields[0], testFields[1],
		},
	)
}

func Test_QueryBuildJoin(t *testing.T) {
	tableZero := elements.TableName{Name: "table_0"}
	tableOne := elements.TableName{Name: "table_1"}
	query := New(tableZero)
	testJoins := elements.JoinClause{
		JoinType: elements.JoinAll,
		Table:    tableOne,
		On: elements.Clause{
			FirstHalf:  elements.FieldExpression{Name: tableZero.Column("id")},
			Operator:   elements.OperatorEqual,
			SecondHalf: elements.FieldExpression{Name: tableOne.Column("id")},
		},
	}

	query.CrossJoin(tableZero, tableOne, "id", "id")
	if failproof.AssertEqual(t, len(query.joins), 1) {
		failproof.AssertEqualCompare(t, compareJoins, query.joins[0], testJoins)
		failproof.AssertEqual(t, query.joins[0].Build(), testJoins.Build())
		failproof.AssertEqual(t, testJoins.Build(), "JOIN table_1 ON `table_0`.`id` = `table_1`.`id`")
	}
}

func Test_QueryJoinTypes(t *testing.T) {
	tableZero := elements.TableName{Name: "table_0"}
	tableOne := elements.TableName{Name: "table_1"}
	query := New(tableZero)

	query.CrossJoin(tableZero, tableOne, "id", "id").
		InnerJoin(tableZero, tableOne, "id", "id").
		OuterJoin(tableZero, tableOne, "id", "id").
		LeftJoin(tableZero, tableOne, "id", "id").
		RightJoin(tableZero, tableOne, "id", "id")

	if failproof.AssertEqual(t, len(query.joins), 5) {
		failproof.AssertEqual(t, query.joins[0].JoinType, elements.JoinAll)
		failproof.AssertEqual(t, query.joins[1].JoinType, elements.JoinInner)
		failproof.AssertEqual(t, query.joins[2].JoinType, elements.JoinOuter)
		failproof.AssertEqual(t, query.joins[3].JoinType, elements.JoinLeft)
		failproof.AssertEqual(t, query.joins[4].JoinType, elements.JoinRight)
	}
}
