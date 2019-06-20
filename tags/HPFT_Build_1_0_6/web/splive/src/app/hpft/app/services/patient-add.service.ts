import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

import { EnvironmentProvider } from '../../../shared/environment-provider';
import { PayloadResponse } from '../../../shared/models/api/payload-models';
import { ServerApiInterfaceService } from '../../../shared/services/api/server-api-interface.service';

import { ListingService } from '../../../shared/services/listing.service';
import { DataListRequest, DataListResponse } from '../../../shared/models/api/data-list-models';
import { ServicepointListResponse } from '../../../prod-shared/models/api/servicepoint-models';
import { ServicepointConfigureListResponse } from '../../../prod-shared/models/api/service-configuration-models';
import { PatientSearchRequestFilter, PatientSearchResponseFilter } from '../models/api/patient-data-models';

@Injectable()

export class PatientAddService extends ListingService<PatientSearchRequestFilter, PatientSearchResponseFilter>{
  patientid: number;
  constructor(private serverApiInterfaceService: ServerApiInterfaceService) {
    super();
  }
  getDataList(dataListRequest: DataListRequest<PatientSearchRequestFilter>, implicitErrorHandling = true):
    Observable<PayloadResponse<DataListResponse<PatientSearchResponseFilter>>> {
    return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/patient/list/master',
      dataListRequest, implicitErrorHandling);
  }
  // getPatientDataList(dataListRequest: DataListRequest<PatientSearchRequestFilter>, implicitErrorHandling = true):
  //   Observable<PayloadResponse<DataListResponse<PatientSearchResponseFilter>>> {
  //   return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/patient/list/master',
  //     dataListRequest, implicitErrorHandling);
  // }
  getServicepointList(implicitErrorHandling = true):
    Observable<PayloadResponse<ServicepointListResponse[]>> {
    return this.serverApiInterfaceService.get(EnvironmentProvider.appbaseurl + '/api/v1/servicepoint/list/short',
      implicitErrorHandling);
  }
  getServicepointConfigureList(implicitErrorHandling = true):
    Observable<PayloadResponse<ServicepointConfigureListResponse[]>> {
    return this.serverApiInterfaceService.get(EnvironmentProvider.appbaseurl + '/api/v1/service/config/list/short',
      implicitErrorHandling);
  }
}
