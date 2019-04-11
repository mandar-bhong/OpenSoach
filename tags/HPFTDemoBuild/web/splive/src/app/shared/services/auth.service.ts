import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

import { EnvironmentProvider } from '../environment-provider';
import { AuthRequest, AuthResponse, ValidateAuthTokenRequest } from '../models/api/auth-models';
import { PayloadResponse } from '../models/api/payload-models';
import { ServerApiInterfaceService } from '../services/api/server-api-interface.service';

@Injectable()
export class AuthService {

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
}
