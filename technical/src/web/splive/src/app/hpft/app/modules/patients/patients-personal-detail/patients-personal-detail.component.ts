import { Component, OnInit, OnDestroy, ViewChild, EventEmitter } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription, Observable, merge } from 'rxjs';
import { EnumDataSourceItem } from '../../../../../shared/models/ui/enum-datasource-item';
import { TranslatePipe } from '../../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';
import { EditRecordBase, EDITABLE_RECORD_STATE, FORM_MODE } from '../../../../../shared/views/edit-record-base';
import { PatientAddModal } from '../../../models/ui/patient-models';
import { PatientService } from '../../../services/patient.service';
import { FormGroup, FormControl } from '@angular/forms';
import { PatientUpdateRequest } from 'app/models/api/patient-models';
import { RecordIDRequestModel } from '../../../../../shared/models/api/common-models';

@Component({
  selector: 'app-patients-personal-detail',
  templateUrl: './patients-personal-detail.component.html',
  styleUrls: ['./patients-personal-detail.component.css']
})
export class PatientsPersonalDetailComponent extends EditRecordBase implements OnInit, OnDestroy {

  dataModel = new PatientAddModal();
  routeSubscription: Subscription;
  patientStates: EnumDataSourceItem<number>[];
  personGender: EnumDataSourceItem<number>[];

  constructor(
    private patientService: PatientService,
    private route: ActivatedRoute,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe,
  ) {
    super();
    this.iconCss = 'fa fa-user';
    this.pageTitle = 'Details';
  }

  ngOnInit() {
    this.createControls();
    this.showBackButton = false;
    this.personGender = this.patientService.getPersonGender();
    this.routeSubscription = this.route.queryParams.subscribe(params => {
      if (params['id']) {
        this.dataModel.patientid = Number(params['id']);
        this.recordState = EDITABLE_RECORD_STATE.UPDATE;
        this.setFormMode(FORM_MODE.VIEW);
        this.getPatientUpdates();
      } else {
        this.recordState = EDITABLE_RECORD_STATE.UPDATE;
        this.setFormMode(FORM_MODE.EDITABLE);
      }
      this.callbackUrl = params['callbackurl'];
    });
  }

  getPatientUpdates() {
    const recordIDRequestModel = new RecordIDRequestModel();
    recordIDRequestModel.admissionid = this.patientService.admissionid;
    recordIDRequestModel.patientid = this.patientService.patientid;
    this.patientService.getPatientNewUpdates(recordIDRequestModel).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        if (payloadResponse.data) {
          this.dataModel.CopyFromUpdateResponse(payloadResponse.data);
          this.recordState = EDITABLE_RECORD_STATE.UPDATE;
          this.patientService.fname = this.dataModel.fname;
          this.patientService.lname = this.dataModel.lname;
          this.patientService.setPatientName(this.dataModel.fname+' '+this.dataModel.lname);
        } else {
          this.appNotificationService.info(this.translatePipe.transform('PATIENT_INFO_DETAILS_NOT_AVAILABLE'));
        }
      }
    });
  }
  save() {
    if (this.recordState === EDITABLE_RECORD_STATE.UPDATE) {
      const patientUpdateRequest = new PatientUpdateRequest();
      this.dataModel.patientid = this.patientService.patientid;
      this.dataModel.CopyToUpdate(patientUpdateRequest);
      this.patientService.updatePatientDetails(patientUpdateRequest).subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.patientService.setPatientName(this.dataModel.fname+' '+this.dataModel.lname);
          this.dataModel.patientid = payloadResponse.data;
          this.recordState = EDITABLE_RECORD_STATE.UPDATE;
          this.setFormMode(FORM_MODE.VIEW);
          this.appNotificationService.success();
        }
        this.inProgress = false;
      });
    }
  }
  createControls(): void {
    this.editableForm = new FormGroup({
      fnameControl: new FormControl(''),
      lnameControl: new FormControl(''),
      mobnoControl: new FormControl(''),
      ageControl: new FormControl(''),
      bloodgrpControl: new FormControl(''),
      genderControl: new FormControl(''),
    });
  }

  //gender value
  getgender(value: number) {
    if (this.personGender && value) {
      return this.personGender.find(a => a.value === value).text;
    }
  }

  closeForm() {
    this.router.navigate([this.callbackUrl], { skipLocationChange: true });
  }

  ngOnDestroy() {
    if (this.routeSubscription) {
      this.routeSubscription.unsubscribe();
    }
  }
}
