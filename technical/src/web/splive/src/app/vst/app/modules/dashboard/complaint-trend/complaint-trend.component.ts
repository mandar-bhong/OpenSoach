import { Component, OnInit } from '@angular/core';
import { curveLinear } from 'd3-shape';

import { ComplaintTrendRequest } from '../../../models/api/dashboard-models';
import { ComplaintTrendModel, TrendChartPerMonthXaxis } from '../../../models/ui/dashboard-models';
import { DashboardService } from '../../../services/dashboard.service';

@Component({
  selector: 'app-complaint-trend',
  templateUrl: './complaint-trend.component.html',
  styleUrls: ['./complaint-trend.component.css',
    '../default-dashboard/default-dashboard.component.css']
})
export class ComplaintTrendComponent implements OnInit {

  curve = curveLinear;
  xAxisLabel = 'Time';
  yAxisLabel = 'Complaint (Count)';
  complainttrenddata: ComplaintTrendModel[] = [];
  complainttrendchartdata = [];
  request = new ComplaintTrendRequest();
  timeline: TrendChartPerMonthXaxis[] = [];
  openLabel = 'Open';
  inprogressLabel = 'In Progress';
  closedLabel = 'Closed';
  legendTitle = 'Status';
  customColors = [
    {
      name: this.openLabel,
      value: '#dc3545'
    },
    {
      name: this.inprogressLabel,
      value: '#37A5CD'
    },
    {
      name: this.closedLabel,
      value: '#28a745'
    }
  ];

  constructor(private dashboardService: DashboardService) { }

  ngOnInit() {
    this.getComplaintTrend();
  }

  getComplaintTrend() {
    const currentDate = new Date();
    this.request.enddate = new Date(Date.UTC(
      currentDate.getFullYear(), currentDate.getMonth(), currentDate.getDate(), currentDate.getHours(), currentDate.getMinutes()));

    const ticks = Date.UTC(this.request.enddate.getUTCFullYear(), this.request.enddate.getUTCMonth() - 11, 1);
    this.request.startdate = new Date(ticks);

    this.dashboardService.getComplaintTrend(this.request).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        payloadResponse.data.forEach(item => {
          const trendModel = new ComplaintTrendModel();
          trendModel.copyFrom(item);
          this.complainttrenddata.push(trendModel);
        });

        this.generateSeriesTimeline();
        this.generateRatingChartData();
      }
    });
  }



  generateRatingChartData() {
    const openData = { name: this.openLabel, series: [] };
    this.complainttrendchartdata.push(openData);
    const inprogressData = { name: this.inprogressLabel, series: [] };
    this.complainttrendchartdata.push(inprogressData);
    const closedData = { name: this.closedLabel, series: [] };
    this.complainttrendchartdata.push(closedData);

    this.timeline.forEach(item => {
      const xAxisDate = new Date(item.year, item.month).toUTCString();
      const trendModel = this.complainttrenddata.find(rating => rating.year === item.year
        && rating.month - 1 === item.month);

      if (trendModel) {
        openData.series.push({ name: xAxisDate, value: trendModel.open });
        inprogressData.series.push({ name: xAxisDate, value: trendModel.inprogress });
        closedData.series.push({ name: xAxisDate, value: trendModel.closed });
      } else {
        openData.series.push({ name: xAxisDate, value: 0 });
        inprogressData.series.push({ name: xAxisDate, value: 0 });
        closedData.series.push({ name: xAxisDate, value: 0 });
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
