import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

import { EnvironmentProvider } from '../../../shared/environment-provider';
import { PayloadResponse } from '../../../shared/models/api/payload-models';
import { ServerApiInterfaceService } from '../../../shared/services/api/server-api-interface.service';
import { ReportRequest, ReportResponse } from '../models/api/report-models';

@Injectable()
export class ReportService {
    constructor(private serverApiInterfaceService: ServerApiInterfaceService) { }

    getReportData(request: ReportRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<any>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/report/view',
            request, implicitErrorHandling);
    }
}
