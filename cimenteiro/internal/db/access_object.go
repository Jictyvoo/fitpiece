package db

import "github.com/jictyvoo/fitpiece/cimenteiro/internal/elements"

type DatabaseAccessObject struct {
	tableName elements.TableName
	columns   []elements.FieldExpression
}
