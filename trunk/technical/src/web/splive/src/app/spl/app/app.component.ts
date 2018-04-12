import { Component, OnInit } from '@angular/core';

import { EnvironmentProvider } from '../../shared/environment-provider';
import { LoginStatusService } from '../../shared/services/login-status.service';
import { environment } from '../environments/environment';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  title = 'ServicePoint.Live';

  constructor(private loginStatusService: LoginStatusService) { }

  ngOnInit() {

    EnvironmentProvider.production = environment.production;
    EnvironmentProvider.baseurl = environment.baseurl;
    EnvironmentProvider.appbaseurl = environment.appbaseurl;
    EnvironmentProvider.prodcode = environment.prodcode;
    this.loginStatusService.init();
  }
}
