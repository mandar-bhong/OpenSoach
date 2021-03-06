import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

import {
    ServiceConfigurationRequest,
    ServiceConfigurationResponse,
    ServiceConfigurationUpdateRequest,
    ServiceConfigureDataListResponse,
    ServiceConfigureFilterRequest,
    ServicepointConfigureListResponse,
    ServicepointConfigureTemplateListRequest,
    ServicePointWithConfigurationResponse,
} from '../../../prod-shared/models/api/service-configuration-models';
import { EnvironmentProvider } from '../../../shared/environment-provider';
import { RecordIDRequest, RecordIDResponse } from '../../../shared/models/api/common-models';
import { DataListRequest, DataListResponse } from '../../../shared/models/api/data-list-models';
import { PayloadResponse } from '../../../shared/models/api/payload-models';
import { ServerApiInterfaceService } from '../../../shared/services/api/server-api-interface.service';
import { ListingService } from '../../../shared/services/listing.service';

@Injectable()
export class SpServiceConfService extends ListingService<ServiceConfigureFilterRequest, ServiceConfigureDataListResponse> {
    constructor(
        private serverApiInterfaceService: ServerApiInterfaceService) {
        super();
    }
    getDataList(dataListRequest: DataListRequest<ServiceConfigureFilterRequest>, implicitErrorHandling = true):
        Observable<PayloadResponse<DataListResponse<ServiceConfigureDataListResponse>>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/service/config/list',
            dataListRequest, implicitErrorHandling);
    }
    addServiceConf(serviceConfigurationRequest: ServiceConfigurationRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<RecordIDResponse>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.appbaseurl + '/api/v1/service/config/add',
            serviceConfigurationRequest, implicitErrorHandling);
    }
    updateServiceConf(serviceConfigurationUpdateRequest: ServiceConfigurationUpdateRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<null>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.appbaseurl + '/api/v1/service/config/update',
            serviceConfigurationUpdateRequest, implicitErrorHandling);
    }
    getServiceConf(request: RecordIDRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<ServiceConfigurationResponse>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/service/config/info',
            request, implicitErrorHandling);
    }
    getServicepointConfigureList(implicitErrorHandling = true):
        Observable<PayloadResponse<ServicepointConfigureListResponse[]>> {
        return this.serverApiInterfaceService.get(EnvironmentProvider.appbaseurl + '/api/v1/service/config/list/short',
            implicitErrorHandling);
    }
    copyTemplateList(request: ServicepointConfigureTemplateListRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<RecordIDResponse>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.appbaseurl + '/api/v1/service/config/copytemplate',
            request, implicitErrorHandling);
    }

    getServicePointsWithConfigurations(implicitErrorHandling = true):
        Observable<PayloadResponse<ServicePointWithConfigurationResponse[]>> {
        return this.serverApiInterfaceService.get(EnvironmentProvider.appbaseurl + '/api/v1/service/spconfig/list/short',
            implicitErrorHandling);
    }

}

