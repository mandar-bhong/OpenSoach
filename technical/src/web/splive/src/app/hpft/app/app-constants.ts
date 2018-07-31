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
        linktitle: 'Service Points',
        pagetitle: 'Service Points',
        linkiconcss: 'fa fa-map-marker'
    },
    {
        url: '/servicepoints/service-associate',
        linktitle: 'Configure',
        pagetitle: 'Configure Chart',
        linkiconcss: 'fa fa-users'
    },
    {
        url: '/charts',
        linktitle: 'Charts',
        pagetitle: 'Charts',
        linkiconcss: 'fa fa-table'
    },
    {
        url: '/charts/configure',
        linktitle: 'Chart Configuration',
        pagetitle: 'Configure Chart',
        linkiconcss: 'fa fa-table'
    },
    {
        url: 'charts/templatelist',
        linktitle: 'Templates',
        pagetitle: 'Chart Templates',
        linkiconcss: 'fa fa-table'
    },
    {
        url: '/foperators',
        linktitle: 'Operators',
        pagetitle: 'Operators',
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
    {
        url: '/complaints',
        linktitle: 'Complaints',
        pagetitle: 'Complaints',
        linkiconcss: 'fa fa-flag'
    },
    {
        url: '/complaints/detail',
        linktitle: 'Details',
        pagetitle: 'Complaint Details',
        linkiconcss: 'fa fa-meh-o'
    },
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
];

export const SIDE_MENU_LINKS: SideMenuModel[] = [
    { url: '/dashboard', level: 0, routingModel: null },
    { url: '/devices', level: 0, routingModel: null },
    { url: '/servicepoints', level: 0, routingModel: null },
    { url: '/charts', level: 0, routingModel: null },
    { url: '/foperators', level: 0, routingModel: null },
    { url: '/complaints', level: 0, routingModel: null },
    { url: '/reports', level: 0, routingModel: null },
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


