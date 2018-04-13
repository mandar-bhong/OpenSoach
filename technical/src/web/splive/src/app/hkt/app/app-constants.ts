import { RoutingModel } from '../../shared/models/ui/routing-model';

export const PROD_CODE = 'SPL_HKT';
export const APP_ROUTES: RoutingModel[] = [
    {
        url: '/dashboard',
        title: 'Dashboard',
        breadcumtitle: 'Dashboard',
        linkiconcss: 'fa fa-dashboard'
    },
    {
        url: '/devices',
        title: 'Devices',
        breadcumtitle: 'Devices',
        linkiconcss: 'fa fa-tablet'
    },
    {
        url: '/servicepoints',
        title: 'Service Points',
        breadcumtitle: 'Service Points',
        linkiconcss: 'fa fa-map-marker'
    },
    {
        url: '/charts',
        title: 'Charts',
        breadcumtitle: 'Charts',
        linkiconcss: 'fa fa-table'
    },
    {
        url: '/foperators',
        title: 'Operators',
        breadcumtitle: 'Operators',
        linkiconcss: 'fa fa-meh-o'
    },
    {
        url: '/complaints',
        title: 'Complaints',
        breadcumtitle: 'Complaints',
        linkiconcss: 'fa fa-flag'
    },
    {
        url: '/reports',
        title: 'Reports',
        breadcumtitle: 'Reports',
        linkiconcss: 'fa fa-file'
    },
    {
        url: '/users',
        title: 'Users',
        breadcumtitle: 'Users',
        linkiconcss: 'fa fa-users'
    }
];

