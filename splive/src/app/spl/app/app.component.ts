import { AfterViewInit, Component, OnDestroy, OnInit } from '@angular/core';
import { NavigationCancel, NavigationEnd, NavigationStart, Router } from '@angular/router';
import { Subscription } from 'rxjs';

import { USER_CATEGORY } from '../../shared/app-common-constants';
import { AppSpecificDataProvider } from '../../shared/app-specific-data-provider';
import { EnvironmentProvider } from '../../shared/environment-provider';
import { AppDataStoreService } from '../../shared/services/app-data-store/app-data-store-service';
import { LoginHandlerService } from '../../shared/services/login-handler.service';
import { environment } from '../environments/environment';
import { APP_IN_MEMORY_STORE_KEYS, APP_LOCAL_STORAGE_KEYS, APP_ROUTES, SIDE_MENU_LINKS } from './app-constants';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})

export class AppComponent implements OnInit, AfterViewInit, OnDestroy {
  title = 'servicepoint.live';
  routerEventSubscription: Subscription;
  loading: boolean;
  constructor(private loginHandlerService: LoginHandlerService,
    private router: Router,
    private appDataStoreService: AppDataStoreService) { }

  ngOnInit() {

    this.populateEnvironmentProvider();
    this.initAppDataStoreService();
    this.populateAppSpecificDataProvider();
    this.loginHandlerService.init();
  }

  populateEnvironmentProvider() {
    EnvironmentProvider.production = environment.production;
    EnvironmentProvider.baseurl = environment.baseurl;
    EnvironmentProvider.prodcode = environment.prodcode;
  }

  populateAppSpecificDataProvider() {
    AppSpecificDataProvider.sideMenuRoutes = SIDE_MENU_LINKS;
    AppSpecificDataProvider.userCateory = USER_CATEGORY.OSU;
    AppSpecificDataProvider.createRouteMap(APP_ROUTES);
  }

  initAppDataStoreService() {
    this.appDataStoreService.appInMemoryStoreKeys = APP_IN_MEMORY_STORE_KEYS;
    this.appDataStoreService.appLocalStorageKeys = APP_LOCAL_STORAGE_KEYS;
    this.appDataStoreService.init();
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
