import { Component, EventEmitter, OnDestroy, OnInit, ViewChild } from '@angular/core';
import { MatPaginator, MatSort } from '@angular/material';
import { Router } from '@angular/router';
import { Observable } from 'rxjs/Observable';
import { merge } from 'rxjs/observable/merge';
import { map } from 'rxjs/operators/map';
import { startWith } from 'rxjs/operators/startWith';
import { switchMap } from 'rxjs/operators/switchMap';
import { Subscription } from 'rxjs/Subscription';

import { DEVICE_STATE } from '../../../../app-common-constants';
import { DataListRequest, DataListResponse } from '../../../../models/api/data-list-models';
import { DeviceDataListResponse, DeviceFilterRequest } from '../../../../models/api/device-models';
import { PayloadResponse } from '../../../../models/api/payload-models';
import { TranslatePipe } from '../../../../pipes/translate/translate.pipe';
import { DeviceSharedService } from '../../../../services/device-shared.service';
import { AppNotificationService } from '../../../../services/notification/app-notification.service';

@Component({
  selector: 'app-device-list-view',
  templateUrl: './device-list-view.component.html',
  styleUrls: ['./device-list-view.component.css']
})

export class DeviceListViewComponent implements OnInit, OnDestroy {
  displayedColumns = ['serialno', 'devname', 'devstate', 'connectionstate', 'batterylevel', 'action'];
  sortByColumns = [
    { text: 'Serial Number', value: 'serialno' },
    { text: 'Device Name', value: 'devname' },
    { text: 'State', value: 'devstate' },
    { text: 'Connection', value: 'connectionstate' },
    { text: 'Bettery Level', value: 'batterylevel' }
  ];
  @ViewChild(MatPaginator)
  paginator: MatPaginator;
  @ViewChild(MatSort)
  sort: MatSort;
  refreshTable: EventEmitter<null> = new EventEmitter();
  dataSource;
  filteredrecords = 0;
  isLoadingResults = true;
  deviceFilterRequest: DeviceFilterRequest;
  dataListFilterChangedSubscription: Subscription;
  deviceState = DEVICE_STATE;
  constructor(public deviceService: DeviceSharedService,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe) { }

  ngOnInit() {

    this.paginator.pageSize = 10;
    this.sort.active = 'serialno';
    this.sort.direction = 'asc';
    this.setDataListing();
    this.dataListFilterChangedSubscription = this.deviceService.dataListSubject.subscribe(value => {
      this.deviceFilterRequest = value;
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
  getDataList(): Observable<PayloadResponse<DataListResponse<DeviceDataListResponse>>> {
    const dataListRequest = new DataListRequest<DeviceFilterRequest>();
    dataListRequest.filter = this.deviceFilterRequest;
    dataListRequest.page = this.paginator.pageIndex + 1;
    dataListRequest.limit = this.paginator.pageSize;
    dataListRequest.orderby = this.sort.active;
    dataListRequest.orderdirection = this.sort.direction;
    return this.deviceService.getDataList(dataListRequest);
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
    this.router.navigate(['devices', 'details'], { queryParams: { id: id, callbackurl: 'devices' }, skipLocationChange: true });
  }

  ngOnDestroy(): void {
    if (this.dataListFilterChangedSubscription) {
      this.dataListFilterChangedSubscription.unsubscribe();
    }
  }
}
