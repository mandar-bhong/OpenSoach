import { Injectable } from '@angular/core';
import { PatientFilterRequest, PatientFilterResponse, LaboratoryReportAddRequest, HospitalListResponse, HospitalSearchRequest } from 'app/models/api/hospital-models';
import { DataListRequest, DataListResponse } from '../../../shared/models/api/data-list-models';
import { Observable } from 'rxjs';
import { PayloadResponse } from '../../../shared/models/api/payload-models';
import { EnvironmentProvider } from '../../../shared/environment-provider';
import { ServerApiInterfaceService } from '../../../shared/services/api/server-api-interface.service';
import { ListingService } from '../../../shared/services/listing.service';


@Injectable()
export class HospitalService extends ListingService<PatientFilterRequest, PatientFilterResponse> {

  cpmid: number;
  admissionid: number;
  selcetdIndex: number;
  hospitalName: string;
  constructor(private serverApiInterfaceService: ServerApiInterfaceService) {
    super();
  }

  getDataList(dataListRequest: DataListRequest<PatientFilterRequest>, implicitErrorHandling = true):
    Observable<PayloadResponse<DataListResponse<PatientFilterResponse>>> {
    return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/patient/list',
      dataListRequest, implicitErrorHandling);
  }
  //Post method for docuemnt upload
  addReportData(request: LaboratoryReportAddRequest, implicitErrorHandling = true):
    Observable<PayloadResponse<any>> {
    return this.serverApiInterfaceService.post(EnvironmentProvider.appbaseurl + '/api/v1/patient/pathologyrecord/add',
      request, implicitErrorHandling);
  }

  //Post method for docuemnt upload
  uploadReportsDocuments(formData: any, implicitErrorHandling = true):
    Observable<PayloadResponse<any>> {
    return this.serverApiInterfaceService.upload(EnvironmentProvider.appbaseurl + '/api/v1/document/upload',
      formData, implicitErrorHandling);
  }
  // Display list for hospital

  getServicepointList(request: HospitalSearchRequest, implicitErrorHandling = true):
    Observable<PayloadResponse<HospitalListResponse[]>> {
    return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.baseurl + '/api/v1/customer/cpm/list/short',
      request, implicitErrorHandling);
  }

}
