import { Component, OnInit } from '@angular/core';
import { curveLinear } from 'd3-shape';

import { TaskTrendRequest } from '../../../models/api/dashboard-models';
import { TaskTrendModel, TrendChartPerMonthXaxis } from '../../../models/ui/dashboard-models';
import { DashboardService } from '../../../services/dashboard.service';

@Component({
  selector: 'app-service-time-weekly',
  templateUrl: './service-time-weekly.component.html',
  styleUrls: ['./service-time-weekly.component.css']
})
export class ServiceTimeWeeklyComponent implements OnInit {

  curve = curveLinear;
  xAxisLabel = 'Time in hours';
  yAxisLabel = 'Days';
  tasktrenddata: TaskTrendModel[] = [];
  tasktrendchartdata = [];
  request = new TaskTrendRequest();
  timeline: TrendChartPerMonthXaxis[] = [];
  ontimeLabel = 'Vehicles';
  formatXAxis;
  selecteddateoption = '0';
  // delayedLabel = 'Delayed';
  // missedLabel = 'Missed';
  // legendTitle = 'Task Status';
  time = 1;
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
    // {
    //   name: this.missedLabel,
    //   value: '#ff5252'
    // },
  ];
  data: any;
  constructor(private dashboardService: DashboardService) { }

  ngOnInit() {
    this.jobexcu();
  }
  // optionChange() {
  //   switch (this.selecteddateoption) {
  //     case '0':
  //       this.jobexcu();
  //       console.log('test');
  //       console.log('this.jobexcu', this.jobexcu() );
  //       break;
  //     case '1':
  //       this.waittime();
  //       console.log('test');
  //       console.log('this.waittime', this.waittime() );
  //       break;
  //   }
  // }
  jobexcu() {
    this.data = [
      {
        'name': 'MON',
        'value': 33
      },
      {
        'name': 'TUE',
        'value': 45
      },
      {
        'name': 'WED',
        'value': 60
      },
      {
        'name': 'THU',
        'value': 70
      },
      {
        'name': 'FRI',
        'value': 80
      },
      {
        'name': 'SAT',
        'value': 70
      },
      {
        'name': 'SUN',
        'value': 100
      }
    ];
  }
  waittime() {
    this.data = [
      {
        'name': 'MON',
        'value': 33
      },
      {
        'name': 'TUE',
        'value': 45
      },
      {
        'name': 'WED',
        'value': 60
      },
      {
        'name': 'THU',
        'value': 70
      },
      {
        'name': 'FRI',
        'value': 80
      },
      {
        'name': 'SAT',
        'value': 11
      },
      {
        'name': 'SUN',
        'value': 33
      }
    ];
  }
  creationtime() {
    this.data = [
      {
        'name': 'MON',
        'value': 70
      },
      {
        'name': 'TUE',
        'value': 55
      },
      {
        'name': 'WED',
        'value': 45
      },
      {
        'name': 'THU',
        'value': 1
      },
      {
        'name': 'FRI',
        'value': 20
      },
      {
        'name': 'SAT',
        'value': 60
      },
      {
        'name': 'SUN',
        'value': 100
      }
    ];
  }
  deliverytime() {
    this.data = [
      {
        'name': 'MON',
        'value': 43
      },
      {
        'name': 'TUE',
        'value': 55
      },
      {
        'name': 'WED',
        'value': 70
      },
      {
        'name': 'THU',
        'value': 30
      },
      {
        'name': 'FRI',
        'value': 61
      },
      {
        'name': 'SAT',
        'value': 35
      },
      {
        'name': 'SUN',
        'value': 100
      }
    ];
  }
}
