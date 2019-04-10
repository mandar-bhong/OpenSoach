import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

import { EnvironmentProvider } from '../../../shared/environment-provider';
import { RecordIDRequest, RecordIDResponse } from '../../../shared/models/api/common-models';
import { DataListRequest, DataListResponse } from '../../../shared/models/api/data-list-models';
import { PayloadResponse } from '../../../shared/models/api/payload-models';
import { ServerApiInterfaceService } from '../../../shared/services/api/server-api-interface.service';
import { ListingService } from '../../../shared/services/listing.service';
import {
  DeviceAddDetailsRequest,
  DeviceAddRequest,
  DeviceAssociateProductListItemResponse,
  DeviceAssociateProductRequest,
  DeviceDataListResponse,
  DeviceDetailsResponse,
  DeviceFilterRequest,
  DeviceMasterUpdateResponse,
  DeviceMasterUpdateRequest,
} from '../models/api/device-models';

@Injectable()
export class DeviceService extends ListingService<DeviceFilterRequest, DeviceDataListResponse> {
  constructor(private serverApiInterfaceService: ServerApiInterfaceService) {
    super();
  }
  getDataList(dataListRequest: DataListRequest<DeviceFilterRequest>, implicitErrorHandling = true):
    Observable<PayloadResponse<DataListResponse<DeviceDataListResponse>>> {
    return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.baseurl + '/api/osu/v1/device/list',
      dataListRequest, implicitErrorHandling);
  }

  addDevice(deviceAddRequest: DeviceAddRequest, implicitErrorHandling = true):
    Observable<PayloadResponse<RecordIDResponse>> {
    return this.serverApiInterfaceService.post(EnvironmentProvider.baseurl + '/api/osu/v1/device/add',
      deviceAddRequest, implicitErrorHandling);
  }

  getDeviceList(dataListFilter: DataListRequest<DeviceFilterRequest>, implicitErrorHandling = true):
    Observable<PayloadResponse<DataListResponse<DeviceDataListResponse>>> {
    return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.baseurl + '/api/osu/v1/device/list',
      dataListFilter, implicitErrorHandling);
  }

  getDeviceDetails(request: RecordIDRequest, implicitErrorHandling = true):
    Observable<PayloadResponse<DeviceDetailsResponse>> {
    return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.baseurl + '/api/osu/v1/device/info/details',
      request, implicitErrorHandling);
  }

  updateDeviceDetails(deviceAddDetailsRequest: DeviceAddDetailsRequest, implicitErrorHandling = true):
    Observable<PayloadResponse<null>> {
    return this.serverApiInterfaceService.post(EnvironmentProvider.baseurl + '/api/osu/v1/device/update/details',
      deviceAddDetailsRequest, implicitErrorHandling);
  }

  associateDeviceToProduct(request: DeviceAssociateProductRequest, implicitErrorHandling = true):
    Observable<PayloadResponse<RecordIDResponse>> {
    return this.serverApiInterfaceService.post(EnvironmentProvider.baseurl + '/api/osu/v1/device/associate/customerproduct',
      request, implicitErrorHandling);
  }

  getDeviceProductAssociation(request: RecordIDRequest, implicitErrorHandling = true):
    Observable<PayloadResponse<DeviceAssociateProductListItemResponse[]>> {
    return this.serverApiInterfaceService.getWithQueryParams(
      EnvironmentProvider.baseurl + '/api/osu/v1/device/list/productassociation',
      request, implicitErrorHandling);
  }
  getDeviceMaster(request: RecordIDRequest, implicitErrorHandling = true):
    Observable<PayloadResponse<DeviceMasterUpdateResponse>> {
    return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.baseurl + '/api/osu/v1/device/info/master',
      request, implicitErrorHandling);
  }
  updateDeviceMaster(deviceMasterUpdateRequest: DeviceMasterUpdateRequest, implicitErrorHandling = true):
    Observable<PayloadResponse<null>> {
    return this.serverApiInterfaceService.post(EnvironmentProvider.baseurl + '/api/osu/v1/device/update/state',
      deviceMasterUpdateRequest, implicitErrorHandling);
  }
}
