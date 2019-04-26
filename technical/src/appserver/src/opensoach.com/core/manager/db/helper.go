package db

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	gmodels "opensoach.com/models"
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

func GetDBTagFromPropName(datamodel interface{}, propName string) string {

	dbTagName := "db"

	t := reflect.TypeOf(datamodel)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if field.Name == propName {
			dbTagValue := field.Tag.Get(dbTagName)
			return dbTagValue
		}
	}

	return ""
}

func GetFilterConditionFormModel(dataModel ...interface{}) string {

	model := dataModel[0]

	OrAndOperator := " AND "

	if len(dataModel) > 1 {
		OrAndOperator = dataModel[1].(gmodels.FilterConfigModel).OrAndOperator
	}

	whereCond := ""
	t := reflect.TypeOf(model)
	val := reflect.ValueOf(model)

	strFields := []reflect.StructField{}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("db")

		hasValue, valueElement := GetValueElementForField(val, field)

		if hasValue {
			switch valueElement.Kind() {
			case reflect.Int, reflect.Int16, reflect.Int64, reflect.Int8:
				whereCond = whereCond + tag + " = " + strconv.FormatInt(valueElement.Int(), 10) + " AND "
			case reflect.String:
				strFields = append(strFields, field)
			}
		}
	}

	for _, item := range strFields {
		tag := item.Tag.Get("db")
		hasValue, valueElement := GetValueElementForField(val, item)

		if hasValue {
			whereCond = whereCond + tag + " like " + "\"%" + fmt.Sprintf("%v", valueElement.String()) + "%\"" + OrAndOperator
		}
	}

	whereCond = strings.TrimRight(whereCond, " AND ")
	whereCond = strings.TrimRight(whereCond, OrAndOperator)

	return whereCond
}

func GetValueElementForField(val reflect.Value, field reflect.StructField) (bool, reflect.Value) {

	isValueIsNil := false

	switch reflect.Indirect(val).FieldByName(field.Name).Kind() {
	case reflect.Ptr:
		isValueIsNil = reflect.Indirect(val).FieldByName(field.Name).IsNil()
	}

	if isValueIsNil == false {

		fieldElement := reflect.Indirect(val).FieldByName(field.Name).Interface()
		valueElement := reflect.ValueOf(fieldElement)

		if valueElement.Kind() == reflect.Ptr {
			valueElement = valueElement.Elem()
		}

		return true, valueElement
	}

	return false, reflect.Value{}

}
