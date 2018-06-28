export enum CUSTOMER_STATE {
    NOT_SET = 0,
    ACTIVE = 1,
    INACTIVE = 2,
    SUSPENDED = 3
}

export const SERVER_SYSTEM_ERROR_MAX_BOUNDARY = 10000;

export const ROUTE_LOGIN = 'auth/login';
export const ROUTE_HOME = '';
export enum USER_CATEGORY {
    OSU = 1,
    CU = 2
}
export enum USER_STATE {
    NOT_SET = 0,
    ACTIVE = 1,
    INACTIVE = 2,
    SUSPENDED = 3
}
export enum SERVICEPOINT_STATE {
    ACTIVE = 1,
    INACTIVE = 2,
    SUSPENDED = 3
}
export enum USER_GENDER {
    NOT_SELECTED = 0,
    MALE = 1,
    FEMALE = 2
}
export enum DEVICE_STATE {
    NOT_SET = 0,
    ACTIVE = 1,
    INACTIVE = 2,
    SUSPENDED = 3
}

export const DEFAULT_PAGE_MENU = 'DEFAULT_PAGE_MENU';

export enum CUSTOMER_PRODUCT_MAPPING_STATE {
    NOT_SET = 0,
    ACTIVE = 1,
    INACTIVE = 2,
    SUSPENDED = 3
}

export enum USER_PRODUCT_MAPPING_STATE {
    NOT_SET = 0,
    ACTIVE = 1,
    INACTIVE = 2,
    SUSPENDED = 3
}
export enum OPERATOR_STATE {
    ACTIVE = 1,
    INACTIVE = 2
}
export enum OPERATOR_AREA {
    OPEN = 1,
    RESTRICTED = 2
}
export enum COMPLAINT_STATE {
    OPEN = 1,
    CLOSE = 2,
    INPROGRESS = 3
}

export enum SEVERIT_STATE {
    HIGH = 1,
    MEDIUM = 2,
    LOW = 3,
    CRITICAL = 4
}
export const APP_SHARED_DATA_STORE_KEYS = {
    AUTH_TOKEN: 'AUTH_TOKEN',
    USER_ROLE: 'USER_ROLE',
    USER_INFO: 'USER_INFO',
};

export const APP_SHARED_IN_MEMORY_STORE_KEYS: string[] = [
    APP_SHARED_DATA_STORE_KEYS.USER_INFO,
];

export const APP_SHARED_LOCAL_STORAGE_KEYS: string[] = [
    APP_SHARED_DATA_STORE_KEYS.AUTH_TOKEN,
    APP_SHARED_DATA_STORE_KEYS.USER_ROLE
];


