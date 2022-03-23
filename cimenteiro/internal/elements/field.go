package elements

import "fmt"

type FieldExpression struct {
	Name  string
	Alias string
}

func (field FieldExpression) Build() string {
	if len(field.Alias) <= 0 {
		return field.Name
	}
	return fmt.Sprintf("%s AS %s", field.Name, field.Alias)
}
