import { Component, EventEmitter, OnDestroy, OnInit, ViewChild } from '@angular/core';
import { MatPaginator, MatSort } from '@angular/material';
import { Router } from '@angular/router';
import { Observable } from 'rxjs/Observable';
import { merge } from 'rxjs/observable/merge';
import { map } from 'rxjs/operators/map';
import { startWith } from 'rxjs/operators/startWith';
import { switchMap } from 'rxjs/operators/switchMap';
import { Subscription } from 'rxjs/Subscription';

import { ComplaintDataListResponse, ComplaintFiltrRequest } from '../../../../models/api/complaint-models';
import { ProdComplaintService } from '../../../../services/complaint/prod-complaint.service';
import { AppNotificationService } from '../../../../../../shared/services/notification/app-notification.service';
import { TranslatePipe } from '../../../../../../shared/pipes/translate/translate.pipe';
import { DataListResponse, DataListRequest } from '../../../../../../shared/models/api/data-list-models';
import { PayloadResponse } from '../../../../../../shared/models/api/payload-models';

@Component({
  selector: 'app-complaint-view',
  templateUrl: './complaint-view.component.html',
  styleUrls: ['./complaint-view.component.css']
})
export class ComplaintViewComponent implements OnInit, OnDestroy {
  displayedColumns = ['spname', 'complainttitle', 'description', 'complaintby', 'complaintstate', 'action'];
  sortByColumns = [{ text: 'Service Point Name', value: 'spname' },
  { text: 'Compalint Title', value: 'complainttitle' },
  { text: 'complaintby', value: 'Complaint By' }
  ];
  @ViewChild(MatPaginator)
  paginator: MatPaginator;
  @ViewChild(MatSort)
  sort: MatSort;
  refreshTable: EventEmitter<null> = new EventEmitter();
  dataSource;
  filteredrecords = 0;
  isLoadingResults = true;
  complaintFiltrRequest: ComplaintFiltrRequest;
  dataListFilterChangedSubscription: Subscription;
  showEditForm = false;
  constructor(public prodComplaintService: ProdComplaintService,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe) { }

  ngOnInit() {
    this.paginator.pageSize = 10;
    this.sort.active = 'spname';
    this.sort.direction = 'asc';

    this.setDataListing();
    this.dataListFilterChangedSubscription = this.prodComplaintService.dataListSubject.subscribe(value => {
      this.complaintFiltrRequest = value;
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
  getDataList(): Observable<PayloadResponse<DataListResponse<ComplaintDataListResponse>>> {
    const dataListRequest = new DataListRequest<ComplaintFiltrRequest>();
    dataListRequest.filter = this.complaintFiltrRequest;
    dataListRequest.page = this.paginator.pageIndex + 1;
    dataListRequest.limit = this.paginator.pageSize;
    dataListRequest.orderby = this.sort.active;
    dataListRequest.orderdirection = this.sort.direction;
    return this.prodComplaintService.getDataList(dataListRequest);
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
    this.router.navigate(['complaints', 'detail'], { queryParams: { id: id, callbackurl: 'complaints' }, skipLocationChange: true });
  }

  ngOnDestroy(): void {
    if (this.dataListFilterChangedSubscription) {
      this.dataListFilterChangedSubscription.unsubscribe();
    }
  }
}
