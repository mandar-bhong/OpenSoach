import { Component, EventEmitter, NgZone, OnDestroy, OnInit, ViewChild } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { MatPaginator, MatSort } from '@angular/material';
import { Router } from '@angular/router';
import { AmazingTimePickerService } from 'amazing-time-picker';
import { merge, Observable, Subscription } from 'rxjs';
import { map, startWith, switchMap } from 'rxjs/operators';

import { ServicepointListResponse } from '../../../../../../prod-shared/models/api/servicepoint-models';
import { DataListRequest, DataListResponse } from '../../../../../../shared/models/api/data-list-models';
import { PayloadResponse } from '../../../../../../shared/models/api/payload-models';
import { TranslatePipe } from '../../../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../../../shared/services/notification/app-notification.service';
import { PatientListDataModel } from '../../../../../app/models/ui/patient-models';
import {
  AdmissionStatusRequest,
  PatientDetaListResponse,
  PatientFilterRequest,
} from '../../../../models/api/patient-data-models';
import { PatientService } from '../../../../services/patient.service';
import { PATIENT_STATE } from 'app/app-constants';

@Component({
  selector: 'app-patient-view',
  templateUrl: './patient-view.component.html',
  styleUrls: ['./patient-view.component.css']
})
export class PatientViewComponent implements OnInit, OnDestroy {

  displayedColumns = ['fname', 'patientregno', 'emergencycontactno', 'mobno', 'spid', 'bedno', 'status', 'action'];
  sortByColumns = [{ text: 'Patient Name', value: 'fname' },
  { text: 'Patient Reg No', value: 'patientregno' },
  { text: 'Emergency Contact', value: 'emergencycontactno' },
  { text: 'Contact No', value: 'mobno' },
  { text: 'Ward', value: 'spid' },
  { text: 'Room/Bed No', value: 'bedno' },
  { text: 'Status', value: 'status' }

  ];
  @ViewChild(MatPaginator)
  paginator: MatPaginator;
  @ViewChild(MatSort)
  sort: MatSort;
  refreshTable: EventEmitter<null> = new EventEmitter();
  dataSource;
  filteredrecords = 0;
  isLoadingResults = true;
  patientFilterRequest: PatientFilterRequest;
  dataListFilterChangedSubscription: Subscription;
  showEditForm = false;
  selectedPatient: PatientListDataModel;
  splist: ServicepointListResponse[] = [];
  dataModel = new PatientListDataModel();
  editableForm: FormGroup;
  selectedStartTime: string;
  PATIENT_STATE = PATIENT_STATE;
  admittedDate: Date;

  constructor(public patientService: PatientService,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private amazingtimepicker: AmazingTimePickerService,
    private translatePipe: TranslatePipe) {
  }

  ngOnInit() {
    this.getServicepointList();
    this.createControls();
    this.dataModel.dischargedon = new Date();
    // this.dataModel.dischargedon = new Date(new Date().toLocaleDateString("en-US")).toISOString().substr(0, 10);
    const dt = new Date();

    this.selectedStartTime = dt.getHours() + ":" + dt.getMinutes();
    this.patientFilterRequest = new PatientFilterRequest();
    this.patientFilterRequest.status = PATIENT_STATE.HOSPITALIZE;
    this.paginator.pageSize = 10;
    this.sort.active = 'fname';
    this.sort.direction = 'asc';
    this.dataListFilterChangedSubscription = this.patientService.dataListSubject.subscribe(value => {
      this.patientFilterRequest = value;
      this.refreshTable.emit();
    });
  }

  createControls(): void {
    this.editableForm = new FormGroup({
      dischargedDateControls: new FormControl('', [Validators.required]),
      dischargedTimeControls: new FormControl('', [Validators.required]),
    });
  }

  // Accept data from ward ie. list of ward
  getServicepointList() {
    this.patientService.getServicepointList().subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.splist = payloadResponse.data;
        this.setDataListing();
      }
    });
  }

  setSelectedPatient(patient: PatientListDataModel) {
    this.selectedPatient = patient;
    this.admittedDate = this.selectedPatient.admittedon;
  }

  // Changing status , patient is dicharge or not
  changestatus() {
    if (this.editableForm.invalid) { return; }
    const admissionStatusRequest = new AdmissionStatusRequest();
    admissionStatusRequest.status = PATIENT_STATE.DISCHARGED;
    admissionStatusRequest.admissionid = this.selectedPatient.admissionid;
    admissionStatusRequest.dischargedon = this.dataModel.dischargedon;
    this.patientService.updateAdmissionStatus(admissionStatusRequest).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.appNotificationService.success();
        this.selectedPatient.status = PATIENT_STATE.DISCHARGED;
      }
    });
  }


  // For Time selection
  openStartTime() {
    const amazingTimePicker = this.amazingtimepicker.open({
      time: this.selectedStartTime,
      theme: 'material-orange',
    });
    // this.selectedStartTime = this.dataModel.dischargedon.toLocaleTimeString();
    amazingTimePicker.afterClose().subscribe(time => {
      this.selectedStartTime = time;
    });

  }

  setDataListing(): void {
    this.sort.sortChange.subscribe(() => this.paginator.pageIndex = 0);
    this.refreshTable.subscribe(() => this.paginator.pageIndex = 0);
    merge(this.sort.sortChange, this.paginator.page, this.refreshTable)
      .pipe(
        startWith({}),
        switchMap(() => {
          this.isLoadingResults = true;
          return this.getDataList();
        }),
        map(data => {
          this.isLoadingResults = false;
          return data;
        }),
      ).subscribe(
        payloadResponse => {
          if (payloadResponse && payloadResponse.issuccess) {
            this.filteredrecords = payloadResponse.data.filteredrecords;
            this.dataSource = payloadResponse.data.records;

            if (this.filteredrecords === 0) {
              this.appNotificationService.info(this.translatePipe.transform('INFO_NO_RECORDS_FOUND'));
            }
          } else {
            this.dataSource = [];

          }
        });
  }

  getDataList(): Observable<PayloadResponse<DataListResponse<PatientDetaListResponse>>> {
    const dataListRequest = new DataListRequest<PatientFilterRequest>();
    dataListRequest.filter = this.patientFilterRequest;
    dataListRequest.page = this.paginator.pageIndex + 1;
    dataListRequest.limit = this.paginator.pageSize;
    dataListRequest.orderby = this.sort.active;
    dataListRequest.orderdirection = this.sort.direction;
    return this.patientService.getDataList(dataListRequest);
  }
  viewDetails(id: number, admissionid: number, personaldetailsid: number) {
    //setting patient id for further use
    this.patientService.patientid = id;
    this.patientService.admissionid = admissionid;
    this.patientService.personaldetailsid = personaldetailsid;
    // this.patientService.selcetdIndex = 0;
    // this.router.navigate(['patients', 'patient_admission'], { queryParams: { id: id, admissionid: admissionid,personaldetailsid:personaldetailsid, callbackurl: 'patients' }, skipLocationChange: true });
    if (this.patientService.admissionid != null){
      this.patientService.selcetdIndex = 0;
      this.router.navigate(['patients', 'patient_admission'], { queryParams: { id: id, admissionid: admissionid, personaldetailsid: personaldetailsid, callbackurl: 'patients' }, skipLocationChange: true });
    }
    else{
      this.patientService.selcetdIndex = 1;
      this.router.navigate(['patients', 'patient_admission'], { queryParams: { id: id, callbackurl: 'patients' }, skipLocationChange: true });
    }
  }



  sortByChanged() {
    this.sort.sortChange.next(this.sort);
  }

  sortDirectionAsc() {
    this.sort.direction = 'asc';
    this.sort.sortChange.next(this.sort);
  }

  sortDirectionDesc() {
    this.sort.direction = 'desc';
    this.sort.sortChange.next(this.sort);
  }

  ngOnDestroy(): void {
    if (this.dataListFilterChangedSubscription) {
      this.dataListFilterChangedSubscription.unsubscribe();
    }
  } // end 

  getSPName(value: number) {
    if (this.splist && value) {
      return this.splist.find(a => a.spid === value).spname;
    }
  }

  cancelStatus() {
    this.dataModel.dischargedon = new Date();
    const dt = new Date();
    this.selectedStartTime = dt.getHours() + ":" + dt.getMinutes();
  }

}
