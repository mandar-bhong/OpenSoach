import { Component, EventEmitter, OnDestroy, OnInit, ViewChild } from '@angular/core';
import { MatPaginator, MatSort } from '@angular/material';
import { Router } from '@angular/router';
import { merge, Observable, Subscription } from 'rxjs';
import { map, startWith, switchMap } from 'rxjs/operators';

import { SpServiceConfService } from '../../../../../../prod-shared/services/spservice/sp-service-conf.service';
import { DataListRequest, DataListResponse } from '../../../../../../shared/models/api/data-list-models';
import { PayloadResponse } from '../../../../../../shared/models/api/payload-models';
import { TranslatePipe } from '../../../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../../../shared/services/notification/app-notification.service';
import {
  ServiceConfigureFilterRequest,
  ServiceConfigureDataListResponse
} from '../../../../../../prod-shared/models/api/service-configuration-models';
import { ServicepointDataListResponse } from '../../../../../../prod-shared/models/api/servicepoint-models';


@Component({
  selector: 'app-chart-view',
  templateUrl: './chart-view.component.html',
  styleUrls: ['./chart-view.component.css']
})
export class ChartViewComponent implements OnInit, OnDestroy {
  displayedColumns = ['conftypecode', 'servconfname', 'action'];
  sortByColumns = [{ text: 'Chart Name', value: 'conftypecode' },
  { text: 'Service Configure Name', value: 'servconfname' }
  ];
  @ViewChild(MatPaginator)
  paginator: MatPaginator;
  @ViewChild(MatSort)
  sort: MatSort;
  refreshTable: EventEmitter<null> = new EventEmitter();
  dataSource;
  filteredrecords = 0;
  isLoadingResults = true;
  serviceConfigureFilterRequest: ServiceConfigureFilterRequest;
  dataListFilterChangedSubscription: Subscription;
  constructor(public spServiceConfService: SpServiceConfService,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe) { }

  ngOnInit() {
    this.paginator.pageSize = 10;
    this.sort.active = 'conftypecode';
    this.sort.direction = 'asc';

    this.setDataListing();
    this.dataListFilterChangedSubscription = this.spServiceConfService.dataListSubject.subscribe(value => {
      this.serviceConfigureFilterRequest = value;
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
  getDataList(): Observable<PayloadResponse<DataListResponse<ServiceConfigureDataListResponse>>> {
    const dataListRequest = new DataListRequest<ServiceConfigureFilterRequest>();
    dataListRequest.filter = this.serviceConfigureFilterRequest;
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
  configure(row: ServicepointDataListResponse) {
    this.router.navigate(['charts', 'configure'], {
      queryParams: { id: row.servconfid, mode: 1, callbackurl: 'servicepoints' }, skipLocationChange: true
    });
  }
  ngOnDestroy(): void {
    if (this.dataListFilterChangedSubscription) {
      this.dataListFilterChangedSubscription.unsubscribe();
    }
  }
}
