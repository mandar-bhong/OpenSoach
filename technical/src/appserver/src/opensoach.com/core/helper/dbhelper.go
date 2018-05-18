package helper

import (
	"regexp"

	"github.com/go-sql-driver/mysql"
	gmodels "opensoach.com/models"
)

func DBQueryParamValidate(queryInput string, emptyAllowed bool) bool {

	var pattern *regexp.Regexp
	if emptyAllowed {
		pattern = regexp.MustCompile("^[A-Za-z0-9_]*$")
	} else {
		pattern = regexp.MustCompile("^[A-Za-z0-9_]+$")
	}

	return pattern.MatchString(queryInput)
}

func GetApplicationErrorCodeFromDBError(dbErr error) (errorHandled bool, errorCode int) {

	if err, ok := dbErr.(*mysql.MySQLError); ok {

		switch err.Number {
		case 1062: //Unique key constrain failed
			return true, gmodels.MOD_OPER_ERR_DATABASE_DUPLICATE_ENTRY
		default:
			return false, int(err.Number)
		}
	}
	return false, gmodels.MOD_OPER_ERR_DATABASE
}
