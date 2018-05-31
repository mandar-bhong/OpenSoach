import { Component, EventEmitter, OnDestroy, OnInit, ViewChild } from '@angular/core';
import { MatPaginator, MatSort } from '@angular/material';
import { Router } from '@angular/router';
import { Observable } from 'rxjs/Observable';
import { merge } from 'rxjs/observable/merge';
import { map } from 'rxjs/operators/map';
import { startWith } from 'rxjs/operators/startWith';
import { switchMap } from 'rxjs/operators/switchMap';
import { Subscription } from 'rxjs/Subscription';

import { DataListRequest, DataListResponse } from '../../../../../shared/models/api/data-list-models';
import { PayloadResponse } from '../../../../../shared/models/api/payload-models';
import { TranslatePipe } from '../../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';
import {
  ServicepointDataListResponse,
  SrevicepointFilterRequest,
} from '../../../../models/api/service-configuration-models';
import { SpServiceConfService } from '../../../../services/spservice/sp-service-conf.service';

@Component({
  selector: 'app-servicepoint-list-view',
  templateUrl: './servicepoint-list-view.component.html',
  styleUrls: ['./servicepoint-list-view.component.css']
})
export class ServicepointListViewComponent implements OnInit, OnDestroy {
  displayedColumns = ['spname', 'spcname', 'devid', 'spstate', 'servconfid', 'action'];
  sortByColumns = [{ text: 'Service Point', value: 'spname' },
  { text: 'Category', value: 'spcname' },
  { text: 'Device', value: 'devname' },
  { text: 'status', value: 'spstate' }
  ];
  @ViewChild(MatPaginator)
  paginator: MatPaginator;
  @ViewChild(MatSort)
  sort: MatSort;
  refreshTable: EventEmitter<null> = new EventEmitter();
  dataSource;
  filteredrecords = 0;
  isLoadingResults = true;
  srevicepointFilterRequest: SrevicepointFilterRequest;
  dataListFilterChangedSubscription: Subscription;
  servicepointDataListResponse: ServicepointDataListResponse;
  showEditForm = false;
  constructor(public spServiceConfService: SpServiceConfService,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe) { }

  ngOnInit() {
    this.paginator.pageSize = 10;
    this.sort.active = 'spname';
    this.sort.direction = 'asc';
    this.setDataListing();
    this.dataListFilterChangedSubscription = this.spServiceConfService.dataListSubject.subscribe(value => {
      this.srevicepointFilterRequest = value;
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
  getDataList(): Observable<PayloadResponse<DataListResponse<ServicepointDataListResponse>>> {
    const dataListRequest = new DataListRequest<SrevicepointFilterRequest>();
    dataListRequest.filter = this.srevicepointFilterRequest;
    dataListRequest.page = this.paginator.pageIndex + 1;
    dataListRequest.limit = this.paginator.pageSize;
    dataListRequest.orderby = this.sort.active;
    dataListRequest.orderdirection = this.sort.direction;
    return this.spServiceConfService.getDataList(dataListRequest);
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
    this.router.navigate(['servicepoints', 'details'], { queryParams: { id: id, callbackurl: 'servicepoints' }, skipLocationChange: true });
  }
  link() {

  }
  ngOnDestroy(): void {
    if (this.dataListFilterChangedSubscription) {
      this.dataListFilterChangedSubscription.unsubscribe();
    }
  }

}
