import { Component, OnInit, EventEmitter, ViewChild, Input, OnDestroy } from '@angular/core';
import { MatSort, MatPaginator, MatTableDataSource } from '@angular/material';
import { merge, Observable, Subscription } from 'rxjs';
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
import { ScheduleDataResponse } from 'app/models/api/schedule-response';
import { ScheduleFilter } from 'app/models/api/schedule-request';
import { ComplaintsModule } from 'app/modules/complaints/complaints.module';
import { ScheduleService } from 'app/services/patient-detail-sevices/schedule.service';


@Component({
  selector: 'app-view-monitor-schedule',
  templateUrl: './view-monitor-schedule.component.html',
  styleUrls: ['./view-monitor-schedule.component.css', '../../transaction-details/transaction-details.component.css'],
  animations: [
    trigger('detailExpand', [
      state('collapsed', style({ height: '0px', minHeight: '0', display: 'none' })),
      state('expanded', style({ height: '*' })),
      transition('expanded <=> collapsed', animate('225ms cubic-bezier(0.4, 0.0, 0.2, 1)')),
    ]),
  ],
})
export class ViewMonitorScheduleComponent implements OnInit, OnDestroy {
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
  scheduleResponse: ScheduleDataResponse<any>[] = [];
  dataListRequest: DataListRequest<TransactionDetailsFilter>;
  isViewSchedule = false;
  dataListFilterChangedSubscription: Subscription;
  scheduleFilter: ScheduleFilter;
  constructor(
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe,
    private patientService: PatientService,
    private scheduleService: ScheduleService) { }
  displayedColumns = ['name', 'startdate', 'enddate', 'view'];
  sortByColumns = [
    { text: 'Name', value: 'name' },
    { text: 'Start', value: 'startDate' },
    { text: 'End', value: 'enddate' }
  ];
  // columnsToDisplay = ['fname', 'date'];
  ngOnInit() {
    console.log('getDataListing executed');
    this.paginator.pageSize = 10;
    this.sort.direction = 'asc';
    this.sort.active = 'enddate';
    this.sort.direction = 'asc';
    this.getDataListing();
    this.dataListFilterChangedSubscription = this.scheduleService.dataListSubject.subscribe(value => {
      this.scheduleFilter = value;
      // this.scheduleFilter.conftypecode = value.conftypecode;
      this.refreshTable.emit();
    });

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
            this.scheduleResponse = [];
            payloadResponse.data.records.forEach((item: any) => {
              const ActionTransactionData = new ScheduleDataResponse<any>();
              Object.assign(ActionTransactionData, item);
              console.log('item', item.txndata);
              const confData = JSON.parse(item.conf);
              ActionTransactionData.conf = confData;
              this.scheduleResponse.push(ActionTransactionData);
            });
            console.log(' this.scheduleResponse', this.scheduleResponse);
            this.dataSource = new MatTableDataSource<ScheduleDataResponse<any>>(this.scheduleResponse);
            if (this.filteredrecords === 0) {
              //   this.appNotificationService.info(this.translatePipe.transform('INFO_NO_RECORDS_FOUND'));
            }
          } else {
            this.dataSource = [];
          }
        }
      );
  }


  getDataList(): Observable<PayloadResponse<DataListResponse<ScheduleDataResponse<string>[]>>> {
    const dataListRequest = new DataListRequest<ScheduleFilter>();
    dataListRequest.orderdirection = this.sort.direction;
    dataListRequest.limit = this.paginator.pageSize;
    dataListRequest.page = this.paginator.pageIndex + 1;
    dataListRequest.orderby = this.sort.active;
    dataListRequest.filter = new ScheduleFilter();
    dataListRequest.filter.admissionid = this.patientService.admissionid;
    dataListRequest.filter.conftypecode = 'Monitor';
    return this.scheduleService.getDataList(dataListRequest);
    // return this.patientService.getScheduleData(dataListRequest);
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
  ngOnDestroy(): void {
    if (this.dataListFilterChangedSubscription) {
      this.dataListFilterChangedSubscription.unsubscribe();
    }
  }
}




