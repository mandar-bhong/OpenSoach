import { Component, OnInit } from '@angular/core';

import { ComplaintSummaryModel } from '../../../models/ui/dashboard-models';
import { DashboardService } from '../../../services/dashboard.service';

@Component({
  selector: 'app-complaint-summary',
  templateUrl: './complaint-summary.component.html',
  styleUrls: ['./complaint-summary.component.css',
    '../default-dashboard/default-dashboard.component.css']
})
export class ComplaintSummaryComponent implements OnInit {

  complaintsummary = new ComplaintSummaryModel();

  selectedoption = '0';
  constructor(private dashboardService: DashboardService) { }
  openLabel = 'Open';
  inprogressLabel = 'In Progress';
  closedLabel = 'Closed';
  customColors = [
    {
      name: this.openLabel,
      value: '#cc3300'
    },
    {
      name: this.inprogressLabel,
      value: '#37A5CD'
    },
    {
      name: this.closedLabel,
      value: '#008000'
    }
  ];

  complaintchartdata = [];
  ngOnInit() {
    this.getComplaintSummaryTillDate();
  }

  optionChange() {
    this.complaintchartdata = [];
    if (this.selectedoption === '1') {
      this.getComplaintSummaryThisMonth();
    } else {
      this.getComplaintSummaryTillDate();
    }
  }

  getComplaintSummaryTillDate() {
    this.dashboardService.getComplaintSummary().subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.complaintsummary.copyFrom(payloadResponse.data);
        console.log('complaint summary till date', this.complaintsummary);
        this.generateChartData();
      }
    });
  }

  getComplaintSummaryThisMonth() {
    const dt = new Date();

    const firstDayofMonth = new Date(dt.getFullYear(), dt.getMonth(), 1);
    const lastDayofMonth = new Date(dt.getFullYear(), dt.getMonth() + 1, 1);

    this.dashboardService.getComplaintSummary(
      { spid: undefined, startdate: firstDayofMonth, enddate: lastDayofMonth }).subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.complaintsummary.copyFrom(payloadResponse.data);
          console.log('complaint summary month', this.complaintsummary);
          this.generateChartData();
        }
      });
  }

  generateChartData() {
    this.complaintchartdata = [
      {
        name: this.openLabel,
        value: this.complaintsummary.open
      },
      {
        name: this.inprogressLabel,
        value: this.complaintsummary.inprogress
      },
      {
        name: this.closedLabel,
        value: this.complaintsummary.closed
      },
    ];
  }
}
