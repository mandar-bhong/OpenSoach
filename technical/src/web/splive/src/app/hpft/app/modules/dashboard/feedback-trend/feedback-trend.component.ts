import { Component, OnInit } from '@angular/core';
import { curveLinear } from 'd3-shape';

import { FeedbackTrendRequest } from '../../../models/api/dashboard-models';
import { FeedbackTrendModel, TrendChartPerMonthXaxis } from '../../../models/ui/dashboard-models';
import { DashboardService } from '../../../services/dashboard.service';

@Component({
  selector: 'app-feedback-trend',
  templateUrl: './feedback-trend.component.html',
  styleUrls: ['./feedback-trend.component.css',
    '../default-dashboard/default-dashboard.component.css']
})

export class FeedbackTrendComponent implements OnInit {

  curve = curveLinear;
  xAxisLabel = 'Time';
  yAxisLabel = 'Rating (Count)';
  feedbacktrenddata: FeedbackTrendModel[] = [];
  feedbacktrendchartdata = [];
  request = new FeedbackTrendRequest();
  timeline: TrendChartPerMonthXaxis[] = [];
  rating1Label = 'Rating 1';
  rating2Label = 'Rating 2';
  rating3Label = 'Rating 3';
  rating4Label = 'Rating 4';
  rating5Label = 'Rating 5';
  legendTitle = 'Ratings';
  customColors = [
    {
      name: this.rating5Label,
      value: '#28a745'
    },
    {
      name: this.rating4Label,
      value: '#19915c'
    },
    {
      name: this.rating3Label,
      value: '#37A5CD'
    },
    {
      name: this.rating2Label,
      value: '#ffc107'
    },
    {
      name: this.rating1Label,
      value: '#dc3545'
    }
  ];

  constructor(private dashboardService: DashboardService) { }

  ngOnInit() {
    this.getFeedbackTrend();
  }

  getFeedbackTrend() {
    const currentDate = new Date();
    this.request.enddate = new Date(Date.UTC(
      currentDate.getFullYear(), currentDate.getMonth(), currentDate.getDate(), currentDate.getHours(), currentDate.getMinutes()));

    const ticks = Date.UTC(this.request.enddate.getUTCFullYear(), this.request.enddate.getUTCMonth() - 11, 1);
    this.request.startdate = new Date(ticks);

    this.dashboardService.getFeedbackTrend(this.request).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        payloadResponse.data.forEach(item => {
          const trendModel = new FeedbackTrendModel();
          trendModel.copyFrom(item);
          this.feedbacktrenddata.push(trendModel);
        });

        this.generateSeriesTimeline();
        this.generateRatingChartData();
      }
    });
  }



  generateRatingChartData() {
    const rating5Data = { name: this.rating5Label, series: [] };
    this.feedbacktrendchartdata.push(rating5Data);
    const rating4Data = { name: this.rating4Label, series: [] };
    this.feedbacktrendchartdata.push(rating4Data);
    const rating3Data = { name: this.rating3Label, series: [] };
    this.feedbacktrendchartdata.push(rating3Data);
    const rating2Data = { name: this.rating2Label, series: [] };
    this.feedbacktrendchartdata.push(rating2Data);
    const rating1Data = { name: this.rating1Label, series: [] };
    this.feedbacktrendchartdata.push(rating1Data);

    this.timeline.forEach(item => {
      const xAxisDate = new Date(item.year, item.month).toUTCString();
      const trendModel = this.feedbacktrenddata.find(rating => rating.year === item.year
        && rating.month === item.month - 1);

      if (trendModel) {
        rating5Data.series.push({ name: xAxisDate, value: trendModel.rating5 });
        rating4Data.series.push({ name: xAxisDate, value: trendModel.rating4 });
        rating3Data.series.push({ name: xAxisDate, value: trendModel.rating3 });
        rating2Data.series.push({ name: xAxisDate, value: trendModel.rating2 });
        rating1Data.series.push({ name: xAxisDate, value: trendModel.rating1 });
      } else {
        rating5Data.series.push({ name: xAxisDate, value: 0 });
        rating4Data.series.push({ name: xAxisDate, value: 0 });
        rating3Data.series.push({ name: xAxisDate, value: 0 });
        rating2Data.series.push({ name: xAxisDate, value: 0 });
        rating1Data.series.push({ name: xAxisDate, value: 0 });
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
