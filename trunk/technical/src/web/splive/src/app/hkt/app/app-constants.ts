import { RoutingModel } from '../../shared/models/ui/routing-model';

export const PROD_CODE = 'SPL_HKT';
export const APP_ROUTES: RoutingModel[] = [
    {
        url: '/dashboard',
        title: 'Dashboard',
        breadcumtitle: 'Dashboard',
        displayinsidemenu: true,
        linkiconcss: 'fa fa-dashboard'
    },
    {
        url: '/devices',
        title: 'Devices',
        breadcumtitle: 'Devices',
        displayinsidemenu: true,
        linkiconcss: 'fa fa-tablet'
    },
    {
        url: '/servicepoints',
        title: 'Service Points',
        breadcumtitle: 'Service Points',
        displayinsidemenu: true,
        linkiconcss: 'fa fa-map-marker'
    },
    {
        url: '/charts',
        title: 'Charts',
        breadcumtitle: 'Charts',
        displayinsidemenu: true,
        linkiconcss: 'fa fa-table'
    },
    {
        url: '/foperators',
        title: 'Operators',
        breadcumtitle: 'Operators',
        displayinsidemenu: true,
        linkiconcss: 'fa fa-meh-o'
    },
    {
        url: '/complaints',
        title: 'Complaints',
        breadcumtitle: 'Complaints',
        displayinsidemenu: true,
        linkiconcss: 'fa fa-flag'
    },
    {
        url: '/reports',
        title: 'Reports',
        breadcumtitle: 'Reports',
        displayinsidemenu: true,
        linkiconcss: 'fa fa-file'
    },
    {
        url: '/users',
        title: 'Users',
        breadcumtitle: 'Users',
        displayinsidemenu: true,
        linkiconcss: 'fa fa-users'
    }
];

