import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

import { EnvironmentProvider } from '../../../environment-provider';
import { AuthRequest, AuthResponse } from '../../../models/api/auth-models';
import { AppDataStoreService } from '../../../services/app-data-store/app-data-store-service';
import { AuthService } from '../../../services/auth.service';
import { LoginStatusService } from '../../../services/login-status.service';

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
    private router: Router,
    private authService: AuthService) { }
  ngOnInit() {
    // TODO: Remove after inetgration
    this.username = 'admin@servicepoint.live';
    this.password = 'admin';
  }

  login() {
    // TODO: Call login api
    const authRequest = new AuthRequest();
    authRequest.username = this.username;
    authRequest.password = this.password;
    authRequest.prodcode = EnvironmentProvider.prodcode;

    this.authService.login(authRequest).subscribe(response => {
      console.log('response', response);
    });

    const authResponse = new AuthResponse();
    authResponse.token = '0123456789';
    this.loginStatusService.login(authResponse);
    this.router.navigate([''], { skipLocationChange: true });
  }

}
