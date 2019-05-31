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
    SnapShotRequest,
    SnapShotResponse,
    TimeRequest,
    TimeResponse,
    ServiceTimeAvrResponse,
    ServiceTimeAvrRequest,
    VehicleServiceTrendMontlyResponse,
    VehicleServiceTrendMontlyRequest,
    VehicleServiceTrendWeeklyResponse,
    VehicleServiceTrendWeeklyRequest
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

    getComplaintSummary(request = new ComplaintSummaryRequest(), implicitErrorHandling = true):
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
    getSnapShot(request: SnapShotRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<SnapShotResponse[]>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/dashboard/vehicle/snapshot',
            request, implicitErrorHandling);
    }
    getTime(request: TimeRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<TimeResponse>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/dashboard/vehicle/averagetime',
            request, implicitErrorHandling);
    }
    getSeviceTimeMonth(request = new ServiceTimeAvrRequest(), implicitErrorHandling = true):
        Observable<PayloadResponse<ServiceTimeAvrResponse[]>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/dashboard/averagetime/permonth',
            request, implicitErrorHandling);
    }
    getVehicleServiceTrendMontly(request = new VehicleServiceTrendMontlyRequest(), implicitErrorHandling = true):
        Observable<PayloadResponse<VehicleServiceTrendMontlyResponse[]>> {
        return this.serverApiInterfaceService.getWithQueryParams(
            EnvironmentProvider.appbaseurl + '/api/v1/dashboard/vehicleserviced/permonth',
            request, implicitErrorHandling);
    }
    getVehicleServiceTrendWeekly(request = new VehicleServiceTrendWeeklyRequest(), implicitErrorHandling = true):
        Observable<PayloadResponse<VehicleServiceTrendWeeklyResponse[]>> {
        return this.serverApiInterfaceService.getWithQueryParams(
            EnvironmentProvider.appbaseurl + '/api/v1/dashboard/vehicleserviced/perweek',
            request, implicitErrorHandling);
    }
}
