import { Injectable } from '@angular/core';
import { Observable, Subject } from 'rxjs';

import { ServicepointConfigureListResponse } from '../../../prod-shared/models/api/service-configuration-models';
import { ServicepointListResponse } from '../../../prod-shared/models/api/servicepoint-models';
import { EnvironmentProvider } from '../../../shared/environment-provider';
import { RecordIDRequest, RecordIDResponse, RecordIDRequestModel } from '../../../shared/models/api/common-models';
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
    StatusChangeRequest,
    AdmissionAddResponseModel,
    AdmissionStatusRequest,
    MedicalDetailsRequest,
    MedicalDetailsResponse,
    PresentComplaint,
    ReasonForAdmission,
    HistoryPresentIllness,
    PastHistory,
    TreatmentBeforeAdmission,
    InvestigationBeforeAdmission,
    FamilyHistory,
    Allergies,
    PersonalHistory,
    PersonAccompanyingInfo,
    PersonalDetailsResponse,
    PathologyFilterRequest,
    PathologyResponse,
    TreatmentResponse,
    TreatmentFilterRequest,
    PersonalDetailsRequest,
    CheckPatientResponse,
    DrInchargeListResponse,
    PersonalHistoryInfo,
} from '../models/api/patient-models';
import { TransactionDetailsFilter } from 'app/models/api/transaction-details';
import { ActionTransactionResponse } from 'app/models/api/transaction-details-response';
import { ScheduleFilter } from 'app/models/api/schedule-request';
import { ScheduleDataResponse } from 'app/models/api/schedule-response';
import { DoctorOrderRequest } from 'app/models/api/doctor-orders-request';
import { DoctorOrderResponse } from 'app/models/api/doctor-order-response';
import { SaveFileService } from '../../../shared/services/save-file.service';
import { FileDownloadRequest } from 'app/models/api/file-download-request';
import { PathologyReportAddRequest } from 'app/models/api/pathology-report-add-request';
import { TreatmentReportRequest } from 'app/models/api/treatment-report-request';


@Injectable()
export class PatientService extends ListingService<PatientFilterRequest, PatientDetaListResponse> {
    patientid: number;
    admissionid: number;
    selcetdIndex: number;
    disableTab: boolean;
    fname: string;
    lname: string;
    admittedon: Date;
    drincharge: number;
    admissionIdReceived = new Subject<number>();
    medicaldetialsid: number;
    personaldetailsid: number;
    treatmentid: number;
    pathologyid: number;

    constructor(private serverApiInterfaceService: ServerApiInterfaceService,
        private saveFileService: SaveFileService) {
        super();
    }

    setAdmissionId(value: number) {
        this.admissionIdReceived.next(value);
    }

    //Get data for admission master table
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
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/patient/info/details',
            request, implicitErrorHandling);
    }

    //Update New Patient Response
    getPatientNewUpdates(request: RecordIDRequestModel, implicitErrorHandling = true):
        Observable<PayloadResponse<PatientUpdateResponse>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/patient/info/details',
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

    //Update patient status on discharge
    updateAdmissionStatus(admissionStatusRequest: AdmissionStatusRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<null>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.appbaseurl + '/api/v1/patient/update/status',
            admissionStatusRequest, implicitErrorHandling);
    }

    //Post method for patient medical Complaint
    medicalAddPatientComplaint(medicalDetailsRequest: PresentComplaint, implicitErrorHandling = true):
        Observable<PayloadResponse<any>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.appbaseurl + '/api/v1/patient/medicaldetails/update/presentcomplaints',
            medicalDetailsRequest, implicitErrorHandling);
    }

    //Post method for patient medical Reason For Admission
    medicalAddPatientReasonForAdmission(reasonForAdmission: ReasonForAdmission, implicitErrorHandling = true):
        Observable<PayloadResponse<any>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.appbaseurl + '/api/v1/patient/medicaldetails/update/reasonforadmission',
            reasonForAdmission, implicitErrorHandling);
    }

    //Post method for patient medical History Present Illness
    medicalAddPatientHistoryPresentIllness(historyPresentIllness: HistoryPresentIllness, implicitErrorHandling = true):
        Observable<PayloadResponse<any>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.appbaseurl + '/api/v1/patient/medicaldetails/update/historypresentillness',
            historyPresentIllness, implicitErrorHandling);
    }

    //Post method for patient medical Past History
    medicalAddPatientPastHistory(pastHistory: PastHistory, implicitErrorHandling = true):
        Observable<PayloadResponse<any>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.appbaseurl + '/api/v1/patient/medicaldetails/update/pasthistory',
            pastHistory, implicitErrorHandling);
    }

    //Post method for patient medical Treatment Before Admission
    medicalAddPatientTreatmentBeforeAdmission(treatmentBeforeAdmission: TreatmentBeforeAdmission, implicitErrorHandling = true):
        Observable<PayloadResponse<any>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.appbaseurl + '/api/v1/patient/medicaldetails/update/treatmentbeforeadmission',
            treatmentBeforeAdmission, implicitErrorHandling);
    }

    //Post method for patient medical Investigation Before Admission
    medicalAddPatientInvestigationBeforeAdmission(investigationBeforeAdmission: InvestigationBeforeAdmission, implicitErrorHandling = true):
        Observable<PayloadResponse<any>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.appbaseurl + '/api/v1/patient/medicaldetails/update/investigationbeforeadmission',
            investigationBeforeAdmission, implicitErrorHandling);
    }

    //Post method for patient medical Family History
    medicalAddPatientFamilyHistory(familyHistory: FamilyHistory, implicitErrorHandling = true):
        Observable<PayloadResponse<any>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.appbaseurl + '/api/v1/patient/medicaldetails/update/familyhistory',
            familyHistory, implicitErrorHandling);
    }

    //Post method for patient medical Allergies
    medicalAddPatientAllergies(allergiess: Allergies, implicitErrorHandling = true):
        Observable<PayloadResponse<any>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.appbaseurl + '/api/v1/patient/medicaldetails/update/allergies',
            allergiess, implicitErrorHandling);
    }

    //Post method for patient medical Personal History
    medicalAddPatientPersonalHistory(personalHistory: PersonalHistory, implicitErrorHandling = true):
        Observable<PayloadResponse<PersonalHistoryInfo[]>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.appbaseurl + '/api/v1/patient/medicaldetails/update/personalhistory',
            personalHistory, implicitErrorHandling);
    }

    //Post method for patient personal Add Accompanying
    personalAddAccompanying(personalDetailsRequest: PersonalDetailsRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<any>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.appbaseurl + '/api/v1/patient/personaldetails/update/personaccompanying',
            personalDetailsRequest, implicitErrorHandling);
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
        Observable<PayloadResponse<PersonalDetailsResponse>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/patient/personaldetails/info',
            request, implicitErrorHandling);
    }
    // service function for getting transaction details.
    getActionTransaction(dataListRequest: DataListRequest<TransactionDetailsFilter>, implicitErrorHandling = true):
        Observable<PayloadResponse<DataListResponse<ActionTransactionResponse<string>[]>>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/patient/list/actiontxn',
            dataListRequest, implicitErrorHandling);
    }
    // service function for getting schedule details.
    getScheduleData(dataListRequest: DataListRequest<ScheduleFilter>, implicitErrorHandling = true):
        Observable<PayloadResponse<DataListResponse<ScheduleDataResponse<string>[]>>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/patient/list/patientconf',
            dataListRequest, implicitErrorHandling);
    }
    // service function for getting schedule details.
    getScheduleDataById(dataListRequest: ScheduleFilter, implicitErrorHandling = true):
        Observable<PayloadResponse<ScheduleDataResponse<string>>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/patient/config/info',
            dataListRequest, implicitErrorHandling);
    }

    //Get Patient MedicalDetailsResponse Response
    getPatientMedicalDetail(request: RecordIDRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<MedicalDetailsResponse>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/patient/medicaldetails/info',
            request, implicitErrorHandling);
    }


    //Get Patient MedicalDetailsResponse Response
    getPatientMedicalID(request: RecordIDRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<any>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/patient/admission/info/details',
            request, implicitErrorHandling);
    }
    //Get Patient Personal detail person accompanying Response
    getPatientPersonalDetail(request: RecordIDRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<PersonalDetailsResponse>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/patient/personaldetails/info',
            request, implicitErrorHandling);
    }

    //Get data for pathology report
    getPathologyList(dataListRequest: DataListRequest<PathologyFilterRequest>, implicitErrorHandling = true):
        Observable<PayloadResponse<DataListResponse<PathologyResponse>>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/patient/list/pathologyrecord',
            dataListRequest, implicitErrorHandling);
    }

    //Get data for pathology report
    getTreatmentList(dataListRequest: DataListRequest<TreatmentFilterRequest>, implicitErrorHandling = true):
        Observable<PayloadResponse<DataListResponse<TreatmentResponse>>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/patient/list/treatment',
            dataListRequest, implicitErrorHandling);
    }

    // service function for getting doctor order details
    getDoctorOrderDetails(dataListRequest: DataListRequest<DoctorOrderRequest>, implicitErrorHandling = true):
        Observable<PayloadResponse<DataListResponse<DoctorOrderResponse>>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/patient/list/doctororders',
            dataListRequest, implicitErrorHandling);
    }
    // service fucntion for downloading file
    downloadFile(request: FileDownloadRequest, implicitErrorHandling = true):
        Observable<Blob> {
        return this.serverApiInterfaceService.downloadFile(EnvironmentProvider.appbaseurl + '/api/v1/document/download',
            request, implicitErrorHandling);
    }

    saveFile(data: Blob, filename: string) {
        this.saveFileService.saveFile(data, filename);
    }
    //Post method for docuemnt upload
    uploadReportsDocuments(formData: any, implicitErrorHandling = true):
        Observable<PayloadResponse<any>> {
        return this.serverApiInterfaceService.upload(EnvironmentProvider.appbaseurl + '/api/v1/document/upload',
            formData, implicitErrorHandling);
    }
    //Post method for docuemnt upload
    addReportData(request: PathologyReportAddRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<any>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.appbaseurl + '/api/v1/patient/pathologyrecord/add',
            request, implicitErrorHandling);
    }
    //Post method for docuemnt upload
    addTeatmentReportData(request: TreatmentReportRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<any>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.appbaseurl + '/api/v1/patient/treatment/add',
            request, implicitErrorHandling);
    }

    //  Update Admission Response
    getStatusCheck(request: RecordIDRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<CheckPatientResponse>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/patient/admission/info/status',
            request, implicitErrorHandling);
    }

    getDrInchargeList(implicitErrorHandling = true):
        Observable<PayloadResponse<DrInchargeListResponse[]>> {
        return this.serverApiInterfaceService.get(EnvironmentProvider.appbaseurl + '/api/v1/user/doctorlist',
            implicitErrorHandling);
    }

}
