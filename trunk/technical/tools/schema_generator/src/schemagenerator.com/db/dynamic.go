package db

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

const MODEL_DB_ATTRIBUTE_TAG = "dbattr"
const MODEL_DB_AUTO_TAG = "auto"
const MODEL_DB_PRIMARY_TAG = "pri"

var queries map[string]string

func init() {
	queries = make(map[string]string)
}

func GetInsertDynamicQuery(tablename string, insStruct interface{}) string {
	query := ""
	key := tablename + "Insert"

	val, ok := queries[key]
	if ok {
		query = val
	} else {
		t := reflect.TypeOf(insStruct)
		query = "INSERT INTO " + tablename + " ("
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			tag := field.Tag.Get("db")
			attrTagVal, _ := field.Tag.Lookup(MODEL_DB_ATTRIBUTE_TAG)
			if strings.Contains(attrTagVal, MODEL_DB_AUTO_TAG) == false {
				query = query + tag + ","
			}
		}
		query = strings.TrimRight(query, ",")
		query = query + " ) values ( "
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			tag := field.Tag.Get("db")
			attrTagVal, _ := field.Tag.Lookup(MODEL_DB_ATTRIBUTE_TAG)
			if strings.Contains(attrTagVal, MODEL_DB_AUTO_TAG) == false {
				query = query + ":" + tag + ","
			}
		}
		query = strings.TrimRight(query, ",")
		query = query + ")"

		queries[key] = query
	}
	// fmt.Println(query)
	return query
}

func GetUpdateDynamicQuery(tablename string, updateStruct interface{}) (error, string) {
	query := ""
	key := tablename + "Update"

	val, ok := queries[key]
	if ok {
		query = val
	} else {
		primaryKey := ""
		t := reflect.TypeOf(updateStruct)

		val := reflect.ValueOf(updateStruct)
		modelName := reflect.Indirect(val).Type().Name()

		query = "UPDATE " + tablename + " SET "
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			tag := field.Tag.Get("db")
			attrTagVal, _ := field.Tag.Lookup(MODEL_DB_ATTRIBUTE_TAG)
			if strings.Contains(attrTagVal, MODEL_DB_PRIMARY_TAG) {
				primaryKey = tag
			} else {
				if strings.Contains(attrTagVal, MODEL_DB_AUTO_TAG) == false {
					query = query + tag + " = :" + tag + ", "
				}
			}
		}
		query = strings.TrimRight(query, ", ")

		if primaryKey != "" {
			query = query + " WHERE " + primaryKey + " = :" + primaryKey
			queries[key] = query
		} else {
			return errors.New(fmt.Sprintf("Unable to find dbattr tag in provided model. Model : %v", modelName)), ""
		}
	}
	// fmt.Println(query)
	return nil, query
}

func GetDeleteDynamicQuery(tablename string, deleteStruct interface{}) (error, string) {
	query := ""
	key := tablename + "Delete"

	val, ok := queries[key]
	if ok {
		query = val
	} else {
		primaryKey := ""
		t := reflect.TypeOf(deleteStruct)

		val := reflect.ValueOf(deleteStruct)
		modelName := reflect.Indirect(val).Type().Name()

		query = "DELETE FROM " + tablename + " WHERE "
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			tag := field.Tag.Get("db")
			attrTagVal, _ := field.Tag.Lookup(MODEL_DB_ATTRIBUTE_TAG)
			if strings.Contains(attrTagVal, MODEL_DB_PRIMARY_TAG) {
				primaryKey = tag
			}
		}
		if primaryKey != "" {
			query = query + primaryKey + " = :" + primaryKey
			queries[key] = query
		} else {
			return errors.New(fmt.Sprintf("Unable to find dbattr tag in provided model. Model : %v", modelName)), ""
		}
	}
	// fmt.Println(query)
	return nil, query
}

func GetSelectAllDynamicQuery(tablename string, destination interface{}) string {
	query := ""
	key := tablename + "SelectAll"

	val, ok := queries[key]
	if ok {
		query = val
	} else {
		val := reflect.ValueOf(destination)
		t := reflect.Indirect(val).Type().Elem()
		query = "Select "
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			tag := field.Tag.Get("db")
			query = query + tag + ","
		}
		query = strings.TrimRight(query, ",")
		query = query + " FROM " + tablename

		queries[key] = query
	}
	// fmt.Println(query)
	return query
}

func GetSelectByIdDynamicQuery(tablename string, destination interface{}) (error, string) {
	query := ""
	key := tablename + "SelectById"

	val, ok := queries[key]
	if ok {
		query = val
	} else {
		primaryKey := ""
		val := reflect.ValueOf(destination)
		t := reflect.Indirect(val).Type().Elem()
		modelName := reflect.Indirect(val).Type().Elem().Name()

		query = "SELECT "
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			tag := field.Tag.Get("db")
			attrTagVal, _ := field.Tag.Lookup(MODEL_DB_ATTRIBUTE_TAG)
			if strings.Contains(attrTagVal, MODEL_DB_PRIMARY_TAG) {
				primaryKey = tag
			}
			query = query + tag + ", "
		}
		query = strings.TrimRight(query, ", ")

		if primaryKey != "" {
			query = query + " FROM " + tablename + " WHERE " + primaryKey + " = ?"
			queries[key] = query
		} else {
			return errors.New(fmt.Sprintf("Unable to find dbattr tag in provided model. Model : %v", modelName)), ""
		}
	}
	// fmt.Println(query)
	return nil, query
}

func GetSelectByFilterDynamicQuery(tablename string, filter interface{}, args ...string) string {
	t := reflect.TypeOf(filter)
	query := "SELECT "
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("db")
		query = query + tag + ", "
	}
	query = strings.TrimRight(query, ", ")
	query = query + " FROM " + tablename + " WHERE "
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("db")
		for _, each := range args {
			if field.Name == each {
				query = query + tag + " =? " + " AND "
			}
		}
	}
	query = strings.TrimRight(query, " AND ")
	// fmt.Println(query)
	return query
}

func GetFilterValues(tablename string, filterStruct interface{}, args ...string) []interface{} {
	var tagValues []interface{}
	val := reflect.ValueOf(filterStruct)
	t := reflect.TypeOf(filterStruct)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		for _, each := range args {
			if field.Name == each {
				value := reflect.Indirect(val).FieldByName(field.Name).Interface()
				tagValues = append(tagValues, value)
			}
		}
	}
	// fmt.Println(tagValues)
	return tagValues
}
