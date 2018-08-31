import { Component, OnInit } from '@angular/core';
import { curveLinear } from 'd3-shape';

import { TaskTrendRequest } from '../../../models/api/dashboard-models';
import { TaskTrendModel, TrendChartPerMonthXaxis } from '../../../models/ui/dashboard-models';
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
  tasktrenddata: TaskTrendModel[] = [];
  tasktrendchartdata = [];
  request = new TaskTrendRequest();
  timeline: TrendChartPerMonthXaxis[] = [];
  ontimeLabel = 'Vehicles';
  // delayedLabel = 'Delayed';
  // missedLabel = 'Missed';
  // legendTitle = 'Task Status';
  customColors = [
    {
      name: this.ontimeLabel,
      value: '#28a745'
    },
    // {
    //   name: this.delayedLabel,
    //   value: '#ffc107'
    // },
    // {
    //   name: this.missedLabel,
    //   value: '#ff5252'
    // },
  ];
data:any;
  constructor(private dashboardService: DashboardService) { }

  ngOnInit() {
   this.data= [
      {
        "name": "MON",
        "value": 82
      },
      {
        "name": "TUE",
        "value": 85
      },
      {
        "name": "WED",
        "value": 60
      },
      {
        "name": "THU",
        "value": 70
      },
      {
        "name": "FRI",
        "value": 80
      },
      {
        "name": "SAT",
        "value": 112
      },
      {
        "name": "SUN",
        "value": 100
      }
    ];
   // this.getTaskTrend();
  }

//   getTaskTrend() {
//     const currentDate = new Date();
//     this.request.enddate = new Date(Date.UTC(
//       currentDate.getFullYear(), currentDate.getMonth(), currentDate.getDate(), currentDate.getHours(), currentDate.getMinutes()));

//     const ticks = Date.UTC(this.request.enddate.getUTCFullYear(), this.request.enddate.getUTCMonth() - 11, 1);
//     this.request.startdate = new Date(ticks);

// //     this.dashboardService.getTaskTrend(this.request).subscribe(payloadResponse => {
// //       if (payloadResponse && payloadResponse.issuccess) {
// //         payloadResponse.data.forEach(item => {
// //           const trendModel = new TaskTrendModel();
// //           trendModel.copyFrom(item);
// //           this.tasktrenddata.push(trendModel);
// //         });
// // console.log('data received');
// //         this.generateSeriesTimeline();
// //         this.generateRatingChartData();
// //       }
// //     });

//     this.generateSeriesTimeline();
//         this.generateRatingChartData();
//   }



//   generateRatingChartData() {
//     const ontimeData = { name: this.ontimeLabel, series: [] };
//     this.tasktrendchartdata.push(ontimeData);
//     const delayedData = { name: this.delayedLabel, series: [] };
//     this.tasktrendchartdata.push(delayedData);
//     const missedData = { name: this.missedLabel, series: [] };
//     this.tasktrendchartdata.push(missedData);

//     let dummyOnTime =1000;
//     let dummyDelayed =205;
//     let dummyMissed =50;
//     this.timeline.forEach(item => {
//       const xAxisDate = new Date(item.year, item.month).toUTCString();
//       // const trendModel = this.tasktrenddata.find(rating => rating.year === item.year
//       //   && rating.month-1 === item.month);

//       // if (trendModel) {
//       //   ontimeData.series.push({ name: xAxisDate, value: trendModel.ontime });
//       //   delayedData.series.push({ name: xAxisDate, value: trendModel.delayed });
//       //   missedData.series.push({ name: xAxisDate, value: trendModel.missed });
//       // } else {
//         ontimeData.series.push({ name: xAxisDate, value: dummyOnTime });
//         delayedData.series.push({ name: xAxisDate, value: dummyDelayed });
//         missedData.series.push({ name: xAxisDate, value: dummyMissed });

//         dummyOnTime=dummyOnTime+12;
//         dummyDelayed=dummyDelayed-13;
//         dummyMissed=dummyMissed-3;
//       //}
//     });
//   }

//   generateSeriesTimeline() {
//     let month = this.request.startdate.getMonth();
//     let year = this.request.startdate.getFullYear();
//     for (let i = 0; i < 12; i++) {
//       this.timeline.push({ year: year, month: month });
//       month = month + 1;
//       if (month > 11) {
//         month = 0;
//         year = year + 1;
//       }
//     }
//   }

//   formatXAxis(value: string) {
//     const date = new Date(value);
//     return date.toLocaleDateString('en-US', { month: 'short' });
//   }
}
