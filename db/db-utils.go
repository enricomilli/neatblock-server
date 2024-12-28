package db

import (
	"fmt"
	"reflect"
	"strings"
)

func BuildUpsertQuery(tableName string, data interface{}, uniqueKey string) (string, map[string]interface{}) {
	// Get the value and handle pointer types
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// Ensure we're working with a struct
	if v.Kind() != reflect.Struct {
		panic("data must be a struct or pointer to struct")
	}

	t := v.Type()

	fields := []string{}
	placeholders := []string{}
	updates := []string{}
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
			if dbTag != uniqueKey {
				updates = append(updates, fmt.Sprintf("%s = EXCLUDED.%s", dbTag, dbTag))
			}
			values[dbTag] = v.Field(i).Interface()
		}
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s) ON CONFLICT (%s) DO UPDATE SET %s",
		tableName,
		strings.Join(fields, ", "),
		strings.Join(placeholders, ", "),
		uniqueKey,
		strings.Join(updates, ", "),
	)

	return query, values
}

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
