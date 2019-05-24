import { Injectable } from '@angular/core';
import { ScheduleFilter } from 'app/models/api/schedule-request';
import { ScheduleDataResponse } from 'app/models/api/schedule-response';
import { Observable } from 'rxjs';
import { EnvironmentProvider } from '../../../../shared/environment-provider';
import { DataListRequest, DataListResponse } from '../../../../shared/models/api/data-list-models';
import { PayloadResponse } from '../../../../shared/models/api/payload-models';
import { ServerApiInterfaceService } from '../../../../shared/services/api/server-api-interface.service';
import { ListingService } from '../../../../shared/services/listing.service';
import { ServicepointConfigureListResponse } from '../../../../prod-shared/models/api/service-configuration-models';
import { ServicepointListResponse } from '../../../../prod-shared/models/api/servicepoint-models';

@Injectable()
export class ScheduleService extends ListingService<ScheduleFilter, ScheduleDataResponse<string>[]>{
  patientid: number;
  admissionid: number;
  constructor(private serverApiInterfaceService: ServerApiInterfaceService) {
    super();
  }

  getDataList(dataListRequest: DataListRequest<ScheduleFilter>, implicitErrorHandling = true):
    Observable<PayloadResponse<DataListResponse<ScheduleDataResponse<string>[]>>> {
    return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/patient/list/patientconf',
      dataListRequest, implicitErrorHandling);
  }
}
