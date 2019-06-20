package helper

import (
	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/manager/db"
)

var SUB_MODULE_NAME = "SPL.helper"

func DBQueryParamValidate(queryInput string) bool {
	return ghelper.DBQueryParamValidate(queryInput, true)
}

func GetDBTagFromJSONTag(model interface{}, jsonTag string) string {
	return db.GetDBTagFromJSONTag(model, jsonTag)
}

func GetFilterConditionFormModel(model interface{}) string {
	return db.GetFilterConditionFormModel(model)
}
