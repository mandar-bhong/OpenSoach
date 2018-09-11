import { Component, OnInit } from '@angular/core';
import { curveLinear } from 'd3-shape';

import { FeedbackTrendRequest } from '../../../models/api/dashboard-models';
import { FeedbackTrendModel, TrendChartPerMonthXaxis } from '../../../models/ui/dashboard-models';
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
  feedbacktrenddata: FeedbackTrendModel[] = [];
  feedbacktrendchartdata = [];
  request = new FeedbackTrendRequest();
  timeline: TrendChartPerMonthXaxis[] = [];
  rating1Label = 'Wait time';
  rating2Label = 'Job Creation';
  rating3Label = 'Job Execution';
  rating4Label = 'Delivery Time';
  // rating5Label = 'Rating 5';
  legendTitle = 'Option';
  test = [];
  customColors = [
    // {
    //   name: this.rating5Label,
    //   value: '#28a745'
    // },
    {
      name: this.rating4Label,
      value: '#19915c'
    },
    {
      name: this.rating3Label,
      value: '#ffc107'
    },
    {
      name: this.rating2Label,
      value: '#FF4C89'
    },
    {
      name: this.rating1Label,
      value: '#37A5CD'
    }
  ];
  // data: any;
  data = [];
  constructor(private dashboardService: DashboardService) { }

  ngOnInit() {
    this.getTrend();
  }
  // getFeedbackTrend() {
  //   const currentDate = new Date();
  //   this.request.enddate = new Date(Date.UTC(
  //     currentDate.getFullYear(), currentDate.getMonth(), currentDate.getDate(), currentDate.getHours(), currentDate.getMinutes()));

  //   const ticks = Date.UTC(this.request.enddate.getUTCFullYear(), this.request.enddate.getUTCMonth() - 11, 1);
  //   this.request.startdate = new Date(ticks);

  //   this.dashboardService.getFeedbackTrend(this.request).subscribe(payloadResponse => {
  //     if (payloadResponse && payloadResponse.issuccess) {
  //       payloadResponse.data.forEach(item => {
  //         const trendModel = new FeedbackTrendModel();
  //         trendModel.copyFrom(item);
  //         this.feedbacktrenddata.push(trendModel);
  //       });

  //       this.generateSeriesTimeline();
  //       this.generateRatingChartData();
  //     }
  //   });
  // }

  getTrend() {
    const currentDate = new Date();
    this.request.enddate = new Date(Date.UTC(
      currentDate.getFullYear(), currentDate.getMonth(), currentDate.getDate(), currentDate.getHours(), currentDate.getMinutes()));

    const ticks = Date.UTC(this.request.enddate.getUTCFullYear(), this.request.enddate.getUTCMonth() - 11, 1);
    this.request.startdate = new Date(ticks);

    this.generateSeriesTimeline();
    this.generateRatingChartData();

    this.timeline.forEach(item => {
      const xAxisDate = new Date(item.year, item.month).toUTCString();
      this.data.push({ name: xAxisDate, value: 10 + Math.floor(Math.random() * 100) });
    });
  }

  generateRatingChartData() {
    // const rating5Data = { name: this.rating5Label, series: [] };
    // this.feedbacktrendchartdata.push(rating5Data);
    const rating4Data = { name: this.rating4Label, series: [] };
    this.feedbacktrendchartdata.push(rating4Data);
    const rating3Data = { name: this.rating3Label, series: [] };
    this.feedbacktrendchartdata.push(rating3Data);
    const rating2Data = { name: this.rating2Label, series: [] };
    this.feedbacktrendchartdata.push(rating2Data);
    const rating1Data = { name: this.rating1Label, series: [] };
    this.feedbacktrendchartdata.push(rating1Data);

    const rating4 = 100;
    const rating3 = 35;
    const rating2 = 50;
    const rating1 = 20;
    this.timeline.forEach(item => {
      const xAxisDate = new Date(item.year, item.month).toUTCString();
      const trendModel = this.feedbacktrenddata.find(rating => rating.year === item.year
        && rating.month - 1 === item.month);
      // if (trendModel) {
      //   // rating5Data.series.push({ name: xAxisDate, value: trendModel.rating5 });
      //   rating4Data.series.push({ name: xAxisDate, value: trendModel.rating4 });
      //   rating3Data.series.push({ name: xAxisDate, value: trendModel.rating3 });
      //   rating2Data.series.push({ name: xAxisDate, value: trendModel.rating2 });
      //   rating1Data.series.push({ name: xAxisDate, value: trendModel.rating1 });
      // } else {
      //   // rating5Data.series.push({ name: xAxisDate, value: 0 });
      //   rating4Data.series.push({ name: xAxisDate, value: 0 });
      //   rating3Data.series.push({ name: xAxisDate, value: 0 });
      //   rating2Data.series.push({ name: xAxisDate, value: 0 });
      //   rating1Data.series.push({ name: xAxisDate, value: 0 });
      // }
      rating4Data.series.push({ name: xAxisDate, value: rating4 });
      rating3Data.series.push({ name: xAxisDate, value: rating3 });
      rating2Data.series.push({ name: xAxisDate, value: rating2 });
      rating1Data.series.push({ name: xAxisDate, value: rating1 });

      // rating4 = rating4 + 12;
      // dummyDelayed = dummyDelayed - 13;
      // dummyMissed = dummyMissed - 3;

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
// this.test = [
//   {
//     'name': 'Wait time',
//     'series': [
//       {
//         value: 35,
//         name: 'Sat, 30 Sep 2017 18:30:00 GMT'
//       },
//       {
//         value: 48,
//         name: 'Tue, 31 Oct 2017 18:30:00 GMT'
//       },
//       {
//         value: 20,
//         name: 'Thu, 30 Nov 2017 18:30:00 GMT'
//       },
//       {
//         value: 56,
//         name: 'Sun, 31 Dec 2017 18:30:00 GMT'
//       },
//       {
//         value: 95,
//         name: 'Wed, 31 Jan 2018 18:30:00 GMT'
//       },
//       {
//         name: 'Wed, 28 Feb 2018 18:30:00 GMT',
//         value: 54
//       },
//       {
//         name: 'Sat, 31 Mar 2018 18:30:00 GMT',
//         value: 54
//       }
//     ]
//   }
// ];
