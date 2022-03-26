package cimenteiro

import (
	"database/sql"
	"github.com/Wrapped-Owls/fitpiece/cimenteiro/builder"
	"github.com/Wrapped-Owls/fitpiece/cimenteiro/internal/elements"
)

type Table = elements.TableName

type Cimenteiro struct {
	db          *sql.DB
	builder     builder.QueryBuilder
	PlaceHolder string
}

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

func (cimenteiro *Cimenteiro) Attach(table Table) *builder.QueryBuilder {
	if len(cimenteiro.PlaceHolder) < 1 {
		cimenteiro.PlaceHolder = "?"
	}
	cimenteiro.builder = builder.New(table)
	return &cimenteiro.builder
}

func (cimenteiro Cimenteiro) Select() (*sql.Rows, error) {
	generator := cimenteiro.generator()
	sqlStatement, values := generator.Select()
	return cimenteiro.db.Query(sqlStatement, values...)
}

func (cimenteiro Cimenteiro) Update(args map[string]any) (sql.Result, error) {
	generator := cimenteiro.generator()
	sqlStatement, values := generator.Update(args)
	return cimenteiro.db.Exec(sqlStatement, values...)
}

func (cimenteiro Cimenteiro) Insert(args []any) (sql.Result, error) {
	generator := cimenteiro.generator()
	sqlStatement, values := generator.Insert(args...)
	return cimenteiro.db.Exec(sqlStatement, values...)
}

func (cimenteiro Cimenteiro) Delete() (sql.Result, error) {
	generator := cimenteiro.generator()
	sqlStatement, values := generator.Delete()
	return cimenteiro.db.Exec(sqlStatement, values...)
}
