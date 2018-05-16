import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';

import { EnvironmentProvider } from '../../../shared/environment-provider';
import { RecordIDRequest } from '../../../shared/models/api/common-models';
import { DataListRequest, DataListResponse } from '../../../shared/models/api/data-list-models';
import { DeviceDataListResponse, DeviceDetailsResponse, DeviceFilterRequest } from '../../models/api/device-models';
import { PayloadResponse } from '../../../shared/models/api/payload-models';
import { ServerApiInterfaceService } from '../../../shared/services/api/server-api-interface.service';
import { ListingService } from '../../../shared/services/listing.service';

@Injectable()
export class ProdDeviceService extends ListingService<DeviceFilterRequest, DeviceDataListResponse> {
    constructor(private serverApiInterfaceService: ServerApiInterfaceService) {
        super();
    }
    getDataList(dataListRequest: DataListRequest<DeviceFilterRequest>, implicitErrorHandling = true):
        Observable<PayloadResponse<DataListResponse<DeviceDataListResponse>>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.baseurl + '/api/cu/v1/device/list',
            dataListRequest, implicitErrorHandling);
    }

    getDeviceList(dataListFilter: DataListRequest<DeviceFilterRequest>, implicitErrorHandling = true):
        Observable<PayloadResponse<DataListResponse<DeviceDataListResponse>>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.baseurl + '/api/cu/v1/device/list',
            dataListFilter, implicitErrorHandling);
    }

    getDeviceDetails(request: RecordIDRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<DeviceDetailsResponse>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.baseurl + '/api/cu/v1/device/info/details',
            request, implicitErrorHandling);
    }
}
