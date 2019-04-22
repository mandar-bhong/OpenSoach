import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

import { EnvironmentProvider } from '../../../shared/environment-provider';
import { PayloadResponse } from '../../../shared/models/api/payload-models';
import { ServerApiInterfaceService } from '../../../shared/services/api/server-api-interface.service';
import {
    ComplaintSummaryRequest,
    ComplaintSummaryResponse,
    DeviceSummaryResponse,
    FeedbackSummaryRequest,
    FeedbackSummaryResponse,
    ServicePointSummaryResponse,
    TaskSummaryRequest,
    TaskSummaryResponse,
    FeedbackTrendResponse,
    FeedbackTrendRequest,
    TaskTrendRequest,
    TaskTrendResponse,
    ComplaintTrendResponse,
    ComplaintTrendRequest,
    PatientSummaryRequest,
    PatientSummaryResponse,
} from '../models/api/dashboard-models';

@Injectable()
export class DashboardService {

    constructor(private serverApiInterfaceService: ServerApiInterfaceService) { }
    getDeviceSummary(implicitErrorHandling = true):
        Observable<PayloadResponse<DeviceSummaryResponse>> {
        return this.serverApiInterfaceService.get(EnvironmentProvider.appbaseurl + '/api/v1/dashboard/device/summary',
            implicitErrorHandling);
    }

    getServicePointSummary(implicitErrorHandling = true):
        Observable<PayloadResponse<ServicePointSummaryResponse>> {
        return this.serverApiInterfaceService.get(EnvironmentProvider.appbaseurl + '/api/v1/dashboard/location/summary',
            implicitErrorHandling);
    }

    getComplaintSummary(request= new ComplaintSummaryRequest(), implicitErrorHandling = true):
        Observable<PayloadResponse<ComplaintSummaryResponse>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/dashboard/complaint/summary',
            request, implicitErrorHandling);
    }

    getFeedbackSummary(request = new FeedbackSummaryRequest(), implicitErrorHandling = true):
        Observable<PayloadResponse<FeedbackSummaryResponse>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/dashboard/feedback/summary',
            request, implicitErrorHandling);
    }

    getTaskSummary(request = new TaskSummaryRequest(), implicitErrorHandling = true):
        Observable<PayloadResponse<TaskSummaryResponse>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/dashboard/task/summary',
            request, implicitErrorHandling);
    }

    getFeedbackTrend(request = new FeedbackTrendRequest(), implicitErrorHandling = true):
        Observable<PayloadResponse<FeedbackTrendResponse[]>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/dashboard/feedback/permonth',
            request, implicitErrorHandling);
    }

    getTaskTrend(request = new TaskTrendRequest(), implicitErrorHandling = true):
        Observable<PayloadResponse<TaskTrendResponse[]>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/dashboard/task/permonth',
            request, implicitErrorHandling);
    }

    getComplaintTrend(request = new ComplaintTrendRequest(), implicitErrorHandling = true):
        Observable<PayloadResponse<ComplaintTrendResponse[]>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/dashboard/complaint/permonth',
            request, implicitErrorHandling);
    }

    getPatientSummary(request= new PatientSummaryRequest(), implicitErrorHandling = true):
        Observable<PayloadResponse<PatientSummaryResponse>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/dashboard/patient/summary',
            request, implicitErrorHandling);
    }
}
