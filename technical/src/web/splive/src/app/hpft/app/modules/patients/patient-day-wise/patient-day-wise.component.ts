import { ChangeDetectorRef, Component, Input, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';

import { ChartDataViewModel, ChartTransactionModel } from '../../../../../hkt/app/models/ui/chart-conf-models';
import { SpServiceTxnService } from '../../../../../prod-shared/services/spservice/sp-service-txn.service';
import { TranslatePipe } from '../../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';
import { PatientDataModel } from '../../../models/ui/patient-models';
import { PatientService } from '../../../services/patient.service';

@Component({
  selector: 'app-patient-day-wise',
  templateUrl: './patient-day-wise.component.html',
  styleUrls: ['./patient-day-wise.component.css']
})
export class PatientDayWiseComponent implements OnInit {
  dataModel = new ChartDataViewModel();
  datachartModel = new PatientDataModel();
  txns: ChartTransactionModel[];
  displayedColumns = ['time', 'taskname', 'fopname', 'value', 'comment'];
  dataSource = [];
  @Input()
  patientDayWiseTxn: ChartTransactionModel[] = [];
  constructor(private patientService: PatientService,
    private route: ActivatedRoute,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private abc: ChangeDetectorRef,
    private translatePipe: TranslatePipe,
    private spServiceTxnService: SpServiceTxnService) { }

  ngOnInit() {
    console.log('recved data');
    console.log(this.patientDayWiseTxn);
    this.dataSource = this.patientDayWiseTxn;
  }


}

