import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

import { EnvironmentProvider } from '../../shared/environment-provider';
import { PayloadResponse } from '../../shared/models/api/payload-models';
import { ServerApiInterfaceService } from '../../shared/services/api/server-api-interface.service';
import { SplBaseURLResponse } from '../models/api/spl-conf-models';

@Injectable()
export class SplConfService {
    constructor(
        private serverApiInterfaceService: ServerApiInterfaceService) { }

    getSplBaseUrl(implicitErrorHandling = true):
        Observable<PayloadResponse<SplBaseURLResponse>> {
        return this.serverApiInterfaceService.get(EnvironmentProvider.appbaseurl + '/api/v1/splprod/baseurl',
            implicitErrorHandling);
    }
}

