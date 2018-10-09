import { Component, OnInit } from '@angular/core';
import { curveLinear } from 'd3-shape';

import { TaskTrendRequest, VehicleServiceTrendMontlyRequest } from '../../../models/api/dashboard-models';
import {
  TaskTrendModel, TrendChartPerMonthXaxis, SeriveTimeAvrModel,
  VehicleServiceTrendMonthlyModel
} from '../../../models/ui/dashboard-models';
import { DashboardService } from '../../../services/dashboard.service';

@Component({
  selector: 'app-vehicle-monthly-trend',
  templateUrl: './vehicle-monthly-trend.component.html',
  styleUrls: ['./vehicle-monthly-trend.component.css',
    '../default-dashboard/default-dashboard.component.css']
})
export class VehicleMonthlyTrendComponent implements OnInit {

  curve = curveLinear;
  xAxisLabel = 'Time';
  yAxisLabel = 'Vehicles';
  vehicletrenddata: VehicleServiceTrendMonthlyModel[] = [];
  tasktrendchartdata = [];
  request = new VehicleServiceTrendMontlyRequest();
  timeline: TrendChartPerMonthXaxis[] = [];
  ontimeLabel = 'Vehicles';
  colorScheme = {
    domain: ['#E466C9', '#956EE8', '#00B1EA', '#00DDC6', '#FFB467', '#FF6859', '#245AAE']
  };
  data = [];
  constructor(private dashboardService: DashboardService) { }

  ngOnInit() {
    this.getSeviceTimeMonth();
    // this.getTrend();
  }
  getSeviceTimeMonth() {
    const currentDate = new Date();
    this.request.enddate = new Date(Date.UTC(
      currentDate.getFullYear(), currentDate.getMonth(), currentDate.getDate(), currentDate.getHours(), currentDate.getMinutes()));

    const ticks = Date.UTC(this.request.enddate.getUTCFullYear(), this.request.enddate.getUTCMonth() - 11, 1);
    this.request.startdate = new Date(ticks);

    this.dashboardService.getVehicleServiceTrendMontly(this.request).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        payloadResponse.data.forEach(item => {
          const trendModel = new VehicleServiceTrendMonthlyModel();
          trendModel.copyFrom(item);
          this.vehicletrenddata.push(trendModel);
          console.log('this.tasktrenddata testing');
          console.log(this.vehicletrenddata);
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

  //   this.timeline.forEach(item => {
  //     const xAxisDate = new Date(item.year, item.month).toUTCString();
  //     this.data.push({ name: xAxisDate, value: 3000 + Math.floor(Math.random() * 1000) });
  //   });
  // }


  generateRatingChartData() {
    this.timeline.forEach(item => {
      const xAxisDate = new Date(item.year, item.month).toUTCString();
      const trendModel = this.vehicletrenddata.find(rating => rating.year === item.year
        && rating.month - 1 === item.month);
      // console.log('trendModel testing', trendModel);
      if (trendModel) {
        this.tasktrendchartdata.push({ name: xAxisDate, value: trendModel.vehicleserviced});
        // console.log('tasktrendchartdata');
        // console.log(this.tasktrendchartdata);
      } else {
        this.tasktrendchartdata.push({ name: xAxisDate, value: 0 });
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
