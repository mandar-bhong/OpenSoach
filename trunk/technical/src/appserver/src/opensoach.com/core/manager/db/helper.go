package db

import (
	"reflect"
)

func GetDBTagFromJSONTag(user interface{}, jsonTag string) string {

	dbTagName := "db"
	jsonTagName := "json"

	t := reflect.TypeOf(user)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		tag := field.Tag.Get(jsonTagName)

		if tag == jsonTag {
			dbTagValue := field.Tag.Get(dbTagName)

			return dbTagValue
		}
	}

	return ""
}
