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
        url: '/customers/update-customer',
        linktitle: 'Details',
        pagetitle: 'Customer Details',
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
        linktitle: 'Add User',
        pagetitle: 'Add New User',
        linkiconcss: 'fa fa-users'
    },
    {
        url: '/users/user-detail',
        linktitle: 'User Detail',
        pagetitle: 'User Detail',
        linkiconcss: 'fa fa-users'
    },
    {
        url: '/devices',
        linktitle: 'Devices',
        pagetitle: 'Devices',
        linkiconcss: 'fa fa-tablet'
    }
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
