import { Component, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs';

import { EnumDataSourceItem } from '../../../../../shared/models/ui/enum-datasource-item';
import { TranslatePipe } from '../../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';
import { EditRecordBase } from '../../../../../shared/views/edit-record-base';
import {
  MedicalDetailAddRequest,
  PatientDataAddRequest,
  PatientDetailAddRequest,
} from '../../../models/api/patient-models';
import { PatientDataModel } from '../../../models/ui/patient-models';
import { PatientService } from '../../../services/patient.service';
import { ServicepointListResponse } from '../../../../../prod-shared/models/api/servicepoint-models';
import { ServicepointConfigureListResponse } from '../../../../../prod-shared/models/api/service-configuration-models';

@Component({
  selector: 'app-patient-add',
  templateUrl: './patient-add.component.html',
  styleUrls: ['./patient-add.component.css']
})
export class PatientAddComponent extends EditRecordBase implements OnInit, OnDestroy {
  dataModel = new PatientDataModel();
  routeSubscription: Subscription;
  patientStates: EnumDataSourceItem<number>[];
  splist: ServicepointListResponse[] = [];
  spconfigures: ServicepointConfigureListResponse[] = [];
  constructor(
    private patientService: PatientService,
    private route: ActivatedRoute,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe,
  ) {
    super();
  }

  ngOnInit() {
    this.getServicepointList();
    this.getServicepointConfigureList();
    this.patientStates = this.patientService.getPatientStates();
    this.dataModel.patientdetails = new PatientDetailAddRequest();
    this.dataModel.medicaldetails = new MedicalDetailAddRequest();
    // this.dataModel.patientfiletemplate = new PatientFileAddRequest();
    this.routeSubscription = this.route.queryParams.subscribe(params => {
      if (params['id']) {
        this.dataModel.patientid = Number(params['id']);
        // this.getPatientDetails();
      }
      this.callbackUrl = params['callbackurl'];
    });
  }
  getPatientDetails() {
    this.patientService.getPatientDetails({ recid: this.dataModel.patientid }).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        if (payloadResponse.data) {
          this.dataModel.copyFrom(payloadResponse.data);
        }
      }
    });
  }
  getServicepointList() {
    this.patientService.getServicepointList().subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.splist = payloadResponse.data;
      }
    });
  }
  getServicepointConfigureList() {
    // get Existing servicepoint configure short list
    this.patientService.getServicepointConfigureList().subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.spconfigures = payloadResponse.data;
      }
    });
  }
  save() {
    const patientDataAddRequest = new PatientDataAddRequest();
    this.dataModel.copyTo(patientDataAddRequest);
    this.patientService.addPatient(patientDataAddRequest).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.appNotificationService.success();
        // this.dataModel. = payloadResponse.data;
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
}
