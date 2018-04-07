import { Injectable } from '@angular/core';
import { CanActivate, Router } from '@angular/router';

import { LoginStatusService } from './services/login-status.service';

@Injectable()
export class AuthGuard implements CanActivate {

    constructor(private router: Router,
        private loginStatusService: LoginStatusService) { }

    canActivate() {
        console.log('this.router.url', this.router.url);
        if (this.loginStatusService.isLoggedIn) {
            return true;
        } else {
            console.log('no permission');
            this.router.navigate(['auth/login'], { skipLocationChange: true });
            return false;
        }
    }
}
