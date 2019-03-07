import { Component, OnInit, EventEmitter, ViewChild, Input } from '@angular/core';


@Component({
  selector: 'app-schedule-details',
  templateUrl: './schedule-details.component.html',
  styleUrls: ['./schedule-details.component.css']

})
export class ScheduleDetailsComponent implements OnInit {
  @Input() patientconfid: number;

  constructor() { }
  displayedColumns = ['name', 'startdate', 'enddate', 'view'];

  ngOnInit() {

  }


  // getScheduleDataById() {
  //   const dataListRequest = new ScheduleFilter();
  //   dataListRequest.recid = this.patientconfid
  //   console.log('this.patientconfid', this.patientconfid);
  //   this.patientService.getScheduleDataById(dataListRequest).subscribe((schedulePayloadResponse) => {
  //     if (schedulePayloadResponse.issuccess) {
  //       console.log('schedulePayloadResponse', schedulePayloadResponse);
  //     }
  //   });
  // }


}
