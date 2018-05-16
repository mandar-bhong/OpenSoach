import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';

import { DEVICE_STATE } from '../app-common-constants';
import { EnvironmentProvider } from '../environment-provider';
import { RecordIDRequest } from '../models/api/common-models';
import { DataListRequest, DataListResponse } from '../models/api/data-list-models';
import { DeviceDataListResponse, DeviceDetailsResponse, DeviceFilterRequest } from '../models/api/device-models';
import { PayloadResponse } from '../models/api/payload-models';
import { EnumDataSourceItem } from '../models/ui/enum-datasource-item';
import { EnumNumberDatasource } from '../utility/enum-number-datasource';
import { ServerApiInterfaceService } from './api/server-api-interface.service';
import { ListingService } from './listing.service';

@Injectable()
export class DeviceSharedService extends ListingService<DeviceFilterRequest, DeviceDataListResponse> {
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

    getDeviceStates(): EnumDataSourceItem<number>[] {
        return EnumNumberDatasource.getDataSource('DEVICE_STATE_', DEVICE_STATE);
    }

    getDeviceState(state: number) {
        return 'DEVICE_STATE_' + state;
    }
}
