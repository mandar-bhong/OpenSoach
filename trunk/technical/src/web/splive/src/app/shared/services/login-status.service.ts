import { Injectable } from '@angular/core';

import { AuthResponse } from '../models/api/auth-models';
import { APP_DATA_STORE_KEYS, AppDataStoreService } from './app-data-store/app-data-store-service';

@Injectable()
export class LoginStatusService {
    isLoggedIn: boolean;
    authToken: string;
    userRole: string;

    constructor(private appDataStoreService: AppDataStoreService) { }

    init() {
        this.authToken = this.appDataStoreService.getDataStore(APP_DATA_STORE_KEYS.AUTH_TOKEN)
            .getObject<string>(APP_DATA_STORE_KEYS.AUTH_TOKEN);
        if (this.authToken) {
            // TODO: Call API for validating token.
            this.isLoggedIn = true;
        }

        this.userRole = this.appDataStoreService.getDataStore(APP_DATA_STORE_KEYS.USER_ROLE)
            .getObject<string>(APP_DATA_STORE_KEYS.USER_ROLE);
    }

    login(authResponse: AuthResponse) {
        this.isLoggedIn = true;
        this.authToken = authResponse.token;
        this.userRole = authResponse.urolecode;

        this.appDataStoreService.getDataStore(APP_DATA_STORE_KEYS.AUTH_TOKEN)
            .setObject<string>(APP_DATA_STORE_KEYS.AUTH_TOKEN, this.authToken);

        this.appDataStoreService.getDataStore(APP_DATA_STORE_KEYS.USER_ROLE)
            .setObject<string>(APP_DATA_STORE_KEYS.USER_ROLE, this.userRole);
    }

    logout() {
        this.isLoggedIn = false;
        this.authToken = null;
        this.userRole = null;

        this.appDataStoreService.getDataStore(APP_DATA_STORE_KEYS.AUTH_TOKEN)
            .removeObject(APP_DATA_STORE_KEYS.AUTH_TOKEN);

        this.appDataStoreService.getDataStore(APP_DATA_STORE_KEYS.USER_ROLE)
            .removeObject(APP_DATA_STORE_KEYS.USER_ROLE);
    }
}
