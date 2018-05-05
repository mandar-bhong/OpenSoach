import { AfterViewInit, Component, OnDestroy, OnInit } from '@angular/core';
import { NavigationCancel, NavigationEnd, NavigationStart, Router } from '@angular/router';
import { Subscription } from 'rxjs/Subscription';

import { USER_CATEGORY } from '../../shared/app-common-constants';
import { AppSpecificDataProvider } from '../../shared/app-specific-data-provider';
import { EnvironmentProvider } from '../../shared/environment-provider';
import { LoginHandlerService } from '../../shared/services/login-handler.service';
import { environment } from '../environments/environment';
import { APP_ROUTES, SIDE_MENU_LINKS } from './app-constants';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})

export class AppComponent implements OnInit, AfterViewInit, OnDestroy {
  title = 'servicepoint.Live';
  routerEventSubscription: Subscription;
  loading: boolean;
  constructor(private loginHandlerService: LoginHandlerService,
    private router: Router) { }

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
    AppSpecificDataProvider.userCateory = USER_CATEGORY.CU;
    AppSpecificDataProvider.createRouteMap(APP_ROUTES);
  }

  ngAfterViewInit() {
    this.routerEventSubscription = this.router.events
      .subscribe((event) => {
        if (event instanceof NavigationStart) {
          this.loading = true;
        } else if (
          event instanceof NavigationEnd ||
          event instanceof NavigationCancel
        ) {
          this.loading = false;
        }
      });
  }

  ngOnDestroy() {
    if (this.routerEventSubscription) {
      this.routerEventSubscription.unsubscribe();
    }
  }
}
