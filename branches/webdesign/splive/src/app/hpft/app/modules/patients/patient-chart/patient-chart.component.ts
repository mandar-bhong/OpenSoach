import { Component, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs';

import { ServicepointListResponse } from '../../../../../prod-shared/models/api/servicepoint-models';
import { SpServiceTxnService } from '../../../../../prod-shared/services/spservice/sp-service-txn.service';
import { EnumDataSourceItem } from '../../../../../shared/models/ui/enum-datasource-item';
import { TranslatePipe } from '../../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';
import { EditRecordBase } from '../../../../../shared/views/edit-record-base';
import { PatientDetailAddRequest } from '../../../models/api/patient-models';
import { ChartDataViewModel, PatientDayWiseTxn } from '../../../models/ui/chart-conf-models';
import { PatientDataModel } from '../../../models/ui/patient-models';
import { PatientService } from '../../../services/patient.service';

// import { ChartDataViewModel } from '../../../../../hkt/app/models/ui/chart-conf-models';
// import { reverse } from 'dns';

@Component({
  selector: 'app-patient-chart',
  templateUrl: './patient-chart.component.html',
  styleUrls: ['./patient-chart.component.css']
})
export class PatientChartComponent extends EditRecordBase implements OnInit, OnDestroy {
  dataModel = new PatientDataModel();
  routeSubscription: Subscription;
  patientStates: EnumDataSourceItem<number>[];
  // selecteddateoption = 0;
  selectedDate = new Date();
  splist: ServicepointListResponse[] = [];
  testdate;
  servinid;
  patient = [];
  count = [];
  datachartModel = new ChartDataViewModel();
  patientDayWiseTransactions: PatientDayWiseTxn[] = [];
  constructor(private patientService: PatientService,
    private route: ActivatedRoute,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe,
    private spServiceTxnService: SpServiceTxnService) {
    super();
    this.iconCss = 'fa fa-user';
    this.pageTitle = 'Patient File';
  }

  ngOnInit() {
    this.dataModel.patientdetails = new PatientDetailAddRequest();
    this.patientStates = this.patientService.getPatientStates();
    this.routeSubscription = this.route.queryParams.subscribe(params => {
      if (params['id']) {
        this.dataModel.patientid = Number(params['id']);
        this.getPatientDetails();
      }
      this.callbackUrl = params['callbackurl'];
    });
  }
  getPatientDetails() {
    this.patientService.getPatientDetails({ recid: this.dataModel.patientid }).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        if (payloadResponse.data) {
          this.dataModel.copyFrom(payloadResponse.data);
          this.testdate = payloadResponse.data;
          this.subTitle = this.dataModel.patientdetails.patientname;
          this.isEditable = false;
          this.getPatientTransactions();
          this.getServicepointList();
        }
      }
    });
  }
  getPatientTransactions() {
    this.spServiceTxnService.getPatientTransactions({
      spid: this.testdate.spid,
      servinid: this.testdate.servinid,
    }).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.datachartModel.copyFromTransactions(payloadResponse.data);
        console.log('this.datachartModel', this.datachartModel);
        this.datachartModel.txns.forEach(item => {
          const existingDayWiseRecord = this.patientDayWiseTransactions.find(a => a.day.toDateString() === item.txndate.toDateString());
          if (existingDayWiseRecord) {
            existingDayWiseRecord.txn.push(item);
            console.log('existingDayWiseRecord', existingDayWiseRecord);
          } else {
            const newDayWiseRecord = new PatientDayWiseTxn();
            newDayWiseRecord.day = item.txndate;
            newDayWiseRecord.txn = [];
            newDayWiseRecord.txn.push(item);
            this.patientDayWiseTransactions.push(newDayWiseRecord);
          }
        });

        console.log('sorting');
        this.patientDayWiseTransactions.sort(this.sortTxnDatesDesc);
      }
    });
  }

  private sortTxnDatesDesc(a:PatientDayWiseTxn, b:PatientDayWiseTxn) {
    if (a.day < b.day) {
     return 1;
    }
    if (a.day > b.day) {
     return -1;
    }
    return 0;
   }

  getServicepointList() {
    this.patientService.getServicepointList().subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.splist = payloadResponse.data;
      }
    });
  }
  getpatientstate(value: number) {
    if (this.patientStates && value) {
      return this.patientStates.find(a => a.value === value).text;
    }
  }
  getwardname(value: number) {
    if (this.splist && value) {
      const item = this.splist.find(a => a.spid === value);
      if (item) {
        return item.spname;
      }
    }
  }
  closeForm() {
    this.router.navigate([this.callbackUrl], { skipLocationChange: true });
  }
  ngOnDestroy() {
    if (this.routeSubscription) {
      this.routeSubscription.unsubscribe();
    }
  }
}
