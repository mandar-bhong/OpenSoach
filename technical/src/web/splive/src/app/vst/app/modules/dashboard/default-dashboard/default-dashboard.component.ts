import { Component, OnInit } from '@angular/core';

import {
  ComplaintSummaryModel, DeviceSummaryModel, ServicePointSummaryModel,
  SnapshotModel, TimeModel
} from '../../../models/ui/dashboard-models';
import { DashboardService } from '../../../services/dashboard.service';
import { EnumDataSourceItem } from '../../../../../shared/models/ui/enum-datasource-item';
import { SNAPSHOT_STATE } from '../../../app-constants';
import { SnapShotResponse } from '../../../models/api/dashboard-models';

@Component({
  selector: 'app-default-dashboard',
  templateUrl: './default-dashboard.component.html',
  styleUrls: ['./default-dashboard.component.css']
})
export class DefaultDashboardComponent implements OnInit {
  snapshot = new SnapshotModel();
  snapshotStates: EnumDataSourceItem<number>[];
  snapshotdata: SnapshotModel[] = [];
  snapshotdatamap: Map<number, SnapshotModel>;
  tokendata = new SnapshotModel();
  jobcreated = new SnapshotModel();
  jobinprogress = new SnapshotModel();
  jobcompleted = new SnapshotModel();
  vehical = new SnapshotModel();
  xyz;
  timedata = new TimeModel();
  constructor(private dashboardService: DashboardService) { }
  devicesummary = new DeviceSummaryModel();
  spsummary = new ServicePointSummaryModel();
  complaintsummary = new ComplaintSummaryModel();
  ngOnInit() {
    this.getDeviceSummary();
    this.getServicePointSummary();
    this.getComplaintSummary();
    this.getSnapShot();
    this.getTime();
  }

  getDeviceSummary() {
    this.dashboardService.getDeviceSummary().subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.devicesummary.copyFrom(payloadResponse.data);
      }
    });
  }

  getServicePointSummary() {
    this.dashboardService.getServicePointSummary().subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.spsummary.copyFrom(payloadResponse.data);
      }
    });
  }

  getComplaintSummary() {
    this.dashboardService.getComplaintSummary().subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.complaintsummary.copyFrom(payloadResponse.data);
      }
    });
  }
  getSnapShot() {
    this.snapshot.startdate = new Date();
    this.snapshot.enddate = new Date();
    this.snapshot.startdate.setHours(0, 0, 0, 0);
    this.snapshot.enddate.setHours(24, 0, 0, 0);
    this.dashboardService.getSnapShot({
      startdate: this.snapshot.startdate,
      enddate: this.snapshot.enddate
    }).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.snapshotdatamap = new Map(payloadResponse.data.map(ab => [ab.status, ab] as [number, SnapshotModel]));
        console.log('this.snapshotdatamap');
        console.log(this.snapshotdatamap);

        if (this.snapshotdatamap.get(SNAPSHOT_STATE.TOKEN_GENERATED) != null) {
          this.tokendata = this.snapshotdatamap.get(SNAPSHOT_STATE.TOKEN_GENERATED);
          console.log('tokendata');
          console.log(this.tokendata);
        }

        if (this.snapshotdatamap.get(SNAPSHOT_STATE.JOB_CREATED) != null) {
          this.jobcreated = this.snapshotdatamap.get(SNAPSHOT_STATE.JOB_CREATED);
          console.log('jobcreated');
          console.log(this.jobcreated);
        }

        if (this.snapshotdatamap.get(SNAPSHOT_STATE.JOB_INPROGRESS) != null) {
          this.jobinprogress = this.snapshotdatamap.get(SNAPSHOT_STATE.JOB_INPROGRESS);
          console.log('jobinprogress');
          console.log(this.jobinprogress);
        }

        if (this.snapshotdatamap.get(SNAPSHOT_STATE.JOB_COMPLETED) != null) {
          this.jobcompleted = this.snapshotdatamap.get(SNAPSHOT_STATE.JOB_COMPLETED);
          console.log('jobcompleted');
          console.log(this.jobcompleted);
        }

        if (this.snapshotdatamap.get(SNAPSHOT_STATE.VEHICAL_DELIVERY) != null) {
          this.vehical = this.snapshotdatamap.get(SNAPSHOT_STATE.VEHICAL_DELIVERY);
          console.log('vehical');
          console.log(this.vehical);
        }
      }
    });
  }
  getTime() {
    this.timedata.startdate = new Date();
    this.timedata.enddate = new Date();
    this.timedata.startdate.setHours(0, 0, 0, 0);
    this.timedata.enddate.setHours(24, 0, 0, 0);
    this.dashboardService.getTime({
      startdate: this.timedata.startdate,
      enddate: this.timedata.enddate
    }).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        // this.timedata.copyFrom(payloadResponse.data);
        const wtime = payloadResponse.data.waittime / 60; // sec to min
        this.timedata.waittime = wtime;
        console.log('wtime');
        console.log(wtime);
        const jetime = payloadResponse.data.jobexetime / 3600; // sec to hrs
        this.timedata.jobexetime = jetime;
        console.log('jetime');
        console.log(jetime);
        const jctime = payloadResponse.data.jobcreationtime / 60; // sec to min
        this.timedata.jobcreationtime = jctime;
        console.log('jctime');
        console.log(jctime);
        const dtime = payloadResponse.data.deliverytime / 60; // sec to min
        this.timedata.deliverytime = dtime;
        console.log('dtime');
        console.log(dtime);
      }
    });
  }
  windowfull() {

  }
}
