package cimenteiro

import (
	"github.com/Wrapped-Owls/fitpiece/cimenteiro/builder"
	"github.com/Wrapped-Owls/fitpiece/cimenteiro/internal/elements"
)

type Table = elements.TableName

func NewQueryBuilder(table Table) builder.QueryBuilder {
	return builder.New(table)
}
