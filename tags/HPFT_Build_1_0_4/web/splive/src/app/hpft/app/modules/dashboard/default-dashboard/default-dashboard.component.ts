import { Component, OnInit } from '@angular/core';

import { ComplaintSummaryModel, DeviceSummaryModel, ServicePointSummaryModel, PatientSummaryModel } from '../../../models/ui/dashboard-models';
import { DashboardService } from '../../../services/dashboard.service';

@Component({
  selector: 'app-default-dashboard',
  templateUrl: './default-dashboard.component.html',
  styleUrls: ['./default-dashboard.component.css']
})
export class DefaultDashboardComponent implements OnInit {

  constructor(private dashboardService: DashboardService) { }
  devicesummary = new DeviceSummaryModel();
  spsummary = new ServicePointSummaryModel();
  patientsummary = new PatientSummaryModel();
  ngOnInit() {
    this.getDeviceSummary();
    this.getServicePointSummary();    
    this.getPatientSummary();
  }

  getDeviceSummary() {
    this.dashboardService.getDeviceSummary().subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.devicesummary.copyFrom(payloadResponse.data);
      }
    });
  }

  getServicePointSummary() {
    this.dashboardService.getServicePointSummary().subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.spsummary.copyFrom(payloadResponse.data);
      }
    });
  } 

  getPatientSummary() {
    this.dashboardService.getPatientSummary().subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.patientsummary.copyFrom(payloadResponse.data);
      }
    });
  }
}
