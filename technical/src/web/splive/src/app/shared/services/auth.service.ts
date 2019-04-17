import { Injectable } from '@angular/core';
import { Observable, Subject } from 'rxjs';

import { EnvironmentProvider } from '../environment-provider';
import { AuthRequest, AuthResponse, ValidateAuthTokenRequest } from '../models/api/auth-models';
import { PayloadResponse } from '../models/api/payload-models';
import { ServerApiInterfaceService } from '../services/api/server-api-interface.service';
import { ChangeUserPasswordRequest, ActivationChangePassword } from '../models/api/user-models';
import { HttpClient } from 'selenium-webdriver/http';

@Injectable()
export class AuthService {
    code = new Subject<string>();
    userid: number;

    constructor(private serverApiInterfaceService: ServerApiInterfaceService) { }
    login(authRequest: AuthRequest, implicitErrorHandling = true): Observable<PayloadResponse<AuthResponse>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.baseurl + '/api/v1/login', authRequest, implicitErrorHandling);
    }

    logout(implicitErrorHandling = true): Observable<PayloadResponse<null>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.baseurl + '/api/v1/logout', null, implicitErrorHandling);
    }

    validateAuthToken(validateAuthTokenRequest: ValidateAuthTokenRequest,
        implicitErrorHandling = true): Observable<PayloadResponse<boolean>> {
        return this.serverApiInterfaceService.getWithQueryParams(
            EnvironmentProvider.baseurl + '/api/v1/validateauthtoken', validateAuthTokenRequest, implicitErrorHandling);
    }

    changeUserPassword(changeUserPasswordRequest: ChangeUserPasswordRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<any>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.baseurl + '/api/v1/user/create/password',
            changeUserPasswordRequest, implicitErrorHandling);
    }

    
    // changeUserPassword(changeUserPasswordRequest: ChangeUserPasswordRequest, implicitErrorHandling = true):
    //     Observable<PayloadResponse<any>> {
    //     return this.serverApiInterfaceService.post('http://172.105.232.148/api/v1/user/create/password',
    //         changeUserPasswordRequest, implicitErrorHandling);
    // }

    getActivationPerams(activationChangePassword: ActivationChangePassword, implicitErrorHandling = true):
        Observable<PayloadResponse<any>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.baseurl + '/api/v1/user/activation',
            activationChangePassword, implicitErrorHandling);
    }

    // getActivationPerams(activationChangePassword: ActivationChangePassword, implicitErrorHandling = true):
    //     Observable<PayloadResponse<any>> {
    //     return this.serverApiInterfaceService.post('http://172.105.232.148/api/v1/user/activation',
    //         activationChangePassword, implicitErrorHandling);
    // }
}
