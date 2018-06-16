import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';

@Injectable()
export class LoginStatusProviderService {
    isLoggedIn: boolean;
    authToken: string;
    userRole: string;

    logginStatusChanged = new Subject<boolean>();
    changeLoginStatus(value: boolean) {
        this.logginStatusChanged.next(value);
    }
}

