export class RoutingModel {
    url: string;
    linktitle: string;
    linkiconcss: string;
    pagetitle: string;
}

export class SideMenuModel {
    url: string;
    level: number;
    routingModel: RoutingModel;
}
