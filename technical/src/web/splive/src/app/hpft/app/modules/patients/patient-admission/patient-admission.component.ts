import { Component, OnDestroy, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs';
import { ServicepointListResponse } from '../../../../../prod-shared/models/api/servicepoint-models';
import { TranslatePipe } from '../../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';
import { EDITABLE_RECORD_STATE, EditRecordBase, FORM_MODE } from '../../../../../shared/views/edit-record-base';
import { AdmissionAddRequest, AdmissionUpdateRequest, DrInchargeListResponse } from '../../../models/api/patient-data-models';
import { AdmissionAddModel } from '../../../models/ui/patient-models';
import { PatientService } from '../../../services/patient.service';
import { PATIENT_STATE } from 'app/app-constants';
import { stringify } from '@angular/compiler/src/util';
import { analyzeAndValidateNgModules } from '@angular/compiler';

@Component({
  selector: 'app-patient-admission',
  templateUrl: './patient-admission.component.html',
  styleUrls: ['./patient-admission.component.css']
})
export class PatientAdmissionComponent extends EditRecordBase implements OnInit, OnDestroy {

  dataModel = new AdmissionAddModel();
  routeSubscription: Subscription;
  splist: ServicepointListResponse[];
  drlist: DrInchargeListResponse[];
  PATIENT_STATE = PATIENT_STATE;
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
    this.getServicepointList();
    this.getDrInchargeList();

    this.routeSubscription = this.route.queryParams.subscribe(params => {
      if (params['admissionid']) {
        this.dataModel.admissionid = Number(params['admissionid']);
        this.recordState = EDITABLE_RECORD_STATE.UPDATE;
        this.setFormMode(FORM_MODE.VIEW);
        // if (this.dataModel.status = PATIENT_STATE.DISCHARGED) {
        //   this.setFormMode(FORM_MODE.VIEW)
        // }
        this.getAdmissionUpdates();
      }
      else {
        this.recordState = EDITABLE_RECORD_STATE.ADD;
        this.setFormMode(FORM_MODE.EDITABLE);

      }

      this.callbackUrl = params['callbackurl'];
    });
  }

  // Accept data from ward ie. list of ward
  getServicepointList() {
    this.patientService.getServicepointList().subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.splist = payloadResponse.data;
      }
      else {
        this.appNotificationService.error();
      }
    });
  }

  getDrInchargeList() {
    this.patientService.getDrInchargeList().subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.drlist = payloadResponse.data;
      }
      else {
        this.appNotificationService.error();
      }
    });
  }

  createControls(): void {
    this.editableForm = new FormGroup({
      patientwardControls: new FormControl('', [Validators.required]),
      bednumberControls: new FormControl(''),
      doctorinchargeControls: new FormControl('', [Validators.required]),
      admittedDateControls: new FormControl(''),
      patientregnoControls: new FormControl('', [Validators.required]),
    });
  }

  //Sending data for admission add
  save() {
    if (this.editableForm.invalid) { return; }
    this.inProgress = true;
    if (this.recordState == EDITABLE_RECORD_STATE.ADD) {
      const admissionAddRequest = new AdmissionAddRequest();
      this.dataModel.patientid = this.patientService.patientid;
      this.dataModel.status = PATIENT_STATE.HOSPITALIZE;
      this.dataModel.copyTo(admissionAddRequest);
      this.patientService.admissionAddPatient(admissionAddRequest).subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.dataModel.admissionid = payloadResponse.data.admissionid;
          this.patientService.setAdmissionId(payloadResponse.data.admissionid);
          this.appNotificationService.success();
          this.patientService.medicaldetialsid = payloadResponse.data.medicaldetailsid;
          this.patientService.personaldetailsid = payloadResponse.data.personaldetailsid;
          this.patientService.admissionid = payloadResponse.data.admissionid;
          this.recordState = EDITABLE_RECORD_STATE.UPDATE;
          this.setFormMode(FORM_MODE.VIEW);
        }
        this.inProgress = false;
      });
    }
    else {
      const admissionUpdateRequest = new AdmissionUpdateRequest();
      this.dataModel.admissionid = this.patientService.admissionid;
      this.dataModel.copyToUpdate(admissionUpdateRequest);
      this.patientService.updateAdmissionRequest(admissionUpdateRequest).subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.appNotificationService.success();
          this.recordState = EDITABLE_RECORD_STATE.UPDATE;
          this.setFormMode(FORM_MODE.VIEW);
        }
        this.inProgress = false;
      });
    }
  }

  getAdmissionUpdates() {
    this.patientService.getAdmissionUpdates({ recid: this.dataModel.admissionid }).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        if (payloadResponse.data) {
          this.dataModel.copyFromUpdateResponse(payloadResponse.data);
          this.recordState = EDITABLE_RECORD_STATE.UPDATE;
          this.setFormMode(FORM_MODE.VIEW);
          if (this.dataModel.status === PATIENT_STATE.DISCHARGED) {
            this.isEditable = false;
          }
        }
      }
    });
  }

  closeForm() {
    this.router.navigate([this.callbackUrl], { skipLocationChange: true });
  }

  ngOnDestroy() {
    if (this.routeSubscription) {
      this.routeSubscription.unsubscribe();
    }
  }

  getSPName(value: number) {
    if (this.splist && value) {
      return this.splist.find(a => a.spid === value).spname;
    }
  }
  getDrName(value: number) {
    if (this.drlist && value) {
      const user = this.drlist.find(a => a.usrid === value);
      return user.fname + ' ' + user.lname;
    }
  }

}