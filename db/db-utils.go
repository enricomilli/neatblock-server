package db

import (
	"fmt"
	"reflect"
	"strings"
)

func BuildInsertQuery(tableName string, data interface{}) (string, map[string]interface{}) {
	v := reflect.ValueOf(data)
	t := v.Type()

	fields := []string{}
	placeholders := []string{}
	values := make(map[string]interface{})

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		dbTag := field.Tag.Get("db")
		if dbTag != "" && dbTag != "-" {
			if strings.Contains(dbTag, ",") {
				dbTag = strings.Split(dbTag, ",")[0]
			}
			fields = append(fields, dbTag)
			placeholders = append(placeholders, ":"+dbTag)
			values[dbTag] = v.Field(i).Interface()
		}
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		tableName,
		strings.Join(fields, ", "),
		strings.Join(placeholders, ", "),
	)

	return query, values
}
