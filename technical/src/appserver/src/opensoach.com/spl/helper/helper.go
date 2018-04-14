package helper

import (
	"opensoach.com/core/manager/db"
)

func GetDBTagFromJSONTag(model interface{}, jsonTag string) string {
	return db.GetDBTagFromJSONTag(model, jsonTag)
}
