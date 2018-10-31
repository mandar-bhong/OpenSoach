import { Component, OnInit } from '@angular/core';
import { curveLinear } from 'd3-shape';

import { TaskTrendRequest, VehicleServiceTrendWeeklyRequest } from '../../../models/api/dashboard-models';
import {
  TaskTrendModel, TrendChartPerMonthXaxis, VehicleServiceTrendWeeklyModel,
} from '../../../models/ui/dashboard-models';
import { DashboardService } from '../../../services/dashboard.service';

@Component({
  selector: 'app-vehicle-weekly-trend',
  templateUrl: './vehicle-weekly-trend.component.html',
  styleUrls: ['./vehicle-weekly-trend.component.css',
    '../default-dashboard/default-dashboard.component.css']
})
export class VehicleWeeklyTrendComponent implements OnInit {

  curve = curveLinear;
  xAxisLabel = 'Vehicles';
  yAxisLabel = 'Days';
  tasktrenddata: VehicleServiceTrendWeeklyModel[] = [];
  tasktrendchartdata = [];
  a = [];
  b = [];
  request = new VehicleServiceTrendWeeklyRequest();
  timeline: TrendChartPerMonthXaxis[] = [];
  // timeday: VehicleChartPerWeeklyYaxis[] = [];
  ontimeLabel = 'Vehicles';
  // delayedLabel = 'Delayed';
  // missedLabel = 'Missed';
  // legendTitle = 'Task Status';
  formatXAxis;
  // formatYAxis;
  customColors = [
    {
      name: 'SUN',
      value: '#245AAE'
    },
    {
      name: 'MON',
      value: '#E466C9'
    },
    {
      name: 'TUE',
      value: '#956EE8'
    },
    {
      name: 'WED',
      value: '#00B1EA'
    },
    {
      name: 'THU',
      value: '#00DDC6'
    },
    {
      name: 'FRI',
      value: '#FFB467'
    },
    {
      name: 'SAT',
      value: '#FF6859'
    }
  ];
  data: any;
  constructor(private dashboardService: DashboardService) { }

  ngOnInit() {
    // this.data = [
    //   {
    //     'name': 'MON',
    //     'value': 82
    //   },
    //   {
    //     'name': 'TUE',
    //     'value': 85
    //   },
    //   {
    //     'name': 'WED',
    //     'value': 60
    //   },
    //   {
    //     'name': 'THU',
    //     'value': 70
    //   },
    //   {
    //     'name': 'FRI',
    //     'value': 80
    //   },
    //   {
    //     'name': 'SAT',
    //     'value': 11
    //   },
    //   {
    //     'name': 'SUN',
    //     'value': 100
    //   }
    // ];
    this.getTaskTrend();
  }

  getTaskTrend() {
    const currentDate = new Date();
    this.request.enddate = new Date(Date.UTC(
      currentDate.getFullYear(), currentDate.getMonth(), currentDate.getDate(), currentDate.getHours(), currentDate.getMinutes()));
    console.log('this.request.enddate', this.request.enddate);

    const ticks = Date.UTC(this.request.enddate.getUTCFullYear(), this.request.enddate.getUTCMonth(),
      this.request.enddate.getUTCDate() - 6);
    console.log('this.request.enddate.getUTCDay()', this.request.enddate.getUTCDay() - 6);
    this.request.startdate = new Date(ticks);
    console.log('ticks', ticks);
    console.log('this.request.startdate ', this.request.startdate);

    this.dashboardService.getVehicleServiceTrendWeekly(this.request).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        payloadResponse.data.forEach(item => {
          const trendModel = new VehicleServiceTrendWeeklyModel();
          trendModel.copyFrom(item);
          this.tasktrenddata.push(trendModel);
        });
        console.log('data received');
        // this.generateSeriesTimeline();
        this.generateRatingChartData();
      }
    });
  }



  generateRatingChartData() {
    // this.timeline.forEach(item => {
    // const xAxisDate = new Date(item.year, item.month).toUTCString();
    // const trendModel = this.tasktrenddata.find(rating => rating.year === item.year
    //   && rating.month-1 === item.month);
    // if (trendModel) {
    //   this.tasktrendchartdata.push({ name: xAxisDate, value: trendModel.vehicleserviced });
    // } else {
    //  this.tasktrendchartdata.push({ name: xAxisDate, value: trendModel.vehicleserviced });
    // }
    // });




    const start = this.request.startdate;
    const end = this.request.enddate;

    this.tasktrenddata.forEach(a => {
      const ab = this.tasktrenddata;
      if (ab) {
        this.tasktrendchartdata.push({ name: a.servicedate, value: a.vehicleserviced });
        // console.log('a', this.b);
        console.log('this.tasktrendchartdata check', this.tasktrendchartdata);
      } else {
        this.tasktrendchartdata.push({ name: a.servicedate, value: 0 });
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

  formatYAxis(value: string) {
    const date = new Date(value);
    return date.toLocaleDateString('en-US', { weekday: 'short' });
  }
}
