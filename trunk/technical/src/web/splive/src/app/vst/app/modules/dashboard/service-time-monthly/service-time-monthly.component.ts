import { Component, OnInit } from '@angular/core';
import { curveLinear } from 'd3-shape';

import { FeedbackTrendRequest, ServiceTimeAvrRequest } from '../../../models/api/dashboard-models';
import { FeedbackTrendModel, TrendChartPerMonthXaxis, SeriveTimeAvrModel } from '../../../models/ui/dashboard-models';
import { DashboardService } from '../../../services/dashboard.service';

@Component({
  selector: 'app-service-time-monthly',
  templateUrl: './service-time-monthly.component.html',
  styleUrls: ['./service-time-monthly.component.css']
})
export class ServiceTimeMonthlyComponent implements OnInit {

  curve = curveLinear;
  xAxisLabel = 'Time';
  yAxisLabel = 'Time (%)';
  feedbacktrenddata: SeriveTimeAvrModel[] = [];
  feedbacktrendchartdata = [];
  request = new ServiceTimeAvrRequest();
  timeline: TrendChartPerMonthXaxis[] = [];
  waittime = 'Wait time';
  creationtime = 'Job Creation';
  exectiontime = 'Job Execution';
  deliverytime = 'Delivery Time';
  legendTitle = 'Option';
  test = [];
  customColors = [
    {
      name: this.deliverytime,
      value: '#19915c'
    },
    {
      name: this.exectiontime,
      value: '#ffc107'
    },
    {
      name: this.creationtime,
      value: '#FF4C89'
    },
    {
      name: this.waittime,
      value: '#37A5CD'
    }
  ];
  data = [];
  constructor(private dashboardService: DashboardService) { }

  ngOnInit() {
    // this.getTrend();
    this.getSeviceTimeMonth();
  }
  getSeviceTimeMonth() {
    const currentDate = new Date();
    this.request.enddate = new Date(Date.UTC(
      currentDate.getFullYear(), currentDate.getMonth(), currentDate.getDate(), currentDate.getHours(), currentDate.getMinutes()));

    const ticks = Date.UTC(this.request.enddate.getUTCFullYear(), this.request.enddate.getUTCMonth() - 11, 1);
    this.request.startdate = new Date(ticks);

    this.dashboardService.getSeviceTimeMonth(this.request).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        payloadResponse.data.forEach(item => {
          const trendModel = new SeriveTimeAvrModel();
          trendModel.copyFrom(item);
          this.feedbacktrenddata.push(trendModel);
        });

        this.generateSeriesTimeline();
        this.generateRatingChartData();
      }
    });
  }

  // getTrend() {
  //   const currentDate = new Date();
  //   this.request.enddate = new Date(Date.UTC(
  //     currentDate.getFullYear(), currentDate.getMonth(), currentDate.getDate(), currentDate.getHours(), currentDate.getMinutes()));

  //   const ticks = Date.UTC(this.request.enddate.getUTCFullYear(), this.request.enddate.getUTCMonth() - 11, 1);
  //   this.request.startdate = new Date(ticks);

  //   this.generateSeriesTimeline();
  //   this.generateRatingChartData();

  //   this.timeline.forEach(item => {
  //     const xAxisDate = new Date(item.year, item.month).toUTCString();
  //     this.data.push({ name: xAxisDate, value: 10 + Math.floor(Math.random() * 100) });
  //   });
  // }

  generateRatingChartData() {
    const deliveryData = { name: this.deliverytime, series: [] };
    this.feedbacktrendchartdata.push(deliveryData);
    const exectionData = { name: this.exectiontime, series: [] };
    this.feedbacktrendchartdata.push(exectionData);
    const creationtData = { name: this.creationtime, series: [] };
    this.feedbacktrendchartdata.push(creationtData);
    const waitData = { name: this.waittime, series: [] };
    this.feedbacktrendchartdata.push(waitData);

    this.timeline.forEach(item => {
      const xAxisDate = new Date(item.year, item.month).toUTCString();
      const trendModel = this.feedbacktrenddata.find(rating => rating.year === item.year
        && rating.month - 1 === item.month);
      if (trendModel) {
        deliveryData.series.push({ name: xAxisDate, value: trendModel.deliverytime });
        exectionData.series.push({ name: xAxisDate, value: trendModel.jobexetime });
        creationtData.series.push({ name: xAxisDate, value: trendModel.jobcreationtime });
        waitData.series.push({ name: xAxisDate, value: trendModel.waittime });
      } else {
        deliveryData.series.push({ name: xAxisDate, value: 0 });
        exectionData.series.push({ name: xAxisDate, value: 0 });
        creationtData.series.push({ name: xAxisDate, value: 0 });
        waitData.series.push({ name: xAxisDate, value: 0 });
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

