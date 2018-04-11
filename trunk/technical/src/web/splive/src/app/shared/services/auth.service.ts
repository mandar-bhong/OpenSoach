import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';

import { EnvironmentProvider } from '../environment-provider';
import { AuthRequest, AuthResponse } from '../models/api/auth-models';
import { PayloadResponse } from '../models/api/payload-models';
import { ServerApiInterfaceService } from '../services/api/server-api-interface.service';

@Injectable()
export class AuthService {

    constructor(private serverApiInterfaceService: ServerApiInterfaceService) { }
    login(authRequest: AuthRequest, implicitErrorHandling = true): Observable<PayloadResponse<AuthResponse>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.baseurl + '/api/v1/login', authRequest, implicitErrorHandling);
    }
}
