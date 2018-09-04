import { AfterViewInit, Component, OnDestroy, OnInit } from '@angular/core';
import { NavigationCancel, NavigationEnd, NavigationStart, Router } from '@angular/router';
import { Subscription } from 'rxjs';

import { SplConfService } from '../../prod-shared/services/spl-conf.service';
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
    private appDataStoreService: AppDataStoreService,
    private splConfService: SplConfService) { }

  ngOnInit() {

    this.populateEnvironmentProvider();
    this.initAppDataStoreService();
    this.populateAppSpecificDataProvider();
    this.getSplBaseUrl();
  }

  populateEnvironmentProvider() {
    EnvironmentProvider.production = environment.production;
    EnvironmentProvider.appbaseurl = environment.appbaseurl;
    EnvironmentProvider.prodcode = environment.prodcode;
  }

  populateAppSpecificDataProvider() {
    AppSpecificDataProvider.sideMenuRoutes = SIDE_MENU_LINKS;
    AppSpecificDataProvider.userCateory = USER_CATEGORY.CU;
    AppSpecificDataProvider.createRouteMap(APP_ROUTES);
    AppSpecificDataProvider.logoprefix='vehicleservice';
  }

  initAppDataStoreService() {
    this.appDataStoreService.appInMemoryStoreKeys = APP_IN_MEMORY_STORE_KEYS;
    this.appDataStoreService.appLocalStorageKeys = APP_LOCAL_STORAGE_KEYS;
    this.appDataStoreService.init();
  }

  getSplBaseUrl() {
    this.splConfService.getSplBaseUrl().subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        EnvironmentProvider.baseurl = payloadResponse.data.baseurl;
        this.loginHandlerService.init();
      }
    });
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
