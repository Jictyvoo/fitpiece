package db

import "github.com/wrapped-owls/fitpiece/cimenteiro/internal/elements"

type DatabaseAccess struct {
	daoClasses map[string]DatabaseAccessObject
}

func (db *DatabaseAccess) generateDao(tableName string) (DatabaseAccessObject, error) {
	if daoClass, ok := db.daoClasses[tableName]; ok {
		return daoClass, nil
	}
	newDao := DatabaseAccessObject{
		tableName: elements.TableName{
			Name: tableName,
		},
		columns: []elements.FieldExpression{},
	}
	db.daoClasses[tableName] = newDao
	return newDao, nil
}
