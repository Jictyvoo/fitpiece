package database

import (
	"fmt"
	"strings"
)

type TableObject struct {
	columns   []string
	tableName string
}

func NewDAO(columns []string, tableName string) TableObject {
	return TableObject{
		columns:   columns,
		tableName: tableName,
	}
}

func (dao TableObject) Insert(values map[string]interface{}) string {
	generatedAssociative := parseColumnValues(values, ", ")
	columnsString := generatedAssociative[0]
	valuesString := generatedAssociative[1]
	sqlCommand := fmt.Sprintf("INSERT INTO %s(%s) VALUES(%s)", dao.tableName, columnsString, valuesString)
	return sqlCommand
}

func (dao TableObject) Update(elements map[string]interface{}, where string) string {
	updateCommand := valuePerColumn(elements, ", ")
	whereCommand := ""
	if len(where) > 0 {
		whereCommand = "WHERE " + where
	}
	return fmt.Sprintf("UPDATE %s SET %s %s", dao.tableName, updateCommand, whereCommand)
}

func (dao TableObject) Delete(where string) string {
	if len(where) > 0 {
		return fmt.Sprintf("DELETE FROM %s WHERE %s", dao.tableName, where)
	}
	return ""
}

func (dao TableObject) Select(columns []string, where string) string {
	sqlCommand := strings.Builder{}
	sqlCommand.WriteString("SELECT ")
	for count := 0; count < len(columns); count += 1 {
		if count > 0 {
			sqlCommand.WriteString(", ")
		}
		sqlCommand.WriteString(columns[count])
	}
	sqlCommand.WriteString(" FROM " + dao.tableName)
	if len(where) > 0 {
		sqlCommand.WriteString(" WHERE " + where)
	}
	return sqlCommand.String()
}

func (dao TableObject) SelectAll(where string) string {
	whereCommand := ""
	if len(where) > 0 {
		whereCommand = "WHERE " + where
	}
	return fmt.Sprintf("SELECT * FROM %s %s", dao.tableName, whereCommand)
}
