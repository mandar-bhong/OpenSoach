import { Component, Input, OnInit } from '@angular/core';
import { ScheduleDataResponse, SchedularConfigData } from 'app/models/api/schedule-response';
import { ConfigCodeType, FREQUENCY_ZERO, FREQUENCY_ONE, PATIENT_CHECK_STATE } from 'app/app-constants'
@Component({
  selector: 'app-schedule-detail-expand-view',
  templateUrl: './schedule-detail-expand-view.component.html',
  styleUrls: ['./schedule-detail-expand-view.component.css']
})
export class ScheduleDetailExpandViewComponent implements OnInit {
  @Input() schedule: ScheduleDataResponse<string>;
  scheduleDataResponse: ScheduleDataResponse<SchedularConfigData>;
  configCodeType = ConfigCodeType;
  freuencyZero = FREQUENCY_ZERO;
  freuencyOne = FREQUENCY_ONE
  PATIENT_CHECK_STATE:PATIENT_CHECK_STATE;
  constructor() {
  }

  ngOnInit() {
    // SchedularConfigData
    this.scheduleDataResponse = new ScheduleDataResponse<SchedularConfigData>();
    Object.assign(this.scheduleDataResponse, this.schedule);
    console.log('expand schedule executed', this.schedule);
    console.log('processed conf data', this.scheduleDataResponse);
  }
  timeConvert(n) {
    var num = n;
    var hours = (num / 60);
    var rhours = Math.floor(hours);
    var minutes = (hours - rhours) * 60;
    var rminutes = Math.round(minutes);
    if (rhours > 0) {
      if (rminutes > 0) {
        return rhours + " hour & " + rminutes + " minute";
      } else {
        return rhours + " hour";
      }
    } else {
      return rminutes + " minute";
    }
  }
  // code block for check status
  checkStatus(status: number, enddate: string) {
    if (status == 0) {
      const enddt = new Date(enddate);
      const currentdt = new Date();
      if (enddt.getTime() > currentdt.getTime()) {
        return PATIENT_CHECK_STATE.ACTIVE;
      } else {
        return PATIENT_CHECK_STATE.COMPLETED;
      }
    } else {
      return PATIENT_CHECK_STATE.CANCELLED;
    }
  }// end of code block
}
