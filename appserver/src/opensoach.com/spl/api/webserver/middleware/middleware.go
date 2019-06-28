package middleware

import (
	lmodels "opensoach.com/spl/api/models"
)

func Init(config *lmodels.WebServerConfiguration) {

	registerRouters(config)
}
