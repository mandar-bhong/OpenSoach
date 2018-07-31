import { Component, OnInit } from '@angular/core';
import { curveLinear } from 'd3-shape';

import { TaskTrendRequest } from '../../../models/api/dashboard-models';
import { TaskTrendModel, TrendChartPerMonthXaxis } from '../../../models/ui/dashboard-models';
import { DashboardService } from '../../../services/dashboard.service';

@Component({
  selector: 'app-task-trend',
  templateUrl: './task-trend.component.html',
  styleUrls: ['./task-trend.component.css',
    '../default-dashboard/default-dashboard.component.css']
})
export class TaskTrendComponent implements OnInit {

  curve = curveLinear;
  xAxisLabel = 'Time';
  yAxisLabel = 'Task (Count)';
  tasktrenddata: TaskTrendModel[] = [];
  tasktrendchartdata = [];
  request = new TaskTrendRequest();
  timeline: TrendChartPerMonthXaxis[] = [];
  ontimeLabel = 'On Time';
  delayedLabel = 'Delayed';
  legendTitle = 'Task Status';
  customColors = [
    {
      name: this.ontimeLabel,
      value: '#28a745'
    },
    {
      name: this.delayedLabel,
      value: '#ffc107'
    },
  ];

  constructor(private dashboardService: DashboardService) { }

  ngOnInit() {
    this.getTaskTrend();
  }

  getTaskTrend() {
    const currentDate = new Date();
    this.request.enddate = new Date(Date.UTC(
      currentDate.getFullYear(), currentDate.getMonth(), currentDate.getDate(), currentDate.getHours(), currentDate.getMinutes()));

    const ticks = Date.UTC(this.request.enddate.getUTCFullYear(), this.request.enddate.getUTCMonth() - 11, 1);
    this.request.startdate = new Date(ticks);

    this.dashboardService.getTaskTrend(this.request).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        payloadResponse.data.forEach(item => {
          const trendModel = new TaskTrendModel();
          trendModel.copyFrom(item);
          this.tasktrenddata.push(trendModel);
        });

        this.generateSeriesTimeline();
        this.generateRatingChartData();
      }
    });
  }



  generateRatingChartData() {
    const ontimeData = { name: this.ontimeLabel, series: [] };
    this.tasktrendchartdata.push(ontimeData);
    const delayedData = { name: this.delayedLabel, series: [] };
    this.tasktrendchartdata.push(delayedData);

    this.timeline.forEach(item => {
      const xAxisDate = new Date(item.year, item.month).toUTCString();
      const trendModel = this.tasktrenddata.find(rating => rating.year === item.year
        && rating.month === item.month - 1);

      if (trendModel) {
        ontimeData.series.push({ name: xAxisDate, value: trendModel.ontime });
        delayedData.series.push({ name: xAxisDate, value: trendModel.delayed });
      } else {
        ontimeData.series.push({ name: xAxisDate, value: 0 });
        delayedData.series.push({ name: xAxisDate, value: 0 });
      }
    });
  }

  generateSeriesTimeline() {
    let month = this.request.startdate.getMonth();
    let year = this.request.startdate.getFullYear();
    for (let i = 0; i < 12; i++) {
      this.timeline.push({ year: year, month: month });
      month = month + 1;
      if (month > 11) {
        month = 0;
        year = year + 1;
      }
    }
  }

  formatXAxis(value: string) {
    const date = new Date(value);
    return date.toLocaleDateString('en-US', { month: 'short' });
  }
}
