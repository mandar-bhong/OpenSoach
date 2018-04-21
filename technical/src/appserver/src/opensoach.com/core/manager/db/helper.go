package db

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
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

func GetFilterConditionFormModel(model interface{}) string {
	whereCond := ""
	t := reflect.TypeOf(model)
	val := reflect.ValueOf(model)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("db")

		isValueIsNil := reflect.Indirect(val).FieldByName(field.Name).IsNil()

		if isValueIsNil == false {

			fieldElement := reflect.Indirect(val).FieldByName(field.Name).Interface()
			valueElement := reflect.ValueOf(fieldElement)

			if valueElement.Kind() == reflect.Ptr {
				valueElement = valueElement.Elem()
			}

			switch valueElement.Kind() {
			case reflect.Int, reflect.Int16, reflect.Int64, reflect.Int8:
				whereCond = whereCond + tag + " = " + strconv.FormatInt(valueElement.Int(), 10) + " AND "
			case reflect.String:
				whereCond = whereCond + tag + " like " + "\"%" + fmt.Sprintf("%v", valueElement.String()) + "%\"" + " AND "
			}
		}
	}
	whereCond = strings.TrimRight(whereCond, " AND ")
	return whereCond
}
