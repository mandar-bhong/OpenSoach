import { Component, OnInit, Input } from '@angular/core';
import { ScheduleFilter } from 'app/models/api/schedule-request';
import { PatientService } from 'app/services/patient.service';
import { ScheduleDataResponse, SchedularConfigData } from 'app/models/api/schedule-response';
import { ConfigCodeType, FREQUENCY_ZERO, FREQUENCY_ONE } from 'app/app-constants';
import { ScheduleService } from 'app/services/patient-detail-sevices/schedule.service';

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
  constructor(private patientService: PatientService, private scheduleService: ScheduleService) { }
  @Input() patientconfid: number;
  scheduleDataResponse = new ScheduleDataResponse<SchedularConfigData>();
  ngOnInit() {
    console.log('this.patientconfid', this.patientconfid);
    this.getScheduleDataById();
  }
  getScheduleDataById() {
    const dataListRequest = new ScheduleFilter();
    dataListRequest.recid = this.patientconfid
    console.log('this.patientconfid', this.patientconfid);
    this.scheduleService.getScheduleDataById(dataListRequest).subscribe((schedulePayloadResponse) => {
      if (schedulePayloadResponse.issuccess) {
        this.isDataReveived = true;
        console.log('schedulePayloadResponse', schedulePayloadResponse);
        Object.assign(this.scheduleDataResponse, schedulePayloadResponse.data);
        this.scheduleDataResponse.conf = JSON.parse(schedulePayloadResponse.data.conf)
      }
    });
  }
  // code block for check status
  checkStatus(status: number, enddate: string) {
    if (status == 0) {
      const enddt = new Date(enddate);
      const currentdt = new Date();
      if (enddt.getTime() > currentdt.getTime()) {
        return 'Active';
      } else {
        return 'Completed';
      }
    } else {
      return 'Cancelled';
    }
  }// end of code block
}
