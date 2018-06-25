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
  onTimeLabel = 'On Time';
  delayedLabel = 'Delayed';
  customColors = [
    {
      name: this.onTimeLabel,
      value: '#008000'
    },
    {
      name: this.delayedLabel,
      value: '#ffcc00'
    }
  ];

  taskchartdata = [];
  ngOnInit() {
    this.getTaskSummaryTillDate();
  }

  optionChange() {
    this.taskchartdata = [];
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
        this.generateChartData();
      }
    });
  }

  getTaskSummaryThisMonth() {
    const dt = new Date();
    const firstDayofMonth = new Date(dt.getFullYear(), dt.getMonth(), 1);
    this.dashboardService.getTaskSummary(
      { spid: undefined, startdate: firstDayofMonth, enddate: dt }).subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.tasksummary.copyFrom(payloadResponse.data);
          this.generateChartData();
        }
      });
  }

  generateChartData() {
    this.taskchartdata = [
      {
        name: this.onTimeLabel,
        value: this.tasksummary.ontime
      },
      {
        name: this.delayedLabel,
        value: this.tasksummary.delayed
      },
    ];
  }

}
