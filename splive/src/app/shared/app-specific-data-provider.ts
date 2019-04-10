import { USER_CATEGORY } from './app-common-constants';
import { RoutingModel, SideMenuModel } from './models/ui/routing-model';

export class AppSpecificDataProvider {
    static sideMenuRoutes: SideMenuModel[];
    static topMenuRoutes: RoutingModel[];
    static appRoutes: Map<string, RoutingModel>;
    static userCateory: USER_CATEGORY;
    static logoprefix:string;

    static createRouteMap(routes: RoutingModel[]) {
        AppSpecificDataProvider.appRoutes = new Map(routes.map(r => [r.url, r] as [string, RoutingModel]));

        for (let i = 0; i < AppSpecificDataProvider.sideMenuRoutes.length; i++) {
            AppSpecificDataProvider.sideMenuRoutes[i].routingModel = AppSpecificDataProvider.appRoutes.get(
                AppSpecificDataProvider.sideMenuRoutes[i].url);
        }
    }
}
