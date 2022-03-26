package db

import "github.com/wrapped-owls/fitpiece/cimenteiro/internal/elements"

type DatabaseAccessObject struct {
	tableName elements.TableName
	columns   []elements.FieldExpression
}
