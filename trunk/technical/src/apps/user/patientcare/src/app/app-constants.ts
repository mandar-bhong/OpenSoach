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
    ACTION = 'action_tbl'
}

export enum SYNC_TYPE {
    NONE = 0,
    SYNC_FROM_SERVER = 1,
    SYNC_TO_SERVER = 2,
    SYNC_TO_AND_FROM_SERVER = 3
}

export enum SYNC_STATE {
    IN_SYNC = 0,
    NOT_IN_SYNC = 1
}
