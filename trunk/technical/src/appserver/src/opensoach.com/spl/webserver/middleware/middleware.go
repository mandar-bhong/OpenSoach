package middleware

import (
	lmodels "opensoach.com/spl/models"
)

func Init(config *lmodels.WebServerConfiguration) {

	registerRouters(config)
}
