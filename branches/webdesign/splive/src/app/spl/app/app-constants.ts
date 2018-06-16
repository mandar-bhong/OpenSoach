import { RoutingModel, SideMenuModel } from '../../shared/models/ui/routing-model';

export const PROD_CODE = 'SPL';
export const APP_ROUTES: RoutingModel[] = [
    {
        url: '/dashboard',
        linktitle: 'Dashboard',
        pagetitle: 'Dashboard',
        linkiconcss: 'fa fa-dashboard'
    },
    {
        url: '/corporates',
        linktitle: 'Corporates',
        pagetitle: 'Corporates',
        linkiconcss: 'fa fa-building-o'
    },
    {
        url: '/corporates/add',
        linktitle: 'Add',
        pagetitle: 'Add New Corporates',
        linkiconcss: 'fa fa-building-o'
    },
    {
        url: '/corporates/update',
        linktitle: 'Details',
        pagetitle: 'Corporates Details',
        linkiconcss: 'fa fa-building-o'
    },
    {
        url: '/customers',
        linktitle: 'Customers',
        pagetitle: 'Customers',
        linkiconcss: 'fa fa-users'
    },
    {
        url: '/customers/add',
        linktitle: 'Add',
        pagetitle: 'Add New Customer',
        linkiconcss: ''
    },
    {
        url: '/customers/masterupdate',
        linktitle: 'Master Details',
        pagetitle: 'Customer Master Details',
        linkiconcss: 'fa fa-building-o'
    },
    {
        url: '/customers/update',
        linktitle: 'Details',
        pagetitle: 'Customer Details',
        linkiconcss: ''
    },
    {
        url: '/customers/products',
        linktitle: 'Products',
        pagetitle: 'Associated Products',
        linkiconcss: ''
    },
    {
        url: '/customers/servicepoint',
        linktitle: 'Servicepoint',
        pagetitle: 'Servicepoint',
        linkiconcss: ''
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
        url: '/users/masterdetails',
        linktitle: 'Master Details',
        pagetitle: 'User Master Details',
        linkiconcss: 'fa fa-users'
    },
    {
        url: '/users/user-detail',
        linktitle: 'Details',
        pagetitle: 'User Details',
        linkiconcss: 'fa fa-users'
    },
    {
        url: '/users/products',
        linktitle: 'Products',
        pagetitle: 'Associated Products',
        linkiconcss: ''
    },
    {
        url: '/devices',
        linktitle: 'Devices',
        pagetitle: 'Devices',
        linkiconcss: 'fa fa-tablet'
    },
    {
        url: '/devices/add',
        linktitle: 'Add',
        pagetitle: 'Add New Device',
        linkiconcss: 'fa fa-users'
    },
    {
        url: '/devices/update',
        linktitle: 'Details',
        pagetitle: 'Device Details',
        linkiconcss: 'fa fa-users'
    },
    {
        url: '/devices/products',
        linktitle: 'Products',
        pagetitle: 'Associated Products',
        linkiconcss: ''
    },
];

export const SIDE_MENU_LINKS: SideMenuModel[] = [
    { url: '/dashboard', level: 0, routingModel: null },
    { url: '/corporates', level: 0, routingModel: null },
    { url: '/customers', level: 0, routingModel: null },
    { url: '/users', level: 0, routingModel: null },
    { url: '/devices', level: 0, routingModel: null }
];

export const TOP_MENU_LINKS = [
];

export const APP_DATA_STORE_KEYS = {
};

export const APP_IN_MEMORY_STORE_KEYS: string[] = [
];

export const APP_LOCAL_STORAGE_KEYS: string[] = [
];
