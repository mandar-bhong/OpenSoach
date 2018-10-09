import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

import { EnvironmentProvider } from '../../../shared/environment-provider';
import { RecordIDRequest, RecordIDResponse } from '../../../shared/models/api/common-models';
import { DataListRequest, DataListResponse } from '../../../shared/models/api/data-list-models';
import { PayloadResponse } from '../../../shared/models/api/payload-models';
import { ServerApiInterfaceService } from '../../../shared/services/api/server-api-interface.service';
import { ListingService } from '../../../shared/services/listing.service';
import { JobFiltrRequest, JobDataListResponse, JobDetailsDataListResponse,
    StatusChangeRequest, VehicleDetailsResponse, JobDetailslistResponse } from '../models/api/job-models';
import { EnumDataSourceItem } from '../../../shared/models/ui/enum-datasource-item';
import { EnumNumberDatasource } from '../../../shared/utility/enum-number-datasource';
import { SNAPSHOT_STATE } from '../app-constants';


@Injectable()
export class JobService extends ListingService<JobFiltrRequest, JobDataListResponse> {
    constructor(private serverApiInterfaceService: ServerApiInterfaceService) {
        super();
    }
    getDataList(dataListRequest: DataListRequest<JobFiltrRequest>, implicitErrorHandling = true):
        Observable<PayloadResponse<DataListResponse<JobDataListResponse>>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/job/list',
            dataListRequest, implicitErrorHandling);
    }
    getJobStates(): EnumDataSourceItem<number>[] {
        return EnumNumberDatasource.getDataSource('SNAPSHOT_STATE_', SNAPSHOT_STATE);
    }

    getJobState(state: number) {
        return 'SNAPSHOT_STATE_' + state;
    }
    // getDataListDetails(implicitErrorHandling = true):
    //     Observable<PayloadResponse<JobDetailsDataListResponse[]>> {
    //     return this.serverApiInterfaceService.get(EnvironmentProvider.appbaseurl + '/api/v1/patient/list',
    //         implicitErrorHandling);
    // }
    updateStatus(statusChangeRequest: StatusChangeRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<null>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.appbaseurl + '/api/v1/job/update/state',
        statusChangeRequest, implicitErrorHandling);
    }
    getJobsDetails(request: RecordIDRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<VehicleDetailsResponse>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/vehicle/info',
            request, implicitErrorHandling);
    }
    getJobsDetailsList(request: RecordIDRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<JobDetailslistResponse[]>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/job/info',
            request, implicitErrorHandling);
    }
    // getDataListDetails(implicitErrorHandling = true):
    //     Observable<PayloadResponse<JobDetailslistResponse[]>> {
    //     return this.serverApiInterfaceService.get(EnvironmentProvider.appbaseurl + '/api/v1/patient/list',
    //         implicitErrorHandling);
    // }
}
