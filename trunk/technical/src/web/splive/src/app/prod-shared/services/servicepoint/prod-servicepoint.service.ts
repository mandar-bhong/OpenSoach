import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

import { EnvironmentProvider } from '../../../shared/environment-provider';
import { RecordIDRequest, RecordIDResponse } from '../../../shared/models/api/common-models';
import { DataListRequest, DataListResponse } from '../../../shared/models/api/data-list-models';
import { PayloadResponse } from '../../../shared/models/api/payload-models';
import { ServerApiInterfaceService } from '../../../shared/services/api/server-api-interface.service';
import { ListingService } from '../../../shared/services/listing.service';
import { EnumNumberDatasource } from '../../../shared/utility/enum-number-datasource';
import { EnumDataSourceItem } from '../../../shared/models/ui/enum-datasource-item';
import { SERVICEPOINT_STATE } from '../../../shared/app-common-constants';
import {
  SrevicepointFilterRequest,
  ServicepointDataListResponse, ServicepointAssociateRequest, AssociateServicePointDeviceRequest, ServicepointListResponse, SPCategoriesShortDataResponse
} from '../../models/api/servicepoint-models';
import { OperatorServicepointListResponse } from '../../models/api/operator-models';
@Injectable({
  providedIn: 'root'
})
export class ProdServicepointService extends ListingService<SrevicepointFilterRequest, ServicepointDataListResponse> {

  constructor(private serverApiInterfaceService: ServerApiInterfaceService) {
    super();
  }
  getDataList(dataListRequest: DataListRequest<SrevicepointFilterRequest>, implicitErrorHandling = true):
    Observable<PayloadResponse<DataListResponse<ServicepointDataListResponse>>> {
    return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/servicepoint/list',
      dataListRequest, implicitErrorHandling);
  }
  getServicepointList(implicitErrorHandling = true):
    Observable<PayloadResponse<ServicepointListResponse[]>> {
    return this.serverApiInterfaceService.get(EnvironmentProvider.appbaseurl + '/api/v1/servicepoint/list/short',
      implicitErrorHandling);
  }
  getServicepointStates(): EnumDataSourceItem<number>[] {
    return EnumNumberDatasource.getDataSource('SERVICEPOINT_STATE_', SERVICEPOINT_STATE);
  }

  getServicepointState(state: number) {
    return 'SERVICEPOINT_STATE_' + state;
  }

  getCategoriesShortDataList(implicitErrorHandling = true):
    Observable<PayloadResponse<SPCategoriesShortDataResponse[]>> {
    return this.serverApiInterfaceService.get(EnvironmentProvider.appbaseurl + '/api/v1/servicepoint/category/list/short',
      implicitErrorHandling);
  }

  associateConfigure(request: ServicepointAssociateRequest, implicitErrorHandling = true):
    Observable<PayloadResponse<RecordIDResponse>> {
    return this.serverApiInterfaceService.post(EnvironmentProvider.appbaseurl + '/api/v1/service/instance/add',
      request, implicitErrorHandling);
  }
  associateDeviceServicePoint(request: AssociateServicePointDeviceRequest, implicitErrorHandling = true):
    Observable<PayloadResponse<null>> {
    return this.serverApiInterfaceService.post(EnvironmentProvider.appbaseurl + '/api/v1/servicepoint/associate/device',
      request, implicitErrorHandling);
  }
}
