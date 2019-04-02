import { Component, Input, OnInit } from '@angular/core';
import { ScheduleDataResponse, SchedularConfigData } from 'app/models/api/schedule-response';
import { ConfigCodeType, FREQUENCY_ZERO, FREQUENCY_ONE, PATIENT_CHECK_STATE } from 'app/app-constants';
import { TimeConversionHelper } from 'app/helpers/time-conversion-helper';
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
  }

  // import fromstatic class .
  timeConvert(time: number) {
    return TimeConversionHelper.timeConvert(time);
  }
  // code block for check status
  checkStatus(status: number, enddate: string) {
    // import status value from constants.
    if (status == 0) {
      const enddt = new Date(enddate);
      const currentdt = new Date();
      if (enddt.getTime() > currentdt.getTime()) {
        return PATIENT_CHECK_STATE.ACTIVE;
      } else {
        return PATIENT_CHECK_STATE.COMPLETED;
      }
    } else {
      return PATIENT_CHECK_STATE.STOPPED;
    }
  }// end of code block
  
  convertToDate(minutes: number) {
    let date = new Date();
    date.setHours(0, 0, 0, 0);
    date.setMinutes(minutes);
    return date;
    }

}
