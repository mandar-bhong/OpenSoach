import { Component, EventEmitter, OnDestroy, OnInit, ViewChild } from '@angular/core';
import { MatBottomSheet, MatPaginator, MatSort } from '@angular/material';
import { Router } from '@angular/router';
import { merge, Observable, Subscription } from 'rxjs';
import { map, startWith, switchMap } from 'rxjs/operators';

import { DataListRequest, DataListResponse } from '../../../../../shared/models/api/data-list-models';
import { PayloadResponse } from '../../../../../shared/models/api/payload-models';
import { TranslatePipe } from '../../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';
import { ServicepointDataListResponse, SrevicepointFilterRequest } from '../../../../models/api/servicepoint-models';
import { ProdServicepointService } from '../../../../services/servicepoint/prod-servicepoint.service';
import {
  ServicepointDeviceAssociateComponent,
} from '../../servicepoint-device-associate/servicepoint-device-associate.component';
import { ServicepointUpdateComponent } from '../../servicepoint-update/servicepoint-update.component';


@Component({
  selector: 'app-servicepoint-list-view',
  templateUrl: './servicepoint-list-view.component.html',
  styleUrls: ['./servicepoint-list-view.component.css']
})
export class ServicepointListViewComponent implements OnInit, OnDestroy {
  displayedColumns = ['spname', 'spcname', 'devid', 'servconfid', 'action'];
  sortByColumns = [{ text: 'Service Point', value: 'spname' },
  { text: 'Category', value: 'spcname' },
  { text: 'Device', value: 'devname' }
    // { text: 'Status', value: 'spstate' }
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
  constructor(public prodServicepointService: ProdServicepointService,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe,
    private bottomSheet: MatBottomSheet) { }

  ngOnInit() {
    this.paginator.pageSize = 10;
    this.sort.active = 'spname';
    this.sort.direction = 'asc';
    this.setDataListing();
    this.dataListFilterChangedSubscription = this.prodServicepointService.dataListSubject.subscribe(value => {
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
    return this.prodServicepointService.getDataList(dataListRequest);
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

    if (row.servconfid) {
      this.router.navigate(['charts', 'configure'], {
        queryParams: { id: row.servconfid, mode: 1, callbackurl: 'servicepoints' }, skipLocationChange: true
      });
    } else {
      this.router.navigate(['servicepoints', 'service-associate'],
        { queryParams: { id: row.spid, callbackurl: 'servicepoints' }, skipLocationChange: true });
    }
  }
  editServicePoint(row: ServicepointDataListResponse): void {
    const bottomSheetRef = this.bottomSheet.open(ServicepointUpdateComponent, { data: row.spid });
    bottomSheetRef.afterDismissed().subscribe(result => {
      if (result) {
        console.log('after dismiss check', result);
        row.spid = Number(result.spid);
        row.spname = String(result.spname);
        row.spstate = Number(result.spstate);
        row.spcid = Number(result.spcid);
        row.spcname = String(result.spcname);
        // row.spcid = Number(result.spcname);
      }
    });
  }

  showChartData(row: ServicepointDataListResponse) {
    this.router.navigate(['charts'], {
      queryParams: { servconfid: row.servconfid, spid: row.spid, callbackurl: 'servicepoints' }, skipLocationChange: true
    });
  }

  ngOnDestroy(): void {
    if (this.dataListFilterChangedSubscription) {
      this.dataListFilterChangedSubscription.unsubscribe();
    }
  }
  openServicePointDeviceAssociation(sp: ServicepointDataListResponse): void {
    const bottomSheetRef = this.bottomSheet.open(ServicepointDeviceAssociateComponent, { data: sp.spid });
    bottomSheetRef.afterDismissed().subscribe(result => {
      if (result) {
        console.log('after dismiss', result);
        sp.devid = Number(result.devid);
        sp.devname = String(result.devname);
      }
    });
  }
}
