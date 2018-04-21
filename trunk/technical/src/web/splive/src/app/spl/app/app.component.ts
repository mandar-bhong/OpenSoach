import { Component, OnInit } from '@angular/core';

import { EnvironmentProvider } from '../../shared/environment-provider';
import { LoginHandlerService } from '../../shared/services/login-handler.service';
import { environment } from '../environments/environment';
import { AppSpecificDataProvider } from '../../shared/app-specific-data-provider';
import { APP_ROUTES, SIDE_MENU_LINKS } from './app-constants';
import { USER_CATEGORY } from '../../shared/app-common-constants';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})

export class AppComponent implements OnInit {
  title = 'ServicePoint.Live';

  constructor(private loginHandlerService: LoginHandlerService) { }

  ngOnInit() {

    this.populateEnvironmentProvider();
    this.populateAppSpecificDataProvider();
    this.loginHandlerService.init();
  }

  populateEnvironmentProvider() {
    EnvironmentProvider.production = environment.production;
    EnvironmentProvider.baseurl = environment.baseurl;
    EnvironmentProvider.appbaseurl = environment.appbaseurl;
    EnvironmentProvider.prodcode = environment.prodcode;
  }

  populateAppSpecificDataProvider() {
    AppSpecificDataProvider.sideMenuRoutes = SIDE_MENU_LINKS;
    AppSpecificDataProvider.userCateory = USER_CATEGORY.OSU;
    AppSpecificDataProvider.createRouteMap(APP_ROUTES);
  }
}
