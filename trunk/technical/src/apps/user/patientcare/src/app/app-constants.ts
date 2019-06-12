export enum SERVER_WORKER_MSG_TYPE {
    NONE = 0,
    INIT_SERVER_INTERFACE = 1,
    CONNECT_SERVER_INTERFACE = 2,
    CLOSE_SERVER_INTERFACE = 3,
    SEND_MESSAGE = 4,
    UPLOAD_DOCUMENT_COMPLETED = 5,
}

export enum SERVER_WORKER_EVENT_MSG_TYPE {
    NONE = 0,
    SERVER_CONNECTED = 1,
    SERVER_DISCONNECTED = 2,
    DATA_RECEIVED = 3,
    NOTIFICATION_DATA_RECEIVED = 4,
    UPLOAD_DOCUMENT = 5,
}

export enum APP_MODE {
    NONE = 0,
    SHARED_DEVICE = 1,
    USER_DEVICE = 2
}

export enum SYNC_STORE {
    SERVICE_POINT = 'service_point_tbl',
    CONF = 'conf_tbl',
    PATIENT_MASTER = 'patient_master_tbl',
    SCHEDULE = 'schedule_tbl',
    PATIENT_ADMISSION = 'patient_admission_tbl',
    PERSONAL_DETAILS = 'patient_personal_details_tbl',
    MEDICAL_DETAILS = 'patient_medical_details_tbl',
    ACTION_TXN = 'action_txn_tbl',
    ACTION = 'action_tbl',
    DOCTORS_ORDERS = 'doctors_orders_tbl',
    DOCUMENT = 'document_tbl',
    TREATMENT = 'treatment_tbl',
    TREATMENT_DOC = 'treatment_doc_tbl',
    PATHOLOGY_RECORD = 'pathology_record_tbl',
    PATHOLOGY_RECORD_DOC = 'pathology_record_doc_tbl',
    USER = 'mst_user_tbl',
    PATIENT_MONITOR_MAPPING = 'patient_monitor_mapping_view',
}

export enum DB_SYNC_TYPE {
    NONE = 0,
    SYNC_FROM_SERVER = 1,
    SYNC_TO_SERVER = 2,
    SYNC_TO_AND_FROM_SERVER = 3
}

export enum SYNC_PENDING {
    FALSE = 0,
    TRUE = 1
}

export enum SERVER_SYNC_STATE {
    NONE = 0,
    SEND_AUTH_CMD = 1,
    SEND_AUTH_CMD_SUCCESS = 2,
    SEND_AUTH_CMD_FAILURE = 3, // TODO: Not handle as of now
    READ_SYNC_STORE = 4,
    READ_SYNC_STORE_COMPLETED = 5,
    SYNC_TO_SERVER = 6,
    SYNC_TO_SERVER_COMPLETED = 7,
    SYNC_FROM_SERVER = 8,
    SYNC_FROM_SERVER_COMPLETED = 9,
    DIFFERENTIAL_SYNC_INITIALISE = 10,
    DIFFERENTIAL_SYNC_STARTED = 11,
    DIFFERENTIAL_SYNC_COMPLETED = 12
}

export enum CMD_CATEGORY {
    CMD_CAT_DEV_REGISTRATION = 1,
    CMD_CAT_SYNC = 3,
    CMD_CAT_SERVER_NOTIFICATION = 5
}

export enum CMD_ID {
    CMD_DEV_REGISTRATION = 1,
    CMD_GET_STORE_SYNC = 50,
    CMD_APPLY_STORE_SYNC = 51
}
export const freuencyzero = 0
export const freuencyone = 1;

export enum ConfigCodeType {
    MEDICINE = "Medicine",
    MONITOR = "Monitor",
    OUTPUT = "Output",
    INTAKE = "Intake",
    DOCTOR_ORDERS = 'Doctor-Orders'
}
export enum ActionStatus {
    ACTION_ACTIVE = 0,
    ACTION_DELETED = 1
}
export enum ScheuldeStatus {
    SCHEDULE_ACTIVE = 0,
    SCHEDULE_CANCELLED = 1
}

export enum AdmissionStatus {
    Hospitalized = 1,
    Discharged = 2
}

// dev server
export const API_SPL_BASE_URL = "http://172.105.232.148/api";
export const API_APP_BASE_URL = "http://172.105.232.148:91/api";
export const InfluxDb_Log = 'http://172.105.232.148:8086/write?db=spl';

// prod server
// export const API_SPL_BASE_URL = "http://139.162.75.182:91/api";
// export const API_APP_BASE_URL = "http://139.162.75.182/api";

export const ACTION_MISSED_WINDOW = 3 * 60;
export const ACTION_DELAYED_AFTER = 30;
export const ACTION_NEEDS_ATTENTION = 15;
export const ACTION_FUTURE_AFTER = 8 * 60;

export enum ACTION_STATUS {
    // disabled color. there are no actions to be performed in next 3 hours
    NONE = "NONE",
    // default color. there is a action to be performed in next 3 hours, 
    // however there is a still 15 mins more for the action to be performed
    ACTIVE_NORMAL = "ACTIVE_NORMAL",
    // orange, there is action to be performed in another 15 mins or 
    // action was supposed to be performed in last 30 mins
    ACTIVE_NEEDS_ATTENTION = "ACTIVE_NEEDS_ATTENTION",
    // red color, the action is not performed and exceeded 30 mins.
    ACTIVE_DELAYED = "ACTIVE_DELAYED",
    // action to be performed in more than 3 hours
    ACTIVE_FUTURE = "ACTIVE_FUTURE",
    // actions was supposed to be performed more than 3 hours before
    MISSED = "MISSED"

}
export const GRACE_PERIOD = 10; // in minutes

export enum PERSON_ACCMPANYING_GENDER {
    GENDER_MALE = "Male",
    GENDER_FEMALE = "Female",
    GENDER_NOT_SELECTED = "Not_Seleced",
}
export enum MonitorType {
    BLOOD_PRESSURE = 'Blood Pressure',
}
export const MAXIMUM_SCHEDULE_DURATION = 20;
export const MAX_INTERVAL = 23 * 60;
export const MIN_INTERVAL = 5;
export const NUMBER_OF_TIMES_X_INTERVAL = 23;
export enum BuildMode {
    TESTING = 'Testing',
    DEVELOPMENT = 'Development',
    DEBUG = 'Debug',
    PRODUCTION = 'Production',
    STAGING = 'Satging'
}
export enum MessageType {
    ERROR = 'ERROR',
    WARNING = 'WARNING',
    EXCEPTION = 'EXCEPTION'

}
export enum EnvVaraibles {
    BUILD_MODE = 'buildmode'
}
