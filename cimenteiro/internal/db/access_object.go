package db

import "github.com/Wrapped-Owls/fitpiece/cimenteiro/internal/elements"

type DatabaseAccessObject struct {
	tableName elements.TableName
	columns   []elements.FieldExpression
}
