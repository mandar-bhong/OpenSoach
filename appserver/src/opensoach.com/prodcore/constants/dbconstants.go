package constants

const DB_TIME_FORMAT = "2006-01-02 15:04:05"

const DB_USER_STATE_ACTIVE = 1
const DB_USER_CATEGORY_CUSTOMER = 2
const DB_USER_CATEGORY_OS = 1

const DB_DEVICE_STATE_ACTIVE = 1

const DB_DEVICE_CONNECTION_STATE_DISCONNECTED = 0
const DB_DEVICE_CONNECTION_STATE_CONNECTED = 1
const DB_DEVICE_SYNC_STATE_DISCONNECTED = 0

const DB_SERVICE_POINT_STATE_ACTIVE = 1
const DB_SERVICE_POINT_STATE_INACTIVE = 2
const DB_SERVICE_POINT_STATE_SUSPENDED = 3

const LOGGER_LOGGING_LEVEL_ERROR = "Error"
const LOGGER_LOGGING_LEVEL_DEBUG = "Debug"
const LOGGER_LOGGING_LEVEL_INFO = "Info"

const LOGGER_LOGGING_TYPE_STDIO = "Std"
const LOGGER_LOGGING_TYPE_FLUENT = "Fluent"
const LOGGER_LOGGING_TYPE_INFLUXDB = "InfluxDB"

const DB_DOCUMENT_STORAGE_TYPE_FILE_SYSTEM = 1

const DB_NOT_PERSISTANT = 0
const DB_PERSISTANT = 1

// db sync store name
const SYNC_STORE_SERVICE_POINT = "service_point_tbl"
const SYNC_STORE_CONF = "conf_tbl"
const SYNC_STORE_PATIENT_MASTER = "patient_master_tbl"
const SYNC_STORE_PATIENT_CONF = "schedule_tbl"
const SYNC_STORE_PATIENT_ADMISSION = "patient_admission_tbl"
const SYNC_STORE_PERSONAL_DETAILS = "patient_personal_details_tbl"
const SYNC_STORE_MEDICAL_DETAILS = "patient_medical_details_tbl"
const SYNC_STORE_ACTION_TXN = "action_txn_tbl"
const SYNC_STORE_DOCTORS_ORDERS = "doctors_orders_tbl"
const SYNC_STORE_TREATMENT = "treatment_tbl"
const SYNC_STORE_TREATMENT_DOC = "treatment_doc_tbl"
const SYNC_STORE_PATHOLOGY = "pathology_record_tbl"
const SYNC_STORE_PATHOLOGY_DOC = "pathology_record_doc_tbl"
const SYNC_STORE_ACTION = "action_tbl"
const SYNC_STORE_PATIENT_MONITOR_MAPPING_VIEW = "patient_monitor_mapping_view"
