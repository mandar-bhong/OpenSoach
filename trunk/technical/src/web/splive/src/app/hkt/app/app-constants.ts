import { RoutingModel, SideMenuModel } from '../../shared/models/ui/routing-model';

export const PROD_CODE = 'SPL_HKT';
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
        url: '/charts',
        linktitle: 'Charts',
        pagetitle: 'Charts',
        linkiconcss: 'fa fa-table'
    },
    {
        url: '/foperators',
        linktitle: 'Operators',
        pagetitle: 'Operators',
        linkiconcss: 'fa fa-meh-o'
    },
    {
        url: '/complaints',
        linktitle: 'Complaints',
        pagetitle: 'Complaints',
        linkiconcss: 'fa fa-flag'
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
    }
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


