import { Injectable } from '@angular/core';
import { Observable, Subject } from 'rxjs';

import { ServicepointConfigureListResponse } from '../../../prod-shared/models/api/service-configuration-models';
import { ServicepointListResponse } from '../../../prod-shared/models/api/servicepoint-models';
import { EnvironmentProvider } from '../../../shared/environment-provider';
import {  RecordIDRequest, RecordIDResponse } from '../../../shared/models/api/common-models';
import { DataListRequest, DataListResponse } from '../../../shared/models/api/data-list-models';
import { PayloadResponse } from '../../../shared/models/api/payload-models';
import { EnumDataSourceItem } from '../../../shared/models/ui/enum-datasource-item';
import { ServerApiInterfaceService } from '../../../shared/services/api/server-api-interface.service';
import { ListingService } from '../../../shared/services/listing.service';
import { EnumNumberDatasource } from '../../../shared/utility/enum-number-datasource';
import { PATIENT_STATE, PERSON_GENDER } from '../app-constants';
import {
    AdmissionAddRequest,
    AdmissionUpdateRequest,
    AdmissionUpdateResponse,
    PatientAddRequest,
    PatientDataAddRequest,
    PatientDetaListResponse,
    PatientFilterRequest,
    PatientSearchRequestFilter,
    PatientSearchResponseFilter,
    PatientUpdateRequest,
    PatientUpdateResponse,
    PersonDetailResponse,
    StatusChangeRequest,
    AdmissionAddResponseModel,
} from '../models/api/patient-models';


@Injectable()
export class PatientService extends ListingService<PatientFilterRequest, PatientDetaListResponse> {
    patientid: number;
    selcetdIndex: number;
    disableTab: boolean;
    fname: string;
    lname: string;
    drincharge: number;
    admissionIdReceived = new Subject<number>();

    constructor(private serverApiInterfaceService: ServerApiInterfaceService) {
        super();
    }

    setAdmissionId(value: number) {
        this.admissionIdReceived.next(value);
    }
    getDataList(dataListRequest: DataListRequest<PatientFilterRequest>, implicitErrorHandling = true):
        Observable<PayloadResponse<DataListResponse<PatientDetaListResponse>>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/patient/list',
            dataListRequest, implicitErrorHandling);
    }

    getPatientDataList(dataListRequest: DataListRequest<PatientSearchRequestFilter>, implicitErrorHandling = true):
        Observable<PayloadResponse<DataListResponse<PatientSearchResponseFilter>>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/patient/list/master',
            dataListRequest, implicitErrorHandling);
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

    // New
    addPatientData(patientAddRequest: PatientAddRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<RecordIDResponse>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.appbaseurl + '/api/v1/patient/add',
            patientAddRequest, implicitErrorHandling);
    }

    //Post method for patient addmission
    admissionAddPatient(admissionAddRequest: AdmissionAddRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<AdmissionAddResponseModel>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.appbaseurl + '/api/v1/patient/admission/add',
            admissionAddRequest, implicitErrorHandling);
    }

    //Update Patient detail Request
    updatePatientDetails(patientUpdate: PatientUpdateRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<null>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.appbaseurl + '/api/v1/patient/update',
            patientUpdate, implicitErrorHandling);
    }

    //Update Patient Response
    getPatientUpdates(request: RecordIDRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<PatientUpdateResponse>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/patient/info',
            request, implicitErrorHandling);
    }

    //Update Admission Response
    getAdmissionUpdates(request: RecordIDRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<AdmissionUpdateResponse>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/patient/admission/info',
            request, implicitErrorHandling);
    }
    //Update Admission Request
    updateAdmissionRequest(admissionUpdateRequest: AdmissionUpdateRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<null>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.appbaseurl + '/api/v1/patient/admission/update',
            admissionUpdateRequest, implicitErrorHandling);
    }
    //Update status
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

    // for addmission - pateint - person
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
    getPersonGender(): EnumDataSourceItem<number>[] {
        return EnumNumberDatasource.getDataSource('PERSON_GENDER_', PERSON_GENDER);
    }

    getPersonGenders(genders: number) {
        return 'PERSON_GENDER_' + genders;
    }

    //Update Patient Response
    getPatientPersonDetail(request: RecordIDRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<PersonDetailResponse>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/patient/personaldetails/info',
            request, implicitErrorHandling);
    }
}
