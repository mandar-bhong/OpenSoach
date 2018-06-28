import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

import { EnvironmentProvider } from '../../../shared/environment-provider';
import { PayloadResponse } from '../../../shared/models/api/payload-models';
import { ServerApiInterfaceService } from '../../../shared/services/api/server-api-interface.service';
import { SaveFileService } from '../../../shared/services/save-file.service';
import { ReportRequestParams, ReportResponse } from '../models/api/report-models';

@Injectable()
export class ReportService {
    constructor(private serverApiInterfaceService: ServerApiInterfaceService,
        private saveFileService: SaveFileService) { }

    getReportData(request: ReportRequestParams, implicitErrorHandling = true):
        Observable<PayloadResponse<ReportResponse[]>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/report/view',
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
