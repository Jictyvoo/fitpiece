package elements

import "fmt"

type TableName struct {
	Name  string
	alias string
}

func (tableName TableName) String() string {
	if len(tableName.alias) <= 0 {
		return tableName.Name
	}
	return fmt.Sprintf("%s AS %s", tableName.Name, tableName.alias)
}

func (tableName TableName) Column(name string) string {
	return fmt.Sprintf("`%s`.`%s`", tableName.Name, name)
}

func (tableName TableName) ColumnAs(name string, alias string) string {
	return fmt.Sprintf("`%s`.`%s` AS %s", tableName.Name, name, alias)
}
