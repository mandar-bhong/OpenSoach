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
import { ReportRequestParams } from '../models/api/report-models';
import { SaveFileService } from '../../../shared/services/save-file.service';


@Injectable()
export class JobService extends ListingService<JobFiltrRequest, JobDataListResponse> {
    constructor(private serverApiInterfaceService: ServerApiInterfaceService,
        private saveFileService: SaveFileService) {
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
    updateStatus(statusChangeRequest: StatusChangeRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<null>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.appbaseurl + '/api/v1/job/update/state',
        statusChangeRequest, implicitErrorHandling);
    }
    getJobsDetails(request: RecordIDRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<VehicleDetailsResponse>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/job/vehicleinfo',
            request, implicitErrorHandling);
    }
    getJobsDetailsList(request: RecordIDRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<JobDetailslistResponse[]>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/job/info',
            request, implicitErrorHandling);
    }
    generateReport(request: ReportRequestParams, implicitErrorHandling = true):
        Observable<Blob> {
        return this.serverApiInterfaceService.downloadFile(EnvironmentProvider.appbaseurl + '/api/v1/report/generate',
            request, implicitErrorHandling);
    }
    saveReport(data: Blob, filename: string) {
        this.saveFileService.saveFile(data, filename);
    }
}
