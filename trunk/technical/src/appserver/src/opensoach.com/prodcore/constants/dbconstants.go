package constants

const DB_USER_STATE_ACTIVE = 1
const DB_USER_CATEGORY_CUSTOMER = 2
const DB_USER_CATEGORY_OS = 1

const DB_DEVICE_STATE_ACTIVE = 1

const DB_DEVICE_CONNECTION_STATE_CONNECTED = 0
const DB_DEVICE_CONNECTION_STATE_DISCONNECTED = 1
const DB_DEVICE_CONNECTION_STATE_UNKNOWN = 2

const DB_SERVICE_POINT_STATE_ACTIVE = 1
const DB_SERVICE_POINT_STATE_INACTIVE = 2
const DB_SERVICE_POINT_STATE_SUSPENDED = 3

const LOGGER_LOGGING_LEVEL_ERROR = "Error"
const LOGGER_LOGGING_LEVEL_DEBUG = "Debug"
const LOGGER_LOGGING_LEVEL_INFO = "Info"

const LOGGER_LOGGING_TYPE_STDIO = "Std"
const LOGGER_LOGGING_TYPE_FLUENT = "Fluent"
const LOGGER_LOGGING_TYPE_INFLUXDB = "InfluxDB"
