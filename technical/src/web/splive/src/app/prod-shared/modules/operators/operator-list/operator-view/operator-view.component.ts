import { Component, EventEmitter, OnDestroy, OnInit, ViewChild } from '@angular/core';
import { MatPaginator, MatSort } from '@angular/material';
import { Router } from '@angular/router';
import { Observable } from 'rxjs/Observable';
import { merge } from 'rxjs/observable/merge';
import { map } from 'rxjs/operators/map';
import { startWith } from 'rxjs/operators/startWith';
import { switchMap } from 'rxjs/operators/switchMap';
import { Subscription } from 'rxjs/Subscription';

import { OPERATOR_AREA } from '../../../../../shared/app-common-constants';
import { DataListRequest, DataListResponse } from '../../../../../shared/models/api/data-list-models';
import { PayloadResponse } from '../../../../../shared/models/api/payload-models';
import { TranslatePipe } from '../../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';
import { OperatorDataListResponse, OperatorFiltrRequest } from '../../../../models/api/operator-models';
import { ProdOperatorService } from '../../../../services/operator/prod-operator.service';

@Component({
  selector: 'app-operator-view',
  templateUrl: './operator-view.component.html',
  styleUrls: ['./operator-view.component.css']
})
export class OperatorViewComponent implements OnInit, OnDestroy {
  displayedColumns = ['fopname', 'emailid', 'mobileno', 'fopstate', 'foparea', 'action'];
  sortByColumns = [{ text: 'User Name', value: 'usrname' },
  { text: 'Operator Name', value: 'fopname' },
  { text: 'Email id', value: 'emailid' },
  { text: 'Mobile Number', value: 'mobileno' },
  { text: 'Status', value: 'fopstate' },
  { text: 'Area', value: 'foparea' }
  ];
  @ViewChild(MatPaginator)
  paginator: MatPaginator;
  @ViewChild(MatSort)
  sort: MatSort;
  refreshTable: EventEmitter<null> = new EventEmitter();
  dataSource;
  filteredrecords = 0;
  isLoadingResults = true;
  operatorFiltrRequest: OperatorFiltrRequest;
  dataListFilterChangedSubscription: Subscription;
  showEditForm = false;
  operatorAreas = OPERATOR_AREA;
  constructor(public prodOperatorService: ProdOperatorService,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe) { }

  ngOnInit() {
    this.paginator.pageSize = 10;
    this.sort.active = 'fopname';
    this.sort.direction = 'asc';

    this.setDataListing();
    this.dataListFilterChangedSubscription = this.prodOperatorService.dataListSubject.subscribe(value => {
      this.operatorFiltrRequest = value;
      this.refreshTable.emit();
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
  getDataList(): Observable<PayloadResponse<DataListResponse<OperatorDataListResponse>>> {
    const dataListRequest = new DataListRequest<OperatorFiltrRequest>();
    dataListRequest.filter = this.operatorFiltrRequest;
    dataListRequest.page = this.paginator.pageIndex + 1;
    dataListRequest.limit = this.paginator.pageSize;
    dataListRequest.orderby = this.sort.active;
    dataListRequest.orderdirection = this.sort.direction;
    return this.prodOperatorService.getDataList(dataListRequest);
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

  viewDetails(id: number) {
    this.router.navigate(['foperators', 'detail'], { queryParams: { id: id, callbackurl: 'foperators' }, skipLocationChange: true });
  }
  viewAssocate(id: number) {
    this.router.navigate(['foperators', 'associate'], { queryParams: { id: id, callbackurl: 'foperators' }, skipLocationChange: true });
  }
  ngOnDestroy(): void {
    if (this.dataListFilterChangedSubscription) {
      this.dataListFilterChangedSubscription.unsubscribe();
    }
  }
}
