import { Component, OnInit } from '@angular/core';
import { PROD_CODE } from 'app/app-constants';
import { HospitalListResponse, PatientFilterRequest } from 'app/models/api/hospital-models';
import { HospitalFilterModel } from 'app/models/ui/hospital-models';
import { HospitalService } from 'app/services/hospital.service';


@Component({
  selector: 'app-hospital-search',
  templateUrl: './hospital-search.component.html',
  styleUrls: ['./hospital-search.component.css']
})
export class HospitalSearchComponent implements OnInit {

  isExpanded = false;
  dataModel = new HospitalFilterModel();
  hospitallist: HospitalListResponse[] = [];
  constructor(public hospitalService: HospitalService) { }

  ngOnInit() {
    this.getServicepointList();
    this.dataModel.cpmid = 3;
  }

  getServicepointList() {
    const patientFilterRequest = new PatientFilterRequest();
    patientFilterRequest.prodcode = PROD_CODE;
    this.hospitalService.getServicepointList(patientFilterRequest).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.hospitallist = payloadResponse.data;
      }
    });
  }

  search() {
    this.isExpanded = false;
    const patientFilterRequest = new PatientFilterRequest();
    this.dataModel.copyTo(patientFilterRequest);
    this.hospitalService.cpmid = patientFilterRequest.cpmid;
    this.hospitalService.dataListSubjectTrigger(patientFilterRequest);
  }

  panelOpened() {
    this.isExpanded = true;
  }
}
