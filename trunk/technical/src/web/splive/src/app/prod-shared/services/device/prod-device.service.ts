import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';

import { EnvironmentProvider } from '../../../shared/environment-provider';
import { RecordIDRequest } from '../../../shared/models/api/common-models';
import { DataListRequest, DataListResponse } from '../../../shared/models/api/data-list-models';
import { PayloadResponse } from '../../../shared/models/api/payload-models';
import { ServerApiInterfaceService } from '../../../shared/services/api/server-api-interface.service';
import { ListingService } from '../../../shared/services/listing.service';
import {
    DeviceDataListResponse,
    DeviceDetailsResponse,
    DeviceFilterRequest,
    DeviceListItemResponse,
} from '../../models/api/device-models';

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

    getDeviceList(implicitErrorHandling = true):
        Observable<PayloadResponse<DeviceListItemResponse[]>> {
        return this.serverApiInterfaceService.get(EnvironmentProvider.appbaseurl + '/api/v1/device/list/short',
            implicitErrorHandling);
    }

    getDeviceDetails(request: RecordIDRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<DeviceDetailsResponse>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.baseurl + '/api/cu/v1/device/info/details',
            request, implicitErrorHandling);
    }
}
