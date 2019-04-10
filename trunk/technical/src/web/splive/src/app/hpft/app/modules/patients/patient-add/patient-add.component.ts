import { Component, OnDestroy, OnInit, EventEmitter, Output, Input } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs';
import { ServicepointConfigureListResponse } from '../../../../../prod-shared/models/api/service-configuration-models';
import { ServicepointListResponse } from '../../../../../prod-shared/models/api/servicepoint-models';
import { EnumDataSourceItem } from '../../../../../shared/models/ui/enum-datasource-item';
import { TranslatePipe } from '../../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';
import { EditRecordBase, FORM_MODE, EDITABLE_RECORD_STATE } from '../../../../../shared/views/edit-record-base';
import { PatientAddRequest, PatientUpdateRequest } from '../../../models/api/patient-data-models';
import { PatientAddModal } from '../../../models/ui/patient-models';
import { PatientService } from '../../../services/patient.service';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { RecordIDRequestModel } from '../../../../../shared/models/api/common-models';

@Component({
  selector: 'app-patient-add',
  templateUrl: './patient-add.component.html',
  styleUrls: ['./patient-add.component.css']
})
export class PatientAddComponent extends EditRecordBase implements OnInit, OnDestroy {
  @Input()
  editRecordBase: EditRecordBase;
  @Output()
  editClick = new EventEmitter<null>();
  dataModel = new PatientAddModal();
  routeSubscription: Subscription;
  patientStates: EnumDataSourceItem<number>[];
  splist: ServicepointListResponse[] = [];
  spconfigures: ServicepointConfigureListResponse[] = [];
  disabled: boolean = true;
  personGender: EnumDataSourceItem<number>[];
  skipbutton: boolean;

  constructor(
    private patientService: PatientService,
    private route: ActivatedRoute,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe,
  ) {
    super();
    this.iconCss = 'fa fa-user';
    this.pageTitle = 'Patient Registration';
    this.isEditable = false;
  }

  ngOnInit() {
    this.createControls();
    this.personGender = this.patientService.getPersonGender();
    this.routeSubscription = this.route.queryParams.subscribe(params => {
      if (params['id']) {
        this.dataModel.patientid = Number(params['id']);
        this.recordState = EDITABLE_RECORD_STATE.UPDATE;
        this.setFormMode(FORM_MODE.EDITABLE);
        this.skipbutton = true;
        this.getPatientUpdates();
      } else {
        this.recordState = EDITABLE_RECORD_STATE.ADD;
        this.setFormMode(FORM_MODE.EDITABLE);
        this.dataModel.fname = this.patientService.fname;
        this.dataModel.lname = this.patientService.lname;
      }
      this.callbackUrl = params['callbackurl'];
    });
  }

  getPatientUpdates() {
    const recordIDRequestModel = new RecordIDRequestModel();
    recordIDRequestModel.admissionid = this.patientService.admissionid;
    recordIDRequestModel.patientid = this.dataModel.patientid;
    this.patientService.getPatientNewUpdates(recordIDRequestModel).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        if (payloadResponse.data) {
          this.dataModel.CopyFromUpdateResponse(payloadResponse.data);
          this.recordState = EDITABLE_RECORD_STATE.UPDATE;
        } else {
          this.appNotificationService.info(this.translatePipe.transform('PATIENT_INFO_DETAILS_NOT_AVAILABLE'));
        }
      }
    });
  }

  createControls(): void {
    this.editableForm = new FormGroup({
      fnameControl: new FormControl('', [Validators.required]),
      lnameControl: new FormControl('', [Validators.required]),
      mobnoControl: new FormControl(''),
      ageControl: new FormControl(''),
      bloodgrpControl: new FormControl(''),
      genderControl: new FormControl('', [Validators.required]),
    });
  }

  //Sending data for person add
  save() {
    if (this.editableForm.invalid) { return; }
    this.inProgress = true;
    //for redirect on admission table first
    this.patientService.selcetdIndex = 1;

    if (this.recordState === EDITABLE_RECORD_STATE.ADD) {
      this.add();
    }
    else if (this.dataModel.patientid != null) {
      this.update();
    }
  }

  add() {
    const patientAddRequest = new PatientAddRequest();
    this.dataModel.copyTo(patientAddRequest);
    this.patientService.admissionid = null;
    this.patientService.addPatientData(patientAddRequest).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.dataModel.patientid = payloadResponse.data.recid;
        // setting received patient id in service for further use
        if (payloadResponse.data.recid && payloadResponse.data.recid != null) {
          this.patientService.patientid = payloadResponse.data.recid;
        }
        this.patientService.fname = this.dataModel.fname;
        this.patientService.lname = this.dataModel.lname;
        this.patientService.setPatientName(this.dataModel.fname + ' ' + this.dataModel.lname);
        this.recordState = EDITABLE_RECORD_STATE.ADD;
        this.setFormMode(FORM_MODE.VIEW);
        this.appNotificationService.success();
      }
      this.inProgress = false;
    });
    this.router.navigate(['patients', 'patient_admission'], { queryParams: { id: this.patientService.patientid, callbackurl: 'patients' }, skipLocationChange: true });
  }

  update() {
    const patientUpdateRequest = new PatientUpdateRequest();
    this.dataModel.CopyToUpdate(patientUpdateRequest);
    this.patientService.updatePatientDetails(patientUpdateRequest).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.dataModel.patientid = payloadResponse.data;
        this.appNotificationService.success(this.translatePipe.transform('SUCCESS_USERS_DETAILS_SAVED'));
        //setting received patient id in service for further use
        if (payloadResponse.data && payloadResponse.data != null) {
          this.patientService.patientid = payloadResponse.data;
        }
        this.recordState = EDITABLE_RECORD_STATE.ADD;
        this.setFormMode(FORM_MODE.VIEW);
      }
      this.inProgress = false;
    });
    this.router.navigate(['patients', 'patient_admission'], { queryParams: { id: this.dataModel.patientid, callbackurl: 'patients' }, skipLocationChange: true });
  }

  close() {
    this.router.navigate(['patients', 'patient_search'], { queryParams: { callbackurl: 'patients' }, skipLocationChange: true });
  }

  closeForm() {
    this.router.navigate([this.callbackUrl], { skipLocationChange: true });
  }

  skipForm(id: number) {
    //setting patient id for further use
    this.patientService.patientid = id;
    this.patientService.selcetdIndex = 1;
    this.patientService.admissionid = null;
    this.router.navigate(['patients', 'patient_admission'], { queryParams: { id: id, callbackurl: 'patients' }, skipLocationChange: true });
  }

  //gender value
  getgender(value: number) {
    if (this.personGender && value) {
      return this.personGender.find(a => a.value === value).text;
    }
  }

  ngOnDestroy() {
    if (this.routeSubscription) {
      this.routeSubscription.unsubscribe();
    }
  }
}