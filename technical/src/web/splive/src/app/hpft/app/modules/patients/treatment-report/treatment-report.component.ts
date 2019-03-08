import { Component, OnInit, ViewChild, EventEmitter } from '@angular/core';
import { trigger, state, style, transition, animate } from '@angular/animations';
import { MatPaginator, MatSort, MatTableDataSource } from '@angular/material';
import { TreatmentFilterRequest, TreatmentResponse, ActionTreatmentDataValue, JSONBaseDataModel, DocumentTblInfoModel } from 'app/models/api/patient-models';
import { Subscription, merge, Observable } from 'rxjs';
import { PatientService } from 'app/services/patient.service';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';
import { TranslatePipe } from '../../../../../shared/pipes/translate/translate.pipe';
import { startWith, switchMap, map } from 'rxjs/operators';
import { PayloadResponse } from '../../../../../shared/models/api/payload-models';
import { DataListResponse, DataListRequest } from '../../../../../shared/models/api/data-list-models';
import { TreatmentModel } from 'app/models/ui/patient-models';

@Component({
  selector: 'app-treatment-report',
  templateUrl: './treatment-report.component.html',
  styleUrls: ['./treatment-report.component.css'],
  animations: [
    trigger('detailExpand', [
      state('collapsed', style({ height: '0px', minHeight: '0', display: 'none' })),
      state('expanded', style({ height: '*' })),
      transition('expanded <=> collapsed', animate('225ms cubic-bezier(0.4, 0.0, 0.2, 1)')),
    ]),
  ],

})
export class TreatmentReportComponent implements OnInit {

  displayedColumns = ['treatmentdone', 'txndate', 'view'];
  sortByColumns = [
    { text: 'Treatment Done', value: 'treatmentdone' },
    { text: 'Performed On', value: 'txndate' }
  ];

  @ViewChild(MatPaginator)
  paginator: MatPaginator;
  @ViewChild(MatSort)
  sort: MatSort;
  refreshTable: EventEmitter<null> = new EventEmitter();
  dataSource;
  filteredrecords = 0;
  isLoadingResults: boolean;
  isViewSchedule = false;
  treatmentFilterRequest: TreatmentFilterRequest;
  dataListFilterChangedSubscription: Subscription;
  dataModel = new TreatmentModel();
  admissionid: number;
  treatmentid:number;
  treatResponse: TreatmentResponse<ActionTreatmentDataValue>[] = []

  constructor(public patientService: PatientService,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe) { }

  ngOnInit() {
    this.paginator.pageSize = 10;
    this.sort.direction = 'asc';
    this.paginator.pageIndex = 1;
    this.sort.active = 'admissionid';
    // this.sort.active = this.admissionid;
    this.sort.direction = 'asc';
    this.setDataListing();

    this.dataModel.documentData = new JSONBaseDataModel<DocumentTblInfoModel>();
    this.dataModel.documentData.data = new DocumentTblInfoModel();

  }
  setDataListing(): void {
    this.sort.sortChange.subscribe(() => this.paginator.pageIndex = 0);
    this.refreshTable.subscribe(() => this.paginator.pageIndex = 0);
    merge(this.sort.sortChange, this.paginator.page, this.refreshTable)
      .pipe(
        startWith({}),
        switchMap(() => {
          this.isLoadingResults = true;
          return this.getDataList();
        }),
        map(data => {
          this.isLoadingResults = false;
          return data;
        }),
      ).subscribe(
        payloadResponse => {
          if (payloadResponse && payloadResponse.issuccess) {
            this.filteredrecords = payloadResponse.data.filteredrecords;
            payloadResponse.data.records.forEach((item: any) => {
              const treatmentResponse = new TreatmentResponse<ActionTreatmentDataValue>();
              Object.assign(treatmentResponse, item);
              this.treatResponse.push(treatmentResponse);
            });

            this.dataSource = new MatTableDataSource<TreatmentResponse<ActionTreatmentDataValue>>(this.treatResponse);

            if (this.filteredrecords === 0) {
              this.appNotificationService.info(this.translatePipe.transform('INFO_NO_RECORDS_FOUND'));
            }
          } else {
            this.dataSource = [];
          }
        }
      );
  }

  getDataList(): Observable<PayloadResponse<DataListResponse<TreatmentResponse<string>[]>>> {
    const dataListRequest = new DataListRequest<TreatmentFilterRequest>();
    dataListRequest.page = this.paginator.pageIndex;
    dataListRequest.limit = this.paginator.pageSize + 1;
    dataListRequest.orderby = this.sort.active;
    dataListRequest.filter = new TreatmentFilterRequest();
    dataListRequest.filter.admissionid = this.patientService.admissionid;
    this.dataModel.treatmentid = this.patientService.treatmentid;
    dataListRequest.orderdirection = this.sort.direction;
    return this.patientService.getTreatmentList(dataListRequest);
  }

  // code block cehcking obejct is emppty.
  checkEmptyObjects(object): boolean {
    if (Object.keys(object).length > 0) {
      return true;
    } else {
      return false;
    }
  }// end of fucntion 

  // code bloxk for view schedule detsils  of particular action 
  viewSchedule(element) {
    console.log('view schedule clickd');
    this.isViewSchedule = true;
  }

  setOpenCloseSchedule() {
    console.log('view setOpenCloseSchedule clickd');
    this.isViewSchedule = false;
  }
  sortByChanged() {
    this.sort.sortChange.next(this.sort);
  }

  sortDirectionAsc() {
    this.sort.direction = 'asc';
    this.sort.sortChange.next(this.sort);
  }

  sortDirectionDesc() {
    this.sort.direction = 'desc';
    this.sort.sortChange.next(this.sort);
  }

}
