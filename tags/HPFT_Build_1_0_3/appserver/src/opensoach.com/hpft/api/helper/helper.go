package helper

import (
	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/manager/db"
)

func DBQueryParamValidate(queryInput string) bool {
	return ghelper.DBQueryParamValidate(queryInput, true)
}

func GetDBTagFromJSONTag(model interface{}, jsonTag string) string {
	return db.GetDBTagFromJSONTag(model, jsonTag)
}

func GetFilterConditionFormModel(model interface{}) string {
	return db.GetFilterConditionFormModel(model)
}

func GetFilterConditionConfigFormModel(model, filterconfig interface{}) string {
	return db.GetFilterConditionFormModel(model, filterconfig)
}
