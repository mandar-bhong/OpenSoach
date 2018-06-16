import { Component, EventEmitter, OnDestroy, OnInit, ViewChild } from '@angular/core';
import { MatBottomSheet, MatPaginator, MatSort } from '@angular/material';
import { Router } from '@angular/router';
import { merge, Observable, Subscription } from 'rxjs';
import { map, startWith, switchMap } from 'rxjs/operators';

import { DEVICE_STATE } from '../../../../../shared/app-common-constants';
import { DataListRequest, DataListResponse } from '../../../../../shared/models/api/data-list-models';
import { PayloadResponse } from '../../../../../shared/models/api/payload-models';
import { TranslatePipe } from '../../../../../shared/pipes/translate/translate.pipe';
import { AppDeviceService } from '../../../../../shared/services/device/app-device.service';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';
import { DeviceDataListResponse, DeviceFilterRequest } from '../../../../models/api/device-models';
import { ProdDeviceService } from '../../../../services/device/prod-device.service';
import { DeviceDetailsViewComponent } from '../../device-details-view/device-details-view.component';

@Component({
  selector: 'app-device-list-view',
  templateUrl: './device-list-view.component.html',
  styleUrls: ['./device-list-view.component.scss']
})

export class DeviceListViewComponent implements OnInit, OnDestroy {
  displayedColumns = ['serialno', 'devname', 'connectionstate', 'batterylevel', 'action'];
  sortByColumns = [
    { text: 'Serial Number', value: 'serialno' },
    { text: 'Device Name', value: 'devname' },
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
  constructor(private deviceService: ProdDeviceService,
    private appDeviceService: AppDeviceService,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe,
    private bottomSheet: MatBottomSheet) { }

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

  viewDetails(row: DeviceDataListResponse): void {
    const bottomSheetRef = this.bottomSheet.open(DeviceDetailsViewComponent, { data: row.devid });
    bottomSheetRef.afterDismissed().subscribe(result => {
      if (result) {
        console.log('after dismiss', result);
        row.devid = Number(result.devid);
        row.devname = String(result.devname);
      }
    });
  }
  ngOnDestroy(): void {
    if (this.dataListFilterChangedSubscription) {
      this.dataListFilterChangedSubscription.unsubscribe();
    }
  }
}
