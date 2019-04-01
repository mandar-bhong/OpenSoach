import { RoutingModel, SideMenuModel } from '../../shared/models/ui/routing-model';

export const PROD_CODE = 'SPL_HPFT';
export const APP_ROUTES: RoutingModel[] = [
    {
        url: '/dashboard',
        linktitle: 'Dashboard',
        pagetitle: 'Dashboard',
        linkiconcss: 'fa fa-dashboard'
    },
    {
        url: '/devices',
        linktitle: 'Devices',
        pagetitle: 'Devices',
        linkiconcss: 'fa fa-tablet'
    },
    {
        url: '/servicepoints',
        linktitle: 'Wards',
        pagetitle: 'Wards',
        linkiconcss: 'fa fa-map-marker'
    },
    {
        url: '/servicepoints/service-associate',
        linktitle: 'Configure',
        pagetitle: 'Configure Chart',
        linkiconcss: 'fa fa-users'
    },
    // {
    //     url: '/charts',
    //     linktitle: 'Patients',
    //     pagetitle: 'Patients',
    //     linkiconcss: 'fa fa-table'
    // },
    // {
    //     url: '/charts/configure',
    //     linktitle: 'Chart Configuration',
    //     pagetitle: 'Configure Chart',
    //     linkiconcss: 'fa fa-table'
    // },
    // {
    //     url: 'charts/templatelist',
    //     linktitle: 'Templates',
    //     pagetitle: 'Chart Templates',
    //     linkiconcss: 'fa fa-table'
    // },
    {
        url: '/foperators',
        linktitle: 'Medical Attendant',
        pagetitle: 'Medical Attendant',
        linkiconcss: 'fa fa-meh-o'
    },
    {
        url: '/foperators/add',
        linktitle: 'Add',
        pagetitle: 'Add New Operators',
        linkiconcss: 'fa fa-meh-o'
    },
    {
        url: '/foperators/detail',
        linktitle: 'Details',
        pagetitle: 'Operator Details',
        linkiconcss: 'fa fa-meh-o'
    },
    {
        url: '/foperators/associate',
        linktitle: 'Associate',
        pagetitle: 'Operator Associate',
        linkiconcss: 'fa fa-meh-o'
    },
    // {
    //     url: '/complaints',
    //     linktitle: 'Complaints',
    //     pagetitle: 'Complaints',
    //     linkiconcss: 'fa fa-flag'
    // },
    // {
    //     url: '/complaints/detail',
    //     linktitle: 'Details',
    //     pagetitle: 'Complaint Details',
    //     linkiconcss: 'fa fa-meh-o'
    // },
    {
        url: '/reports',
        linktitle: 'Reports',
        pagetitle: 'Reports',
        linkiconcss: 'fa fa-file'
    },
    {
        url: '/users',
        linktitle: 'Users',
        pagetitle: 'Users',
        linkiconcss: 'fa fa-users'
    },
    {
        url: '/users/add-user',
        linktitle: 'Add',
        pagetitle: 'Add New User',
        linkiconcss: 'fa fa-users'
    },
    {
        url: '/users/user-detail',
        linktitle: 'Details',
        pagetitle: 'User Details',
        linkiconcss: 'fa fa-users'
    },
    {
        url: '/users/change-password',
        linktitle: 'Change Password',
        pagetitle: 'Change Password',
        linkiconcss: 'fa fa-users'
    },
    {
        url: '/patients',
        linktitle: 'Patients',
        pagetitle: 'Patients',
        linkiconcss: 'fa fa-users'
    },
    {
        url: '/patients/add',
        linktitle: 'Add',
        pagetitle: 'Add new Patient',
        linkiconcss: 'fa fa-users'
    },
    {
        url: '/patients/patient_chart',
        linktitle: 'File',
        pagetitle: 'Patient File',
        linkiconcss: 'fa fa-users'
    },
    {
        url: '/patients/patient_search',
        linktitle: 'Add',
        pagetitle: 'Add',
        linkiconcss: 'fa fa-users'
    },
    {
        url: '/patients/patient_admission',
        linktitle: 'Admission Details',
        pagetitle: 'Admission Details',
        linkiconcss: 'fa fa-users'
    },


];

export const SIDE_MENU_LINKS: SideMenuModel[] = [
    { url: '/dashboard', level: 0, routingModel: null },
    { url: '/devices', level: 0, routingModel: null },
    { url: '/servicepoints', level: 0, routingModel: null },
    // { url: '/charts', level: 0, routingModel: null },
    { url: '/patients', level: 0, routingModel: null },
    // { url: '/foperators', level: 0, routingModel: null },
    // { url: '/complaints', level: 0, routingModel: null },
    // { url: '/reports', level: 0, routingModel: null },
    { url: '/users', level: 0, routingModel: null }
];

export const TOP_MENU_LINKS = [
];

export const APP_DATA_STORE_KEYS = {
    CUSTOMER_INFO: 'CUSTOMER_INFO',
    CHART_CONFIG: 'CHART_CONFIG'
};

export const APP_IN_MEMORY_STORE_KEYS: string[] = [
    APP_DATA_STORE_KEYS.CUSTOMER_INFO,
    APP_DATA_STORE_KEYS.CHART_CONFIG
];

export const APP_LOCAL_STORAGE_KEYS: string[] = [
];

export enum SERVICE_CONF_TYPE {
    SERVICE_DAILY_CHART = 'SERVICE_DAILY_CHART'
}
export enum PATIENT_STATE {
    NOT_ADMITTED = 0,
    HOSPITALIZE = 1,
    DISCHARGED = 2
}

export enum CHECK_STATE {
    NEW = 0,
    ACKNOWLEDGED = 1
}

export enum PERSON_GENDER {
    NOT_SELECTED = 0,
    MALE = 1,
    FEMALE = 2
}
export enum ConfigCodeType {
    MEDICINE = "Medicine",
    MONITOR = "Monitor",
    OUTPUT = "Output",
    INTAKE = "Intake",
    DOCTOR_ORDERS = 'Doctor-Orders'
}


export enum PATIENT_CHECK_STATE {
    ACTIVE = "Active",
    COMPLETED = "Completed",
    STOPPED = "Stopped"
}


export const FREQUENCY_ZERO = 0
export const FREQUENCY_ONE = 1;

