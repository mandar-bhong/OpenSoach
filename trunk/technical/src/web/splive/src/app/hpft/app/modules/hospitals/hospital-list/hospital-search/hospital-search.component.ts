import { Component, OnInit } from '@angular/core';
import { PatientFilterRequest, HospitalListResponse, HospitalSearchRequest } from 'app/models/api/hospital-models';
import { HospitalFilterModel } from 'app/models/ui/hospital-models';
import { HospitalService } from 'app/services/hospital.service';
import { PROD_CODE } from 'app/app-constants';


@Component({
  selector: 'app-hospital-search',
  templateUrl: './hospital-search.component.html',
  styleUrls: ['./hospital-search.component.css']
})
export class HospitalSearchComponent implements OnInit {

  isExpanded = false;
  dataModel = new HospitalFilterModel();
  hospitallist: HospitalListResponse[] = [];
  hsopitalName: any;
  getCpmid: number;
  constructor(public hospitalService: HospitalService) { }

  ngOnInit() {
    this.getServicepointList();
    this.dataModel.cpmid = 3;
  }


  // getServicepointList() {
  //   this.hospitalService.getServicepointList().subscribe(payloadResponse => {
  //     if (payloadResponse && payloadResponse.issuccess) {
  //       this.hospitallist = payloadResponse.data;
  //       console.log("this.hospitallist",this.hospitallist);
  //     }
  //   });
  // }

  getServicepointList() {
    // const hospitalSearchRequest = new HospitalSearchRequest();
    // hospitalSearchRequest.prodcode = PROD_CODE;

    const patientFilterRequest = new PatientFilterRequest();
    patientFilterRequest.prodcode = PROD_CODE;
    this.hospitalService.getServicepointList(patientFilterRequest).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.hospitallist = payloadResponse.data;
        // this.hospitalService.cpmid =
        console.log("this.hospitallist", this.hospitallist);
        
      }
    });
  }

  search() {
    this.isExpanded = false;
    const patientFilterRequest = new PatientFilterRequest();
    this.dataModel.copyTo(patientFilterRequest);
    console.log('patientFilterRequest', patientFilterRequest);
    this.hospitalService.cpmid = patientFilterRequest.cpmid;
    this.hospitalService.dataListSubjectTrigger(patientFilterRequest);
  }

  panelOpened() {
    this.isExpanded = true;
  }
}
