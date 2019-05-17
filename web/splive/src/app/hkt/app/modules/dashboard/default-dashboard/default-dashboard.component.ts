import { Component, OnInit } from '@angular/core';

import { ComplaintSummaryModel, DeviceSummaryModel, ServicePointSummaryModel } from '../../../models/ui/dashboard-models';
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
  complaintsummary = new ComplaintSummaryModel();
  ngOnInit() {
    this.getDeviceSummary();
    this.getServicePointSummary();
    this.getComplaintSummary();
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

  getComplaintSummary() {
    this.dashboardService.getComplaintSummary().subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.complaintsummary.copyFrom(payloadResponse.data);
      }
    });
  }
}
