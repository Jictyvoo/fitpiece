package builder

import (
	"github.com/Wrapped-Owls/fitpiece/cimenteiro/internal/elements"
	"github.com/wrapped-owls/fitpiece/heartcore/failproof"
	"testing"
)

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
