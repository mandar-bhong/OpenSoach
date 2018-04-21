import { Injectable } from '@angular/core';
import { CanActivate, Router } from '@angular/router';

import { LoginStatusProviderService } from './services/login-status-provider.service';

@Injectable()
export class AuthGuard implements CanActivate {

    constructor(private router: Router,
        private loginStatusProviderService: LoginStatusProviderService) { }

    canActivate() {
        console.log('this.router.url', this.router.url);
        if (this.loginStatusProviderService.isLoggedIn) {
            return true;
        } else {
            console.log('no permission');
            return false;
        }
    }
}
