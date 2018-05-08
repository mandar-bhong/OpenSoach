import { Component, EventEmitter, OnDestroy, OnInit, ViewChild } from '@angular/core';
import { MatPaginator, MatSort } from '@angular/material';
import { Router } from '@angular/router';
import { Observable } from 'rxjs/Observable';
import { merge } from 'rxjs/observable/merge';
import { map } from 'rxjs/operators/map';
import { startWith } from 'rxjs/operators/startWith';
import { switchMap } from 'rxjs/operators/switchMap';
import { Subscription } from 'rxjs/Subscription';

import { DataListRequest, DataListResponse } from '../../../../../../shared/models/api/data-list-models';
import { PayloadResponse } from '../../../../../../shared/models/api/payload-models';
import { CustomerDataListingItemResponse, CustomerFilterRequest } from '../../../../models/api/customer-models';
import { CustomerService } from '../../../../services/customer.service';

@Component({
  selector: 'app-customer-list-view',
  templateUrl: './customer-list-view.component.html',
  styleUrls: ['./customer-list-view.component.css']
})

export class CustomerListViewComponent implements OnInit, OnDestroy {
  displayedColumns = ['custname', 'corpname', 'poc1name', 'poc1emailid', 'poc1mobileno', 'custstate', 'action'];
  sortByColumns = [{ text: 'Customer Name', value: 'custname' },
  { text: 'Corporate Name', value: 'corpname' },
  { text: 'Point of Contact', value: 'poc1name' },
  { text: 'Email Addres', value: 'poc1emailid' },
  { text: 'Mobile no', value: 'poc1mobileno' },
  { text: 'Status', value: 'custstate' }
  ];
  @ViewChild(MatPaginator)
  paginator: MatPaginator;
  @ViewChild(MatSort)
  sort: MatSort;
  refreshTable: EventEmitter<null> = new EventEmitter();
  dataSource;
  filteredrecords = 0;
  isLoadingResults = true;
  customerFilterRequest: CustomerFilterRequest;
  dataListFilterChangedSubscription: Subscription;

  constructor(public customerService: CustomerService,
    private router: Router) { }

  ngOnInit() {
    // set default load parameters
    this.paginator.pageSize = 10;
    this.sort.active = 'custname';
    this.sort.direction = 'asc';

    this.setDataListing();

    this.dataListFilterChangedSubscription = this.customerService.dataListSubject.subscribe(value => {
      this.customerFilterRequest = value;
      this.refreshTable.emit();
    });

  }

  setDataListing(): void {
    this.sort.sortChange.subscribe(() => this.paginator.pageIndex = 0);
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
        } else {
          this.dataSource = [];
        }
      }
    );
  }

  getDataList(): Observable<PayloadResponse<DataListResponse<CustomerDataListingItemResponse>>> {
    const dataListRequest = new DataListRequest<CustomerFilterRequest>();
    dataListRequest.filter = this.customerFilterRequest;
    dataListRequest.page = this.paginator.pageIndex + 1;
    dataListRequest.limit = this.paginator.pageSize;
    dataListRequest.orderby = this.sort.active;
    dataListRequest.orderdirection = this.sort.direction;
    return this.customerService.getDataList(dataListRequest);
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

  editRecord(id: number) {
    this.router.navigate(['customers', 'update'], { queryParams: { id: id, callbackurl: 'customers' }, skipLocationChange: true });
  }

  associateProduct(id: number) {
    this.router.navigate(['customers', 'products'],
      { queryParams: { id: id, callbackurl: 'customers' }, skipLocationChange: true });
  }

  ngOnDestroy(): void {
    if (this.dataListFilterChangedSubscription) {
      this.dataListFilterChangedSubscription.unsubscribe();
    }
  }
}
