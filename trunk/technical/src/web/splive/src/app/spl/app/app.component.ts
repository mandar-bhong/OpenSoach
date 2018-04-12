import { Component, OnInit } from '@angular/core';

import { EnvironmentProvider } from '../../shared/environment-provider';
import { LoginStatusService } from '../../shared/services/login-status.service';
import { environment } from '../environments/environment';
import { AppSpecificDataProvider } from '../../shared/app-specific-data-provider';
import { APP_ROUTES } from './app-constants';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})

export class AppComponent implements OnInit {
  title = 'ServicePoint.Live';

  constructor(private loginStatusService: LoginStatusService) { }

  ngOnInit() {

    this.populateEnvironmentProvider();
    this.populateAppSpecificDataProvider();
    this.loginStatusService.init();
  }

  populateEnvironmentProvider() {
    EnvironmentProvider.production = environment.production;
    EnvironmentProvider.baseurl = environment.baseurl;
    EnvironmentProvider.appbaseurl = environment.appbaseurl;
    EnvironmentProvider.prodcode = environment.prodcode;
  }

  populateAppSpecificDataProvider() {
    AppSpecificDataProvider.approutes = APP_ROUTES;
  }
}
