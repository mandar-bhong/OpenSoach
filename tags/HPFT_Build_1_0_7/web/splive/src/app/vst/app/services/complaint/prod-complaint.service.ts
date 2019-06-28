import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';

import { COMPLAINT_STATE, SEVERIT_STATE } from '../../../../shared/app-common-constants';
import { EnvironmentProvider } from '../../../../shared/environment-provider';
import { RecordIDRequest } from '../../../../shared/models/api/common-models';
import { DataListRequest, DataListResponse } from '../../../../shared/models/api/data-list-models';
import { PayloadResponse } from '../../../../shared/models/api/payload-models';
import { EnumDataSourceItem } from '../../../../shared/models/ui/enum-datasource-item';
import { ServerApiInterfaceService } from '../../../../shared/services/api/server-api-interface.service';
import { ListingService } from '../../../../shared/services/listing.service';
import { EnumNumberDatasource } from '../../../../shared/utility/enum-number-datasource';
import {
  ComplaintDataListResponse,
  ComplaintDetailsResponse,
  ComplaintFiltrRequest,
  ComplaintUpdateRequest,
} from '../../models/api/complaint-models';


@Injectable()
export class ProdComplaintService extends ListingService<ComplaintFiltrRequest, ComplaintDataListResponse> {

  constructor(private serverApiInterfaceService: ServerApiInterfaceService) {
    super();
  }
  getDataList(dataListRequest: DataListRequest<ComplaintFiltrRequest>, implicitErrorHandling = true):
    Observable<PayloadResponse<DataListResponse<ComplaintDataListResponse>>> {
    return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/complaint/list',
      dataListRequest, implicitErrorHandling);
  }
  updateComplaintDetails(complaintUpdateRequest: ComplaintUpdateRequest, implicitErrorHandling = true):
    Observable<PayloadResponse<null>> {
    return this.serverApiInterfaceService.post(EnvironmentProvider.appbaseurl + '/api/v1/complaint/update',
      complaintUpdateRequest, implicitErrorHandling);
  }
  getComplaintDetails(request: RecordIDRequest, implicitErrorHandling = true):
    Observable<PayloadResponse<ComplaintDetailsResponse>> {
    return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/complaint/info',
      request, implicitErrorHandling);
  }
  getComplaintStates(): EnumDataSourceItem<number>[] {
    return EnumNumberDatasource.getDataSource('COMPLAINT_STATE_', COMPLAINT_STATE);
  }
  getComplaintState(states: number) {
    return 'COMPLAINT_STATE_' + states;
  }
  getSeveritiesStates(): EnumDataSourceItem<number>[] {
    return EnumNumberDatasource.getDataSource('SEVERIT_STATE_', SEVERIT_STATE);
  }
  getSeveritiesState(states: number) {
    return 'SEVERIT_STATE_' + states;
  }
}
