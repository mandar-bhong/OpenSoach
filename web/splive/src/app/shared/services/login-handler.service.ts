import { Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { Subject } from 'rxjs';
import { HPFTRouteHelper } from "../../hpft/app/helpers/route-helper";
import { APP_SHARED_DATA_STORE_KEYS, PROD_HPFT, ROUTE_LOGIN } from '../app-common-constants';
import { AppRepoShared } from '../app-repo/app-repo';
import { AuthResponse } from '../models/api/auth-models';
import { CustomerInfo } from '../models/ui/customer-models';
import { UserInfo } from '../models/ui/user-models';
import { AppDataStoreService } from './app-data-store/app-data-store-service';
import { AuthService } from './auth.service';
import { CustomerSharedService } from './customer/customer-shared.service';
import { LoginStatusProviderService } from './login-status-provider.service';
import { AppUserService } from './user/app-user.service';


@Injectable()
export class LoginHandlerService {
    userInfoSubject = new Subject<UserInfo>();
    customerInfoSubject = new Subject<CustomerInfo>();
    userHomeRoute: any;
    constructor(private appDataStoreService: AppDataStoreService,
        private router: Router,
        private authService: AuthService,
        private loginStatusProviderService: LoginStatusProviderService,
        private userSharedService: AppUserService,
        private customerSharedService: CustomerSharedService) {
        this.loginStatusProviderService.logginStatusChanged.subscribe(status => {
            this.logout();
        });
    }

    init() {
        switch (AppRepoShared.appProductCode) {
            case PROD_HPFT:
                this.userHomeRoute = HPFTRouteHelper.getUserHomeRoute;
                break;
            default:
                this.userHomeRoute = this.userHomeRoutHandler;
                break;
        }

        this.loginStatusProviderService.authToken = this.appDataStoreService.getDataStore(APP_SHARED_DATA_STORE_KEYS.AUTH_TOKEN)
            .getObject<string>(APP_SHARED_DATA_STORE_KEYS.AUTH_TOKEN);
        if (this.loginStatusProviderService.authToken) {
            this.validateAuthToken();
        } else {
            this.router.navigate([ROUTE_LOGIN], { skipLocationChange: true });
        }

        this.loginStatusProviderService.userRole = this.appDataStoreService.getDataStore(APP_SHARED_DATA_STORE_KEYS.USER_ROLE)
            .getObject<string>(APP_SHARED_DATA_STORE_KEYS.USER_ROLE);
    }

    login(authResponse: AuthResponse) {
        this.loginStatusProviderService.isLoggedIn = true;
        this.loginStatusProviderService.authToken = authResponse.token;
        this.loginStatusProviderService.userRole = authResponse.urolecode;

        this.appDataStoreService.getDataStore(APP_SHARED_DATA_STORE_KEYS.AUTH_TOKEN)
            .setObject<string>(APP_SHARED_DATA_STORE_KEYS.AUTH_TOKEN, this.loginStatusProviderService.authToken);

        this.appDataStoreService.getDataStore(APP_SHARED_DATA_STORE_KEYS.USER_ROLE)
            .setObject<string>(APP_SHARED_DATA_STORE_KEYS.USER_ROLE, this.loginStatusProviderService.userRole);
        this.navigateToHome();
    }

    logout(navigateToLogin: boolean = true) {
        this.loginStatusProviderService.isLoggedIn = false;
        this.loginStatusProviderService.authToken = null;
        this.loginStatusProviderService.userRole = null;

        this.appDataStoreService.getDataStore(APP_SHARED_DATA_STORE_KEYS.AUTH_TOKEN)
            .removeObject(APP_SHARED_DATA_STORE_KEYS.AUTH_TOKEN);

        this.appDataStoreService.getDataStore(APP_SHARED_DATA_STORE_KEYS.USER_ROLE)
            .removeObject(APP_SHARED_DATA_STORE_KEYS.USER_ROLE);

        this.authService.logout().subscribe();
        if (navigateToLogin) {
            this.router.navigate([ROUTE_LOGIN], { skipLocationChange: true });
        }
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
                this.appDataStoreService.getDataStore(APP_SHARED_DATA_STORE_KEYS.USER_INFO)
                    .setObject<UserInfo>(APP_SHARED_DATA_STORE_KEYS.USER_INFO, userInfo);
                this.userInfoSubject.next(userInfo);
            }
        });
    }

    // TODO: Need to move this method to prod_shared
    // getCustomerLoginInfo() {
    //     if (AppSpecificDataProvider.userCateory === USER_CATEGORY.CU) {
    //         this.customerSharedService.getCusomerLoginInfo().subscribe(payloadResponse => {
    //             if (payloadResponse && payloadResponse.issuccess && payloadResponse.data) {
    //                 const customerInfo = new CustomerInfo();
    //                 customerInfo.copyFrom(payloadResponse.data);
    //                 this.appDataStoreService.getDataStore(APP_SHARED_DATA_STORE_KEYS.CUSTOMER_INFO)
    //                     .setObject<CustomerInfo>(APP_SHARED_DATA_STORE_KEYS.CUSTOMER_INFO, customerInfo);
    //                 this.customerInfoSubject.next(customerInfo);
    //             }
    //         });
    //     }
    // }

    navigateToHome() {
        this.router.navigate([this.userHomeRoute(this.loginStatusProviderService.userRole)], { skipLocationChange: true }).then(a => {
            this.getBasicInfoPostLogin();
        });
    }

    userHomeRoutHandler(userrole: string) {
        this.router.navigate([''], { skipLocationChange: true });
    }

    getBasicInfoPostLogin() {
        this.getUserLoginInfo();
    }
}
