import { RoutingModel } from '../../shared/models/ui/routing-model';

export const PROD_CODE = 'SPL';
export const APP_ROUTES: RoutingModel[] = [
    {
        url: '/dashboard',
        title: 'Dashboard',
        breadcumtitle: 'Dashboard',
        displayinsidemenu: true,
        linkiconcss: 'fa fa-dashboard'
    },
    {
        url: '/corporates',
        title: 'Corporates',
        breadcumtitle: 'Corporates',
        displayinsidemenu: true,
        linkiconcss: 'fa fa-building-o'
    },
    {
        url: '/customers',
        title: 'Customers',
        breadcumtitle: 'Customers',
        displayinsidemenu: true,
        linkiconcss: 'fa fa-users'
    },
    {
        url: '/users',
        title: 'Users',
        breadcumtitle: 'Users',
        displayinsidemenu: true,
        linkiconcss: 'fa fa-users'
    },
    {
        url: '/devices',
        title: 'Devices',
        breadcumtitle: 'Devices',
        displayinsidemenu: true,
        linkiconcss: 'fa fa-tablet'
    }
];
