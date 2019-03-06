import { Component, OnDestroy, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs';

import { ServicepointListResponse } from '../../../../../prod-shared/models/api/servicepoint-models';
import { TranslatePipe } from '../../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';
import { EDITABLE_RECORD_STATE, EditRecordBase, FORM_MODE } from '../../../../../shared/views/edit-record-base';
import { AdmissionAddRequest, AdmissionUpdateRequest } from '../../../models/api/patient-models';
import { AdmissionAddModel } from '../../../models/ui/patient-models';
import { PatientService } from '../../../services/patient.service';

@Component({
  selector: 'app-patient-admission',
  templateUrl: './patient-admission.component.html',
  styleUrls: ['./patient-admission.component.css']
})
export class PatientAdmissionComponent extends EditRecordBase implements OnInit, OnDestroy {

  dataModel = new AdmissionAddModel();
  routeSubscription: Subscription;
  splist: ServicepointListResponse[];
  constructor(
    private patientService: PatientService,
    private route: ActivatedRoute,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe,
  ) {
    super();
    this.iconCss = 'fa fa-user';
    this.pageTitle = 'Admission Detail';
  }

  ngOnInit() {
    this.createControls();
    this.showBackButton = false;
    // this.subTitle = this.translatePipe.transform('OPERATOR_ADD_MODE_TITLE');
    this.getServicepointList();
    this.routeSubscription = this.route.queryParams.subscribe(params => {
      if (params['addid']) {
        this.dataModel.admissionid = Number(params['addid']);
        this.recordState = EDITABLE_RECORD_STATE.UPDATE;
        this.setFormMode(FORM_MODE.VIEW);
        this.getAdmissionUpdates();
      } else {
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
    if (this.recordState === EDITABLE_RECORD_STATE.ADD) {
      const admissionAddRequest = new AdmissionAddRequest();
      this.dataModel.uuid = "9876";
      this.dataModel.patientid = this.patientService.patientid;
      this.dataModel.status = 1;
      this.dataModel.copyTo(admissionAddRequest);
      this.patientService.admissionAddPatient(admissionAddRequest).subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.dataModel.admissionid = payloadResponse.data.admissionid;
          this.patientService.setAdmissionId(payloadResponse.data.admissionid);
          this.appNotificationService.success();
          this.patientService.medicaldetialsid = payloadResponse.data.medicaldetailsid;
          this.patientService.admissionid = payloadResponse.data.admissionid;
          console.log('get id', this.patientService.medicaldetialsid);
          this.recordState = EDITABLE_RECORD_STATE.UPDATE;
          this.setFormMode(FORM_MODE.VIEW);

        }
        this.inProgress = false;
      });
    }
    else {
      // this.update();
      const admissionUpdateRequest = new AdmissionUpdateRequest();
      this.dataModel.uuid = "6767";
      this.dataModel.copyToUpdate(admissionUpdateRequest);
      this.patientService.updateAdmissionRequest(admissionUpdateRequest).subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          // this.patientService.medicaldetialsid = payloadResponse.data.medicaldetailsid;
          // console.log('get id',  this.patientService.medicaldetialsid);
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
}