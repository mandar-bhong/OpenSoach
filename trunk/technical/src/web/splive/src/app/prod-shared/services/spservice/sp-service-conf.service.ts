import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';

import {
    ServiceConfigurationlistResponse,
    ServiceConfigurationRequest,
    ServiceConfigurationUpdateRequest,
} from '../../../prod-shared/models/api/service-configuration-models';
import { EnvironmentProvider } from '../../../shared/environment-provider';
import { RecordIDRequest, RecordIDResponse } from '../../../shared/models/api/common-models';
import { PayloadResponse } from '../../../shared/models/api/payload-models';
import { ServerApiInterfaceService } from '../../../shared/services/api/server-api-interface.service';

@Injectable()
export class SpServiceConfService {
    constructor(
        private serverApiInterfaceService: ServerApiInterfaceService) { }
    addServiceConf(serviceConfigurationRequest: ServiceConfigurationRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<RecordIDResponse>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.appbaseurl + '/api/v1/service/config/add',
            serviceConfigurationRequest, implicitErrorHandling);
    }
    updateServiceConf(serviceConfigurationUpdateRequest: ServiceConfigurationUpdateRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<null>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.baseurl + '/api/v1/service/config/update',
            serviceConfigurationUpdateRequest, implicitErrorHandling);
    }
    getServiceConf(request: RecordIDRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<ServiceConfigurationlistResponse[]>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/service/config/list',
            request, implicitErrorHandling);
    }
}

