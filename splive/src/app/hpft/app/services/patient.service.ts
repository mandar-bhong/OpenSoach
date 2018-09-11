import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

import { EnvironmentProvider } from '../../../shared/environment-provider';
import { PayloadResponse } from '../../../shared/models/api/payload-models';
import { ServerApiInterfaceService } from '../../../shared/services/api/server-api-interface.service';
import { SaveFileService } from '../../../shared/services/save-file.service';
import { ReportRequestParams, ReportResponse } from '../models/api/report-models';
import { ListingService } from '../../../shared/services/listing.service';
import {
    PatientDetaFilterRequest, PatientDataAddRequest, PatientDataListResponse, StatusChangeRequest,
} from '../models/api/patient-models';
import { DataListRequest, DataListResponse } from '../../../shared/models/api/data-list-models';
import { RecordIDRequest, RecordIDResponse } from '../../../shared/models/api/common-models';

import { EnumDataSourceItem } from '../../../shared/models/ui/enum-datasource-item';
import { EnumNumberDatasource } from '../../../shared/utility/enum-number-datasource';
import { ServicepointListResponse } from '../../../prod-shared/models/api/servicepoint-models';
import { ServicepointConfigureListResponse } from '../../../prod-shared/models/api/service-configuration-models';
import { PATIENT_STATE } from '../app-constants';


@Injectable()
export class PatientService {
    constructor(private serverApiInterfaceService: ServerApiInterfaceService,
        private saveFileService: SaveFileService) {
        // super();
    }

    getDataList(implicitErrorHandling = true):
        Observable<PayloadResponse<PatientDataAddRequest[]>> {
        return this.serverApiInterfaceService.get(EnvironmentProvider.appbaseurl + '/api/v1/patient/list',
            implicitErrorHandling);
    }

    getPatientDetails(request: RecordIDRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<PatientDataAddRequest>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/patient/info',
            request, implicitErrorHandling);
    }
    addPatient(patientDataAddRequest: PatientDataAddRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<RecordIDResponse>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.appbaseurl + '/api/v1/patient/add',
            patientDataAddRequest, implicitErrorHandling);
    }
    updateStatus(statusChangeRequest: StatusChangeRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<null>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.appbaseurl + '/api/v1/patient/update/status',
        statusChangeRequest, implicitErrorHandling);
    }
    getPatientStates(): EnumDataSourceItem<number>[] {
        return EnumNumberDatasource.getDataSource('PATIENT_STATE_', PATIENT_STATE);
    }
    getPatientState(states: number) {
        return 'PATIENT_STATE_' + states;
    }
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
    getWardStates(): EnumDataSourceItem<number>[] {
        return EnumNumberDatasource.getDataSource('PATIENT_STATE_', PATIENT_STATE);
    }
    getWardState(states: number) {
        return 'PATIENT_STATE_' + states;
    }
}
