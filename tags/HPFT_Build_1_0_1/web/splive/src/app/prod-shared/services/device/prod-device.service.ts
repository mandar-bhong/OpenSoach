import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

import { CONNECTION_STATE } from '../../../shared/app-common-constants';
import { EnvironmentProvider } from '../../../shared/environment-provider';
import { RecordIDRequest } from '../../../shared/models/api/common-models';
import { DataListRequest, DataListResponse } from '../../../shared/models/api/data-list-models';
import { PayloadResponse } from '../../../shared/models/api/payload-models';
import { EnumDataSourceItem } from '../../../shared/models/ui/enum-datasource-item';
import { ServerApiInterfaceService } from '../../../shared/services/api/server-api-interface.service';
import { ListingService } from '../../../shared/services/listing.service';
import { EnumNumberDatasource } from '../../../shared/utility/enum-number-datasource';
import {
    DeviceDataListResponse,
    DeviceDetailsResponse,
    DeviceDetailsUpdateRequest,
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
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/device/list',
            dataListRequest, implicitErrorHandling);
    }

    getDeviceList(implicitErrorHandling = true):
        Observable<PayloadResponse<DeviceListItemResponse[]>> {
        return this.serverApiInterfaceService.get(EnvironmentProvider.appbaseurl + '/api/v1/device/list/short',
            implicitErrorHandling);
    }
    getDevicesNotAssociatedWithSP(implicitErrorHandling = true):
        Observable<PayloadResponse<DeviceListItemResponse[]>> {
        return this.serverApiInterfaceService.get(EnvironmentProvider.appbaseurl + '/api/v1/device/list/short/noassociation',
            implicitErrorHandling);
    }
    getDeviceDetails(request: RecordIDRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<DeviceDetailsResponse>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/device/info',
            request, implicitErrorHandling);
    }
    updateDeviceDetails(deviceDetailsUpdateRequest: DeviceDetailsUpdateRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<null>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.appbaseurl + '/api/v1/device/update',
            deviceDetailsUpdateRequest, implicitErrorHandling);
    }
    getConnectionStates(): EnumDataSourceItem<number>[] {
        return EnumNumberDatasource.getDataSource('CONNECTION_STATE_', CONNECTION_STATE);
    }
    getConnectionState(state: number) {
        return 'CONNECTION_STATE_' + state;
    }
}
