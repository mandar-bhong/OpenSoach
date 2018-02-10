package webhelper

import (
	"time"
)

const SESSION_KEY = "Authorization"
const SESSION_TIME_OUT time.Duration = time.Minute * 10

const MOD_OPER_SUCCESS int = 0
const MOD_OPER_ERR_SERVER int = 1000
const MOD_OPER_ERR_DATABASE int = 1001
