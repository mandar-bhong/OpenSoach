package models

const SESSION_CLIENT_HEADER_KEY string = "Authorization"

const MOD_OPER_SUCCESS int = 0
const MOD_OPER_ERR_SERVER int = 1000
const MOD_OPER_ERR_DATABASE int = 1001
const MOD_OPER_ERR_INPUT_CLIENT_DATA int = 1501
const MOD_OPER_ERR_USER_TOKEN_NOT_AVAILABLE int = 1502
const MOD_OPER_ERR_USER_SESSION_NOT_AVAILABLE int = 2000
const MOD_OPER_UNAUTHORIZED int = 2001

const MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND int = 5000 // Database record not found or no record affected due to same data

const CACHE_KEY_PREFIX_CPM_ID string = "MST_CPM_ID_"

const PRODUCT_TYPE_HKT string = "HKT"

const DEVICE_PROCESSING_AUTH_TOKEN_NOT_FOUND = 3001
