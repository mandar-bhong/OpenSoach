import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';

import { EnvironmentProvider } from '../../environment-provider';
import { PayloadResponse } from '../../models/api/payload-models';
import { UserLoginInfoResponse } from '../../models/api/user-models';
import { ServerApiInterfaceService } from '../../services/api/server-api-interface.service';

@Injectable()
export class UserSharedService {

    constructor(private serverApiInterfaceService: ServerApiInterfaceService) { }
    getLoginInfo(implicitErrorHandling = true): Observable<PayloadResponse<UserLoginInfoResponse>> {
        return this.serverApiInterfaceService.get(EnvironmentProvider.baseurl + '/api/v1/user/info/login', implicitErrorHandling);
    }
}
