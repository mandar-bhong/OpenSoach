import { Component, EventEmitter, OnDestroy, OnInit, ViewChild } from '@angular/core';
import { MatPaginator, MatSort } from '@angular/material';
import { Router } from '@angular/router';
import { merge, Observable, Subscription } from 'rxjs';
import { map, startWith, switchMap } from 'rxjs/operators';
import { FormControl, FormGroup, Validators, FormBuilder } from '@angular/forms';

import { JobFiltrRequest, JobDataListResponse, StatusChangeRequest } from '../../../../models/api/job-models';
import { AppNotificationService } from '../../../../../../shared/services/notification/app-notification.service';
import { TranslatePipe } from '../../../../../../shared/pipes/translate/translate.pipe';
import { JobService } from '../../../../services/job.service';
import { PayloadResponse } from '../../../../../../shared/models/api/payload-models';
import { DataListResponse, DataListRequest } from '../../../../../../shared/models/api/data-list-models';

@Component({
  selector: 'app-list-view',
  templateUrl: './list-view.component.html',
  styleUrls: ['./list-view.component.css']
})
export class ListViewComponent implements OnInit, OnDestroy {
  displayedColumns = ['vehicleno', 'generatedon', 'intime', 'outtime', 'token', 'state', 'action'];
  sortByColumns = [{ text: 'Vehicle Number', value: 'vehicleno' },
  { text: 'Service Date', value: 'generated_on' },
  { text: 'Intime', value: 'intime' },
  { text: 'Outtime', value: 'outtime' },
  { text: 'Token', value: 'token' },
  { text: 'Status', value: 'state' }
  ];
  @ViewChild(MatPaginator)
  paginator: MatPaginator;
  @ViewChild(MatSort)
  sort: MatSort;
  refreshTable: EventEmitter<null> = new EventEmitter();
  dataSource;
  filteredrecords = 0;
  isLoadingResults = true;
  userFilterRequest: JobFiltrRequest;
  dataListFilterChangedSubscription: Subscription;
  demodata: JobDataListResponse[];
  selected: JobDataListResponse;
  amountform: FormGroup;
  amount: number;
  constructor(private jobService: JobService,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe) { }

  ngOnInit() {
    this.createControls();
    this.paginator.pageSize = 10;
    this.sort.active = 'vehicleno';
    this.sort.direction = 'asc';
    this.setDataListing();
    this.dataListFilterChangedSubscription = this.jobService.dataListSubject.subscribe(value => {
      this.userFilterRequest = value;
      this.refreshTable.emit();
    });
  }
  createControls(): void {
    this.amountform = new FormGroup({
      amountControl: new FormControl('', [Validators.required]),
    });
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
          this.dataSource = payloadResponse.data.records;
          if (this.filteredrecords === 0) {
            this.appNotificationService.info(this.translatePipe.transform('INFO_NO_RECORDS_FOUND'));
          }
        } else {
          this.dataSource = [];
        }
      }
    );
  }

  getDataList(): Observable<PayloadResponse<DataListResponse<JobDataListResponse>>> {
    const dataListRequest = new DataListRequest<JobFiltrRequest>();
    dataListRequest.filter = this.userFilterRequest;
    dataListRequest.page = this.paginator.pageIndex + 1;
    dataListRequest.limit = this.paginator.pageSize;
    dataListRequest.orderby = this.sort.active;
    dataListRequest.orderdirection = this.sort.direction;
    return this.jobService.getDataList(dataListRequest);
  }
  setSelectedStatus(a: JobDataListResponse) {
    this.selected = a;
  }
  changestatus() {
    if (this.amountform.invalid) {
      return;
    } else {
      const statusChangeRequest = new StatusChangeRequest();
      statusChangeRequest.state = 6;
      statusChangeRequest.tokenid = this.selected.tokenid;
      statusChangeRequest.amount = this.amount;
      console.log('statusChangeRequest.amount', statusChangeRequest.amount);
      this.jobService.updateStatus(statusChangeRequest).subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.appNotificationService.success();
          console.log('payloadResponse.issuccess', payloadResponse.issuccess);
          this.selected.state = 6;
        }
      });
    }
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
  // viewDetails(id: number, tokenid: number) {
  //  this.router.navigate(['jobs', 'details'],
  // { queryParams: { id: id, tokenid: tokenid, callbackurl: 'jobs' }, skipLocationChange: true });
  // }
  viewDetails(row: JobDataListResponse) {
    this.router.navigate(['jobs', 'details'],
    { queryParams: { id: row.vehicleid, tokenid: row.tokenid, token: row.token, callbackurl: 'jobs' }, skipLocationChange: true });
  }

  ngOnDestroy(): void {
    if (this.dataListFilterChangedSubscription) {
      this.dataListFilterChangedSubscription.unsubscribe();
    }
  }
}
