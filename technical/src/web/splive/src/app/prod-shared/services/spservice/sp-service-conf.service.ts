import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

import {
    AssociateServicePointDeviceRequest,
    ServiceConfigurationlistResponse,
    ServiceConfigurationRequest,
    ServiceConfigurationUpdateRequest,
    ServicepointAssociateRequest,
    ServicepointConfigureListResponse,
    ServicepointConfigureTemplateListRequest,
    ServicepointDataListResponse,
    SrevicepointFilterRequest,
} from '../../../prod-shared/models/api/service-configuration-models';
import { SERVICEPOINT_STATE } from '../../../shared/app-common-constants';
import { EnvironmentProvider } from '../../../shared/environment-provider';
import { RecordIDRequest, RecordIDResponse } from '../../../shared/models/api/common-models';
import { DataListRequest, DataListResponse } from '../../../shared/models/api/data-list-models';
import { PayloadResponse } from '../../../shared/models/api/payload-models';
import { EnumDataSourceItem } from '../../../shared/models/ui/enum-datasource-item';
import { ServerApiInterfaceService } from '../../../shared/services/api/server-api-interface.service';
import { ListingService } from '../../../shared/services/listing.service';
import { EnumNumberDatasource } from '../../../shared/utility/enum-number-datasource';



@Injectable()
export class SpServiceConfService extends ListingService<SrevicepointFilterRequest, ServicepointDataListResponse> {
    constructor(
        private serverApiInterfaceService: ServerApiInterfaceService) {
        super();
    }
    getDataList(dataListRequest: DataListRequest<SrevicepointFilterRequest>, implicitErrorHandling = true):
        Observable<PayloadResponse<DataListResponse<ServicepointDataListResponse>>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/servicepoint/list',
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
        Observable<PayloadResponse<ServiceConfigurationlistResponse[]>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/service/config/list',
            request, implicitErrorHandling);
    }
    getServicepointStates(): EnumDataSourceItem<number>[] {
        return EnumNumberDatasource.getDataSource('SERVICEPOINT_STATE_', SERVICEPOINT_STATE);
    }

    getServicepointState(state: number) {
        return 'SERVICEPOINT_STATE_' + state;
    }
    associateConfigure(request: ServicepointAssociateRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<RecordIDResponse>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.appbaseurl + '/api/v1/service/instance/add',
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

    associateDeviceServicePoint(request: AssociateServicePointDeviceRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<null>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.appbaseurl + '/api/v1/servicepoint/associate/device',
            request, implicitErrorHandling);
    }
}

