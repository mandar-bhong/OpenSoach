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
  yAxisLabel = ' Count';
  tasktrenddata: TaskTrendModel[] = [];
  tasktrendchartdata = [];
  request = new TaskTrendRequest();
  timeline: TrendChartPerMonthXaxis[] = [];
  hospitalizeLabel = 'Hospitalize';
  dischargedLabel = 'Discharged';
  legendTitle = 'Status';
  customColors = [
    {
      name: this.hospitalizeLabel,
      value: '#28a745'
    },
    {
      name: this.dischargedLabel,
      value: '#52A6CA'
    },
  ];

  constructor(private dashboardService: DashboardService) { }

  ngOnInit() {
    this.getTaskTrendChart();
  }

  getTaskTrendChart() {
    const currentDate = new Date();
    this.request.enddate = new Date(Date.UTC(currentDate.getFullYear(), currentDate.getMonth(), currentDate.getDate(), currentDate.getHours(), currentDate.getMinutes()));
    const ticks = Date.UTC(this.request.enddate.getUTCFullYear(), this.request.enddate.getUTCMonth() - 11, 1);
    this.request.startdate = new Date(ticks);
    this.dashboardService.getTaskTrend(this.request).subscribe(payloadResponse => {
      
      // const dataobject = JSON.parse("{\"issuccess\":true,\"error\":null,\"data\":[\r\n{\r\n\"year\":2018,\r\n\"month\":5,\r\n\"hospitalized\":45,\r\n\"discharged\" :22\r\n},\r\n{\r\n\"year\":2018,\r\n\"month\":6, \r\n\"hospitalized\":81,\r\n\"discharged\" :62\r\n},\r\n{\r\n\"year\":2018,\r\n\"month\":7, \r\n\"hospitalized\":20,\r\n\"discharged\" :66\r\n},\r\n{\r\n\"year\":2018,\r\n\"month\":8, \r\n\"hospitalized\":45,\r\n\"discharged\" :45\r\n},\r\n{\r\n\"year\":2018,\r\n\"month\":9, \r\n\"hospitalized\":10,\r\n\"discharged\" :9\r\n},\r\n{\r\n\"year\":2018,\r\n\"month\":10, \r\n\"hospitalized\":10,\r\n\"discharged\" :9\r\n},\r\n{\r\n\"year\":2018,\r\n\"month\":11, \r\n\"hospitalized\":31,\r\n\"discharged\" :45\r\n},\r\n{\r\n\"year\":2018,\r\n\"month\":12, \r\n\"hospitalized\":50,\r\n\"discharged\" :52\r\n},\r\n{\r\n\"year\":2019,\r\n\"month\":1, \r\n\"hospitalized\":63,\r\n\"discharged\" :88\r\n},\r\n{\r\n\"year\":2019,\r\n\"month\":2, \r\n\"hospitalized\":0,\r\n\"discharged\" :0\r\n},\r\n{\r\n\"year\":2019,\r\n\"month\":3, \r\n\"hospitalized\":210,\r\n\"discharged\" :188\r\n},\r\n{\r\n\"year\":2019,\r\n\"month\":4, \r\n\"hospitalized\":10,\r\n\"discharged\" :9\r\n}\r\n]}");
      // payloadResponse = dataobject;

      if (payloadResponse && payloadResponse.issuccess) {
        console.log("payloadResponse", payloadResponse.data);
        payloadResponse.data.forEach(item => {
          const trendModel = new TaskTrendModel();
          trendModel.copyFrom(item);
          this.tasktrenddata.push(trendModel);
        });
        console.log('data received');
        this.generateSeriesTimeline();
        this.generateRatingChartData();
      }
    });
  }



  generateRatingChartData() {

    const hospitalizedData = { name: this.hospitalizeLabel, series: [] };
    this.tasktrendchartdata.push(hospitalizedData);
    const dischargedData = { name: this.dischargedLabel, series: [] };
    this.tasktrendchartdata.push(dischargedData);

    this.timeline.forEach(item => {
      const xAxisDate = new Date(item.year, item.month).toUTCString();
      const trendModel = this.tasktrenddata.find(rating => rating.year === item.year
        && rating.month - 1 === item.month);
      if (trendModel) {
        hospitalizedData.series.push({ name: xAxisDate, value: trendModel.hospitalized });
        dischargedData.series.push({ name: xAxisDate, value: trendModel.discharged });
      } else {
        hospitalizedData.series.push({ name: xAxisDate, value: 0 });
        dischargedData.series.push({ name: xAxisDate, value: 0 });
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
