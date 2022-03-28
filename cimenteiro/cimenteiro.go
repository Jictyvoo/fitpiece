package cimenteiro

import (
	"database/sql"
	"github.com/jictyvoo/fitpiece/cimenteiro/builder"
	"github.com/jictyvoo/fitpiece/cimenteiro/internal/elements"
)

type Table = elements.TableName

// Cimenteiro is the object that allows queries created using builder.QueryBuilder to
// run safely in the database.
//
// In order to work well, the object instance needs a placeholder.
// This placeholder is used with builder.QueryBuilder and builder.PlaceholderSqlGenerator, to create
// a sql-injection safe query
type Cimenteiro struct {
	db          *sql.DB
	builder     *builder.QueryBuilder
	PlaceHolder string
}

// New creates a new instance of Cimenteiro struct, adding the private field sql.DB,
// and setting '?' as the default placeholder
func New(db *sql.DB) Cimenteiro {
	return Cimenteiro{
		db:          db,
		PlaceHolder: "?",
	}
}

func (cimenteiro *Cimenteiro) generator() builder.PlaceholderSqlGenerator {
	return builder.PlaceholderSqlGenerator{
		Placeholder: cimenteiro.PlaceHolder,
		Query:       cimenteiro.builder,
	}
}

// Attach takes a Table and creates a builder.QueryBuilder based on the given table,
// which is then added to Cimenteiro instance.
//
// The created QueryBuilder is returned to allow users to modify it.
func (cimenteiro *Cimenteiro) Attach(table Table) *builder.QueryBuilder {
	if len(cimenteiro.PlaceHolder) < 1 {
		cimenteiro.PlaceHolder = "?"
	}
	queryBuilder := builder.New(table)
	cimenteiro.builder = &queryBuilder
	return cimenteiro.builder
}

// Select uses the builder.PlaceholderSqlGenerator to generate the query
// and run it using the sql.DB object inside the Cimenteiro instance.
//
// Returns the sql.Rows or an error.
func (cimenteiro Cimenteiro) Select() (*sql.Rows, error) {
	generator := cimenteiro.generator()
	sqlStatement, values := generator.Select()
	return cimenteiro.db.Query(sqlStatement, values...)
}

// Update uses the builder.PlaceholderSqlGenerator to generate the query
// and run it using the sql.DB object inside the Cimenteiro instance.
//
// Returns an sql.Result object, that contains the number of rows affected by the query
func (cimenteiro Cimenteiro) Update(args map[string]any) (sql.Result, error) {
	generator := cimenteiro.generator()
	sqlStatement, values := generator.Update(args)
	return cimenteiro.db.Exec(sqlStatement, values...)
}

// Insert uses the builder.PlaceholderSqlGenerator to generate the query
// and run it using the sql.DB object inside the Cimenteiro instance.
//
// Returns an sql.Result object, that contains the number of rows affected by the query
func (cimenteiro Cimenteiro) Insert(args []any) (sql.Result, error) {
	generator := cimenteiro.generator()
	sqlStatement, values := generator.Insert(args...)
	return cimenteiro.db.Exec(sqlStatement, values...)
}

// Delete uses the builder.PlaceholderSqlGenerator to generate the query
// and run it using the sql.DB object inside the Cimenteiro instance.
//
// Returns an sql.Result object, that contains the number of rows affected by the query
func (cimenteiro Cimenteiro) Delete() (sql.Result, error) {
	generator := cimenteiro.generator()
	sqlStatement, values := generator.Delete()
	return cimenteiro.db.Exec(sqlStatement, values...)
}
