import { Component, OnInit, EventEmitter, ViewChild } from '@angular/core';
import { MatSort, MatPaginator, MatTableDataSource } from '@angular/material';
import { merge, Observable } from 'rxjs';
import { startWith, switchMap, map } from 'rxjs/operators';
import { PayloadResponse } from '../../../../../../shared/models/api/payload-models';
import { DataListResponse, DataListRequest } from '../../../../../../shared/models/api/data-list-models';
import { PatientDetaListResponse, PatientFilterRequest } from 'app/models/api/patient-data-models';
import { AppNotificationService } from '../../../../../../shared/services/notification/app-notification.service';
import { TranslatePipe } from '../../../../../../shared/pipes/translate/translate.pipe';
import { PatientService } from 'app/services/patient.service';
import { TransactionDetailsFilter } from 'app/models/api/transaction-details';
import { ActionTransactionResponse, ActionTransactionDataValue } from 'app/models/api/transaction-details-response';
import { animate, state, style, transition, trigger } from '@angular/animations';
import { ConfigCodeType } from 'app/app-constants';
import { ScheduleDataResponse } from 'app/models/api/schedule-response';

@Component({
  selector: 'app-view-intake-transaction',
  templateUrl: './view-intake-transaction.component.html',
  styleUrls: ['./view-intake-transaction.component.css', '../transaction-details.component.css'],
  animations: [
    trigger('detailExpand', [
      state('collapsed', style({ height: '0px', minHeight: '0', display: 'none' })),
      state('expanded', style({ height: '*' })),
      transition('expanded <=> collapsed', animate('225ms cubic-bezier(0.4, 0.0, 0.2, 1)')),
    ]),
  ],
})
export class ViewIntakeTransactionComponent implements OnInit {

  @ViewChild(MatPaginator)
  paginator: MatPaginator;
  @ViewChild(MatSort)
  sort: MatSort;
  refreshTable: EventEmitter<null> = new EventEmitter();
  dataSource;
  filteredrecords = 0;
  isLoadingResults: boolean;
  patientFilterRequest: PatientFilterRequest;
  expandedElement: ActionTransactionResponse<string> | null;
  transactionResponse: ActionTransactionResponse<ActionTransactionDataValue>[] = [];
  dataListRequest: DataListRequest<TransactionDetailsFilter>;
  isViewSchedule = false;
  ConfigCodeType = ConfigCodeType;
  constructor(
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe,
    private patientService: PatientService) { }
  displayedColumns = ['actionname', 'scheduledtime', 'by', 'view'];
  sortByColumns = [
    { text: 'Name', value: 'actionname' },
    { text: 'Performed On', value: 'scheduledtime' }
  ];
  ngOnInit() {
    this.paginator.pageSize = 10;
    this.sort.direction = 'asc';
    this.sort.active = 'scheduledtime';
    this.getDataListing();
  }
  getDataListing(): void {
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
            this.transactionResponse = [];
            payloadResponse.data.records.forEach((item: any) => {
              const ActionTransactionData = new ActionTransactionResponse<ActionTransactionDataValue>();
              Object.assign(ActionTransactionData, item);
              const txnJsonData = JSON.parse(item.txndata);
              ActionTransactionData.txndata = txnJsonData;
              this.transactionResponse.push(ActionTransactionData);
            });

            this.dataSource = new MatTableDataSource<ActionTransactionResponse<ActionTransactionDataValue>>(this.transactionResponse);

            if (this.filteredrecords === 0) {
              // this.appNotificationService.info(this.translatePipe.transform('INFO_NO_RECORDS_FOUND'));
            }
          } else {
            this.dataSource = [];
          }
        });
  }

  getDataList(): Observable<PayloadResponse<DataListResponse<ActionTransactionResponse<string>[]>>> {
    const dataListRequest = new DataListRequest<TransactionDetailsFilter>();
    dataListRequest.orderdirection = this.sort.direction;
    dataListRequest.limit = this.paginator.pageSize;
    dataListRequest.page = this.paginator.pageIndex + 1;
    dataListRequest.orderby = this.sort.active;
    dataListRequest.filter = new TransactionDetailsFilter();
    dataListRequest.filter.conftypecode = ConfigCodeType.INTAKE;
    dataListRequest.filter.admissionid = this.patientService.admissionid;
    return this.patientService.getActionTransaction(dataListRequest);
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
    this.isViewSchedule = true;
  }

  setOpenCloseSchedule() {
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