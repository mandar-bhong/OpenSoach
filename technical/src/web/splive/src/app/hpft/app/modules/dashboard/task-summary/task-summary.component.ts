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
  missedLabel = 'Missed';
  customColors = [
    {
      name: this.onTimeLabel,
      value: '#28a745'
    },
    {
      name: this.delayedLabel,
      value: '#ffc107'
    },
    {
      name: this.missedLabel,
      value: '#ff5252'
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
    // this.dashboardService.getTaskSummary().subscribe(payloadResponse => {
    //   if (payloadResponse && payloadResponse.issuccess) {
        this.tasksummary.copyFrom({ontime:100,delayed:11, missed:2});
        this.generateChartData();
    //   }
    // });
  }

  getTaskSummaryThisMonth() {
    // const dt = new Date();
    // const firstDayofMonth = new Date(dt.getFullYear(), dt.getMonth(), 1);
    // this.dashboardService.getTaskSummary(
    //   { spid: undefined, startdate: firstDayofMonth, enddate: dt }).subscribe(payloadResponse => {
    //     if (payloadResponse && payloadResponse.issuccess) {
    //       this.tasksummary.copyFrom(payloadResponse.data);
    //       this.generateChartData();
    //     }
    //   });

    this.tasksummary.copyFrom({ontime:12,delayed:2, missed:0});
        this.generateChartData();
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
      {
        name: this.missedLabel,
        value: this.tasksummary.missed
      },
    ];
  }

}
