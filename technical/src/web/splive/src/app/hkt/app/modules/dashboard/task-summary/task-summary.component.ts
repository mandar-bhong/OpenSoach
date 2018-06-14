import { Component, OnInit } from '@angular/core';

import { TaskSummaryModel } from '../../../models/ui/dashboard-models';
import { DashboardService } from '../../../services/dashboard.service';

@Component({
  selector: 'app-task-summary',
  templateUrl: './task-summary.component.html',
  styleUrls: ['./task-summary.component.css',
    '../default-dashboard/default-dashboard.component.css']
})
export class TaskSummaryComponent implements OnInit {

  tasksummary = new TaskSummaryModel();

  selectedoption = '0';
  constructor(private dashboardService: DashboardService) { }

  ngOnInit() {
    this.getTaskSummaryTillDate();
  }

  optionChange() {
    if (this.selectedoption === '1') {
      this.getTaskSummaryThisMonth();
    } else {
      this.getTaskSummaryTillDate();
    }
  }

  getTaskSummaryTillDate() {
    this.dashboardService.getTaskSummary().subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.tasksummary.copyFrom(payloadResponse.data);
        console.log('task summary till date', this.tasksummary);
      }
    });
  }

  getTaskSummaryThisMonth() {
    const dt = new Date();

    const firstDayofMonth = new Date(dt.getFullYear(), dt.getMonth(), 1);
    const lastDayofMonth = new Date(dt.getFullYear(), dt.getMonth() + 1, 1);

    this.dashboardService.getTaskSummary(
      { spid: undefined, startdate: firstDayofMonth, enddate: lastDayofMonth }).subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.tasksummary.copyFrom(payloadResponse.data);
          console.log('task summary month', this.tasksummary);
        }
      });
  }

}
