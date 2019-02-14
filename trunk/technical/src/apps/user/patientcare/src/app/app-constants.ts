export enum SERVER_WORKER_MSG_TYPE {
    NONE = 0,
    INIT_SERVER_INTERFACE = 1,
    CONNECT_SERVER_INTERFACE = 2,
    CLOSE_SERVER_INTERFACE = 3,
    SEND_MESSAGE = 4
}

export enum SERVER_WORKER_EVENT_MSG_TYPE {
    NONE = 0,
    SERVER_CONNECTED = 1,
    SERVER_DISCONNECTED = 2,
    DATA_RECEIVED = 3,
    NOTIFICATION_DATA_RECEIVED = 4
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
    INTAKE = "Intake"
}