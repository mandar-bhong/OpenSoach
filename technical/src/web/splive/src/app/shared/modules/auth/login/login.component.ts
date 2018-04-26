import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

import { EnvironmentProvider } from '../../../environment-provider';
import { AuthRequest, AuthResponse } from '../../../models/api/auth-models';
import { AppDataStoreService } from '../../../services/app-data-store/app-data-store-service';
import { AuthService } from '../../../services/auth.service';
import { LoginHandlerService } from '../../../services/login-handler.service';

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
    private loginHandlerService: LoginHandlerService,
    private router: Router,
    private authService: AuthService) { }
  ngOnInit() {
  }

  login() {
    const authRequest = new AuthRequest();
    authRequest.username = this.username;
    authRequest.password = this.password;
    authRequest.prodcode = EnvironmentProvider.prodcode;

    this.authService.login(authRequest).subscribe(response => {
      if (response && response.issuccess) {
        this.loginHandlerService.login(response.data);
        this.router.navigate([''], { skipLocationChange: true });
      }
    });
  }
}
