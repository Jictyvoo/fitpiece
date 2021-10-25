package database

import (
	"fmt"
	"strings"
)

func normalizeValue(value interface{}) string {
	if value == nil {
		return "NULL"
	}
	return fmt.Sprintf("'%v'", value)
}

func parseColumnValues(values map[string]interface{}, delimiter string) [2]string {
	count := 0
	valueBuilder := strings.Builder{}
	keyBuilder := strings.Builder{}
	for key, value := range values {
		if count > 0 {
			keyBuilder.WriteString(delimiter)
			valueBuilder.WriteString(delimiter)
		}
		keyBuilder.WriteString(key)
		strValue := normalizeValue(value)
		valueBuilder.WriteString(strValue)
		count++
	}
	return [2]string{keyBuilder.String(), valueBuilder.String()}
}

func valuePerColumn(elements map[string]interface{}, delimiter string) string {
	count := 0
	updateBuilder := strings.Builder{}
	for key, value := range elements {
		if count > 0 {
			updateBuilder.WriteString(delimiter)
		}
		updateBuilder.WriteString(fmt.Sprintf("`%s` = %s", key, normalizeValue(value)))
		count++
	}
	return updateBuilder.String()
}
