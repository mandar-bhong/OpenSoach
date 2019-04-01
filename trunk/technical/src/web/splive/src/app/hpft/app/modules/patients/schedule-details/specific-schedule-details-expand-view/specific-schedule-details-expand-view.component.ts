import { Component, OnInit, Input } from '@angular/core';
import { ScheduleFilter } from 'app/models/api/schedule-request';
import { PatientService } from 'app/services/patient.service';
import { ScheduleDataResponse, SchedularConfigData } from 'app/models/api/schedule-response';
import { ConfigCodeType, FREQUENCY_ZERO, FREQUENCY_ONE, PATIENT_CHECK_STATE } from 'app/app-constants';
import { ScheduleService } from 'app/services/patient-detail-sevices/schedule.service';
import { Timeconversion } from 'app/helper';

@Component({
  selector: 'app-specific-schedule-details-expand-view',
  templateUrl: './specific-schedule-details-expand-view.component.html',
  styleUrls: ['./specific-schedule-details-expand-view.component.css']
})
export class SpecificScheduleDetailsExpandViewComponent implements OnInit {
  configCodeType = ConfigCodeType;
  freuencyZero = FREQUENCY_ZERO;
  freuencyOne = FREQUENCY_ONE
  isDataReveived = false;
  PATIENT_CHECK_STATE:PATIENT_CHECK_STATE;
  timeconversion :Timeconversion;
  interval: any;
  constructor(private patientService: PatientService) { }
  @Input() patientconfid: number;
  scheduleDataResponse = new ScheduleDataResponse<SchedularConfigData>();
  ngOnInit() {
    this.getScheduleDataById();
  }
  getScheduleDataById() {
    const dataListRequest = new ScheduleFilter();
    dataListRequest.recid = this.patientconfid
    this.patientService.getScheduleDataById(dataListRequest).subscribe((schedulePayloadResponse) => {
      if (schedulePayloadResponse.issuccess) {
        this.isDataReveived = true;
        Object.assign(this.scheduleDataResponse, schedulePayloadResponse.data);
        this.scheduleDataResponse.conf = JSON.parse(schedulePayloadResponse.data.conf);

        // this.interval  = Timeconversion.timeConvert(this.scheduleDataResponse.conf.interval);
      }
    });
  }

  // import fromstatic class .
  timeConvert(time:number) {
    var num = time;
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
