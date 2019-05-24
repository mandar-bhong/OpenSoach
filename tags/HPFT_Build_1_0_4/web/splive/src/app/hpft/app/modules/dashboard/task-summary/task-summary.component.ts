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
  taskchartdata = [];  // define its type
  // take it from constants
  hospitalizedLabel = 'Hospitalized';
  dischargedLabel = 'Discharged';

  constructor(private dashboardService: DashboardService) { }
  

  customColors = [
    {
      name: this.hospitalizedLabel,
      value: '#28a745'
    },
    {
      name: this.dischargedLabel,
      value: '#52A6CA'
    }
  ]; 

  ngOnInit() {
    this.getTaskSummaryTillDate();
    // this.optionChange();
  }

  optionChange() {
    this.taskchartdata = [];
    // take it from constnats
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
    // check why spid is undefined.
    this.dashboardService.getTaskSummary({ spid: undefined, startdate: firstDayofMonth, enddate: dt }).subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.tasksummary.copyFrom(payloadResponse.data);
          this.generateChartData();
        }
      });
  }

  generateChartData() {
    this.taskchartdata = [
      {
        name: this.hospitalizedLabel,
        value: this.tasksummary.admitted
      },
      {
        name: this.dischargedLabel,
        value: this.tasksummary.discharged
      },
    ];
  }

}
