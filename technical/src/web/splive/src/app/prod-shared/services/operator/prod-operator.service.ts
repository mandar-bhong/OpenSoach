import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';

import { EnvironmentProvider } from '../../../shared/environment-provider';
import { RecordIDRequest, RecordIDResponse } from '../../../shared/models/api/common-models';
import { DataListRequest, DataListResponse } from '../../../shared/models/api/data-list-models';

import { PayloadResponse } from '../../../shared/models/api/payload-models';
import { ServerApiInterfaceService } from '../../../shared/services/api/server-api-interface.service';
import { ListingService } from '../../../shared/services/listing.service';
import {
  OperatorDataListResponse, OperatorFiltrRequest,
  OperatorAddRequest, OperatorUpdateRequest, OperatorDetailsResponse, OperatorAssociateListResponse, OperatorServicepointListResponse
} from '../../models/api/operator-models';
import { OPERATOR_STATE, OPERATOR_AREA } from '../../../shared/app-common-constants';
import { EnumDataSourceItem } from '../../../shared/models/ui/enum-datasource-item';
import { EnumNumberDatasource } from '../../../shared/utility/enum-number-datasource';

@Injectable()
export class ProdOperatorService extends ListingService<OperatorFiltrRequest, OperatorDataListResponse> {

  constructor(private serverApiInterfaceService: ServerApiInterfaceService) {
    super();
  }
  getDataList(dataListRequest: DataListRequest<OperatorFiltrRequest>, implicitErrorHandling = true):
    Observable<PayloadResponse<DataListResponse<OperatorDataListResponse>>> {
    return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/fieldoperator/list',
      dataListRequest, implicitErrorHandling);
  }
  addOperator(operatorAddRequest: OperatorAddRequest, implicitErrorHandling = true):
    Observable<PayloadResponse<RecordIDResponse>> {
    return this.serverApiInterfaceService.post(EnvironmentProvider.appbaseurl + '/api/v1/fieldoperator/add',
      operatorAddRequest, implicitErrorHandling);
  }
  updateOperatorDetails(operatorUpdateRequest: OperatorUpdateRequest, implicitErrorHandling = true):
    Observable<PayloadResponse<null>> {
    return this.serverApiInterfaceService.post(EnvironmentProvider.appbaseurl + '/api/v1/fieldoperator/update',
      operatorUpdateRequest, implicitErrorHandling);
  }
  getOperatorDetails(request: RecordIDRequest, implicitErrorHandling = true):
    Observable<PayloadResponse<OperatorDetailsResponse>> {
    return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/fieldoperator/info',
      request, implicitErrorHandling);
  }
  getOperatorServicpoint(request: RecordIDRequest, implicitErrorHandling = true):
    Observable<PayloadResponse<OperatorAssociateListResponse[]>> {
    return this.serverApiInterfaceService.getWithQueryParams
      (EnvironmentProvider.appbaseurl + '/api/v1/servicepoint/associate/fieldoperator/info',
      request, implicitErrorHandling);
  }
  getServicepointList(implicitErrorHandling = true):
    Observable<PayloadResponse<OperatorServicepointListResponse[]>> {
    return this.serverApiInterfaceService.get(EnvironmentProvider.appbaseurl + '/api/v1/servicepoint/list/short',
      implicitErrorHandling);
  }
  getOperatorStates(): EnumDataSourceItem<number>[] {
    return EnumNumberDatasource.getDataSource('OPERATOR_STATE_', OPERATOR_STATE);
  }
  getOperatorState(states: number) {
    return 'OPERATOR_STATE_' + states;
  }
  getOperatorAreas(): EnumDataSourceItem<number>[] {
    return EnumNumberDatasource.getDataSource('OPERATOR_AREA_', OPERATOR_AREA);
  }
  getOperatorArea(areas: number) {
    return 'OPERATOR_AREA_' + areas;
  }


}
