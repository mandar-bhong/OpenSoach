import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

import { AppDataStoreService } from '../../../services/app-data-store/app-data-store-service';
import { LoginStatusService } from '../../../services/login-status.service';
import { AuthResponse } from '../../../models/api/auth-models';

@Component({
  selector: 'hkt-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  hide = true;
  username: string;
  password: string;
  constructor(private appDataStoreService: AppDataStoreService,
    private loginStatusService: LoginStatusService,
    private router: Router) { }
  ngOnInit() {
    // TODO: Remove after inetgration
    this.username = 'admin@servicepoint.live';
    this.password = 'admin';
  }

  login() {
    // TODO: Call login api
    const authResponse = new AuthResponse();
    authResponse.token = '0123456789';
    this.loginStatusService.login(authResponse);
    this.router.navigate([''], { skipLocationChange: true });
  }

}
