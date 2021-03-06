import { Component, OnInit } from '@angular/core';

import { ServicepointListResponse } from '../../../../../../prod-shared/models/api/servicepoint-models';
import { PatientFilterRequest } from '../../../../models/api/patient-data-models';
import { PatientFilterModel } from '../../../../models/ui/patient-models';
import { PatientService } from '../../../../services/patient.service';
import { PATIENT_STATE } from 'app/app-constants';

@Component({
  selector: 'app-patient-search',
  templateUrl: './patient-search.component.html',
  styleUrls: ['./patient-search.component.css']
})
export class PatientSearchComponent implements OnInit {
  dataModel = new PatientFilterModel();
  isExpanded = false;
  splist: ServicepointListResponse[] = [];
  PATIENT_STATE:PATIENT_STATE;

  patientFilterRequest: PatientFilterRequest;
  constructor(public patientService: PatientService) { }

  ngOnInit() {
    this.getServicepointList();
    this.dataModel.status = PATIENT_STATE.HOSPITALIZE;
  }
  // Accept data from ward ie. list of ward
  getServicepointList() {
    this.patientService.getServicepointList().subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.splist = payloadResponse.data;
      }
    });
  }

  search() {
    this.isExpanded = false;
    const patientFilterRequest = new PatientFilterRequest();
    this.dataModel.copyTo(patientFilterRequest);
    this.patientService.dataListSubjectTrigger(patientFilterRequest);

  }

  panelOpened() {
    this.isExpanded = true;
  }
}