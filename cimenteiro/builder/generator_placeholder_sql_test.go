package builder

import (
	"github.com/jictyvoo/fitpiece/cimenteiro/builder/expressions"
	"github.com/jictyvoo/fitpiece/cimenteiro/internal/elements"
	"github.com/jictyvoo/fitpiece/heartcore/failproof"
	"testing"
	"time"
)

func compareAnySlice(a, b []any) int {
	if len(a) != len(b) {
		return len(a) - len(b)
	}
	difference := 0
	for index := range a {
		if a[index] != b[index] {
			difference++
		}
	}
	return difference
}

func TestPlaceholderSqlGenerator_Select(t *testing.T) {
	tableZero := elements.TableName{Name: "table_0"}
	query := New(tableZero)

	// Test simple query
	query.Where(
		ClauseCreator.Not(
			ClauseCreator.And(
				ClauseCreator.Different("type", "cimento"),
				ClauseCreator.NotIn(expressions.NewValueExpression(49), 10, 20, 30, 40, 50),
			),
		),
	)

	generator := PlaceholderSqlGenerator{Query: &query, Placeholder: "?"}
	sqlStr, args := generator.Select()
	failproof.AssertEqual(t, sqlStr, "SELECT * FROM table_0 WHERE NOT (type <> ? AND ? NOT IN [?, ?, ?, ?, ?])")
	failproof.AssertEqualCompare[[]any](
		t, compareAnySlice,
		args, []any{"cimento", 49, 10, 20, 30, 40, 50},
	)

	// Add fields to query
	query.Fields(
		"id", "name", "type",
	)
	query.FieldsAs(
		"id", "id_field",
		"name", "name_field",
		"type", "type_field",
	)

	sqlStr, args = generator.Select()
	failproof.AssertEqual(
		t, sqlStr,
		"SELECT id, name, type, id AS id_field, name AS name_field, type AS type_field "+
			"FROM table_0 WHERE NOT (type <> ? AND ? NOT IN [?, ?, ?, ?, ?])",
	)
}

func TestPlaceholderSqlGenerator_Select__Join(t *testing.T) {
	tableZero := elements.TableName{Name: "table_0"}
	tableOne := elements.TableName{Name: "table_1"}
	tableTwo := elements.TableName{Name: "table_2"}
	subQuery := New(tableOne)
	subQuery.FieldsAs(
		tableOne.Column("price"), "price_field",
		tableOne.Column("created_at"), "created_at_field",
	)
	subQuery.Where(
		ClauseCreator.Or(
			ClauseCreator.Different("name", "argamassa"),
			ClauseCreator.GreaterThan("id", 630),
		),
	)

	testDate := time.Date(1350, time.July, 7, 11, 36, 8, 49, time.UTC)
	query := New(tableZero)
	query.
		Fields(tableZero.Column("evolution"), "enemy", tableTwo.ColumnAs("book", "book")).
		FieldsAs(
			tableOne.Column("logo"), "logo_path",
			tableZero.Column("created_as"), "first_creation",
		).
		InnerJoinOrigin(tableTwo, "uuid", "zero_id").
		LeftJoin(tableOne, tableTwo, "id", "one_id").
		Where(
			ClauseCreator.Or(
				ClauseCreator.And(
					ClauseCreator.LessEqual("price", 91.5),
					ClauseCreator.InQuery(expressions.NewValueExpression("tijolo"), &subQuery),
				),
				ClauseCreator.Equal(
					"delivered_at",
					testDate,
				),
			),
		)

	generator := PlaceholderSqlGenerator{Query: &query, Placeholder: "?"}
	sqlStr, args := generator.Select()
	failproof.AssertEqual(
		t, sqlStr,
		"SELECT `table_0`.`evolution`, enemy, `table_2`.`book` AS book, `table_1`.`logo` AS logo_path, "+
			"`table_0`.`created_as` AS first_creation "+
			"FROM table_0 INNER JOIN table_2 ON `table_0`.`uuid` = `table_2`.`zero_id` "+
			"LEFT JOIN table_2 ON `table_1`.`id` = `table_2`.`one_id` "+
			"WHERE (price <= ? AND "+
			"? IN (SELECT `table_1`.`price` AS price_field, `table_1`.`created_at` AS created_at_field FROM table_1 WHERE name <> ? OR id > ?)) "+
			"OR delivered_at = ?",
	)
	failproof.AssertEqualCompare[[]any](
		t, compareAnySlice,
		args, []any{91.5, "tijolo", "argamassa", 630, testDate},
	)
}

func TestPlaceholderSqlGenerator_Select__OrderBy(t *testing.T) {
	tableZero := elements.TableName{Name: "table_0"}
	query := New(tableZero)

	generator := PlaceholderSqlGenerator{
		Placeholder: "?",
		Query: query.Where(
			ClauseCreator.Not(
				ClauseCreator.GreaterThan("id", 42),
			),
		).
			OrderBy("id", "name", "value").
			GroupBy("id", "name").
			Having(
				ClauseCreator.NotIn(
					expressions.NewValueExpression("language"), "python", "java", "kotlin", "c++",
				),
			).
			Limit(100).
			Offset(9),
	}

	sqlStr, args := generator.Select()
	failproof.AssertEqual(
		t, sqlStr,
		"SELECT * FROM table_0 WHERE NOT id > ? "+
			"GROUP BY id, name "+
			"HAVING ? NOT IN [?, ?, ?, ?] "+
			"ORDER BY id, name, value ASC "+
			"LIMIT ? OFFSET ?",
	)
	failproof.AssertEqualCompare[[]any](
		t, compareAnySlice,
		args, []any{42, "language", "python", "java", "kotlin", "c++", 100, 9},
	)
}

func TestPlaceholderSqlGenerator_Update(t *testing.T) {
	tableZero := elements.TableName{Name: "table_0"}
	query := New(tableZero)

	// Test simple query
	query.Where(
		ClauseCreator.Or(
			ClauseCreator.In(expressions.NewFieldExpression("pet"), "cat", "dog", "elephant", "tiger", "lion"),
			ClauseCreator.Equal("type", "telha"),
		),
	)

	generator := PlaceholderSqlGenerator{Query: &query, Placeholder: "?"}
	testDate := time.Date(1568, time.May, 19, 5, 18, 59, 26, time.UTC)

	sqlStr, args := generator.Update(map[string]any{
		"name":       "avocado",
		"type":       "fruit",
		"size":       7284,
		"updated_at": testDate,
	})
	failproof.AssertEqual(
		t, sqlStr,
		"UPDATE table_0 SET name = ?, size = ?, type = ?, updated_at = ? WHERE pet IN [?, ?, ?, ?, ?] OR type = ?",
	)
	failproof.AssertEqualCompare[[]any](
		t, compareAnySlice,
		args, []any{"avocado", 7284, "fruit", testDate, "cat", "dog", "elephant", "tiger", "lion", "telha"},
	)
}

func TestPlaceholderSqlGenerator_Insert(t *testing.T) {
	tableZero := elements.TableName{Name: "table_0"}
	query := New(tableZero)
	generator := PlaceholderSqlGenerator{
		Placeholder: "?",
		Query:       query.Fields("name", "type", "size", "created_at"),
	}
	testDate := time.Date(1568, time.May, 19, 5, 18, 59, 26, time.UTC)

	sqlStr, args := generator.Insert("avocado", "fruit", 7284, testDate)
	failproof.AssertEqual(
		t, sqlStr,
		"INSERT INTO table_0 (name, type, size, created_at) VALUES (?,?,?,?)",
	)
	failproof.AssertEqualCompare[[]any](
		t, compareAnySlice,
		args, []any{"avocado", "fruit", 7284, testDate},
	)
}

func TestPlaceholderSqlGenerator_Delete(t *testing.T) {
	tableZero := elements.TableName{Name: "test_table"}
	query := New(tableZero)
	query.Where(
		ClauseCreator.Or(
			ClauseCreator.And(
				ClauseCreator.In(expressions.NewFieldExpression("version"), "1.14", "1.16", "1.17", "1.18"),
				ClauseCreator.Equal("language", "go"),
			),
			ClauseCreator.GreaterEqual("version", "1.13"),
		),
	)

	generator := PlaceholderSqlGenerator{Query: &query, Placeholder: "?"}
	sqlStr, args := generator.Delete()
	failproof.AssertEqual(
		t, sqlStr,
		"DELETE FROM test_table WHERE (version IN [?, ?, ?, ?] AND language = ?) OR version >= ?",
	)
	failproof.AssertEqualCompare[[]any](
		t, compareAnySlice,
		args, []any{"1.14", "1.16", "1.17", "1.18", "go", "1.13"},
	)
}
