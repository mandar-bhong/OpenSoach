import { Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { Subject } from 'rxjs/Subject';

import { ROUTE_HOME, ROUTE_LOGIN, USER_CATEGORY } from '../app-common-constants';
import { AppSpecificDataProvider } from '../app-specific-data-provider';
import { AuthResponse } from '../models/api/auth-models';
import { CustomerInfo } from '../models/ui/customer-models';
import { UserInfo } from '../models/ui/user-models';
import { APP_DATA_STORE_KEYS, AppDataStoreService } from './app-data-store/app-data-store-service';
import { AuthService } from './auth.service';
import { CustomerSharedService } from './customer/customer-shared.service';
import { LoginStatusProviderService } from './login-status-provider.service';
import { UserSharedService } from './user/user-shared.service';

@Injectable()
export class LoginHandlerService {
    userInfoSubject = new Subject<UserInfo>();
    customerInfoSubject = new Subject<CustomerInfo>();
    constructor(private appDataStoreService: AppDataStoreService,
        private router: Router,
        private authService: AuthService,
        private loginStatusProviderService: LoginStatusProviderService,
        private userSharedService: UserSharedService,
        private customerSharedService: CustomerSharedService) {
        this.loginStatusProviderService.logginStatusChanged.subscribe(status => {
            this.logout();
        });
    }

    init() {
        this.loginStatusProviderService.authToken = this.appDataStoreService.getDataStore(APP_DATA_STORE_KEYS.AUTH_TOKEN)
            .getObject<string>(APP_DATA_STORE_KEYS.AUTH_TOKEN);
        if (this.loginStatusProviderService.authToken) {
            this.validateAuthToken();
        } else {
            this.router.navigate([ROUTE_LOGIN], { skipLocationChange: true });
        }

        this.loginStatusProviderService.userRole = this.appDataStoreService.getDataStore(APP_DATA_STORE_KEYS.USER_ROLE)
            .getObject<string>(APP_DATA_STORE_KEYS.USER_ROLE);
    }

    login(authResponse: AuthResponse) {
        this.loginStatusProviderService.isLoggedIn = true;
        this.loginStatusProviderService.authToken = authResponse.token;
        this.loginStatusProviderService.userRole = authResponse.urolecode;

        this.appDataStoreService.getDataStore(APP_DATA_STORE_KEYS.AUTH_TOKEN)
            .setObject<string>(APP_DATA_STORE_KEYS.AUTH_TOKEN, this.loginStatusProviderService.authToken);

        this.appDataStoreService.getDataStore(APP_DATA_STORE_KEYS.USER_ROLE)
            .setObject<string>(APP_DATA_STORE_KEYS.USER_ROLE, this.loginStatusProviderService.userRole);
        this.navigateToHome();
    }

    logout() {
        this.loginStatusProviderService.isLoggedIn = false;
        this.loginStatusProviderService.authToken = null;
        this.loginStatusProviderService.userRole = null;

        this.appDataStoreService.getDataStore(APP_DATA_STORE_KEYS.AUTH_TOKEN)
            .removeObject(APP_DATA_STORE_KEYS.AUTH_TOKEN);

        this.appDataStoreService.getDataStore(APP_DATA_STORE_KEYS.USER_ROLE)
            .removeObject(APP_DATA_STORE_KEYS.USER_ROLE);

        this.authService.logout().subscribe();
        this.router.navigate([ROUTE_LOGIN], { skipLocationChange: true });
    }

    validateAuthToken() {
        this.authService.validateAuthToken({ token: this.loginStatusProviderService.authToken }, false).subscribe(payloadResponse => {
            if (payloadResponse && payloadResponse.issuccess) {
                this.loginStatusProviderService.isLoggedIn = true;
                this.navigateToHome();
            } else {
                this.router.navigate([ROUTE_LOGIN], { skipLocationChange: true });
            }
        });
    }

    getUserLoginInfo() {
        this.userSharedService.getLoginInfo().subscribe(payloadResponse => {
            if (payloadResponse && payloadResponse.issuccess && payloadResponse.data) {
                const userInfo = new UserInfo();
                userInfo.copyFrom(payloadResponse.data);
                this.appDataStoreService.getDataStore(APP_DATA_STORE_KEYS.USER_INFO)
                    .setObject<UserInfo>(APP_DATA_STORE_KEYS.USER_INFO, userInfo);
                this.userInfoSubject.next(userInfo);
            }
        });
    }

    getCustomerLoginInfo() {
        if (AppSpecificDataProvider.userCateory === USER_CATEGORY.CU) {
            this.customerSharedService.getCusomerLoginInfo().subscribe(payloadResponse => {
                if (payloadResponse && payloadResponse.issuccess && payloadResponse.data) {
                    const customerInfo = new CustomerInfo();
                    customerInfo.copyFrom(payloadResponse.data);
                    this.appDataStoreService.getDataStore(APP_DATA_STORE_KEYS.USER_INFO)
                        .setObject<CustomerInfo>(APP_DATA_STORE_KEYS.USER_INFO, customerInfo);
                    this.customerInfoSubject.next(customerInfo);
                }
            });
        }
    }

    navigateToHome() {
        this.router.navigate([ROUTE_HOME], { skipLocationChange: true }).then(a => {
            this.getBasicInfoPostLogin();
        });
    }

    getBasicInfoPostLogin() {
        this.getUserLoginInfo();
        this.getCustomerLoginInfo();
    }
}
