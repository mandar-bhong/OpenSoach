import { Component, OnInit, Input, EventEmitter, ViewChild } from '@angular/core';
import { CustomerService } from '../../../../services/customer.service';
import { CustomerFilterRequest, CustomerDataListingItemResponse } from '../../../../models/api/customer-models';
import { DataListRequest, DataListResponse } from '../../../../../../shared/models/api/data-list-models';
import { PayloadResponse } from '../../../../../../shared/models/api/payload-models';
import { MatPaginator, MatSort } from '@angular/material';
import { merge } from 'rxjs/observable/merge';
import { map } from 'rxjs/operators/map';
import { Observable } from 'rxjs/Observable';
import { startWith } from 'rxjs/operators/startWith';
import { switchMap } from 'rxjs/operators/switchMap';

export class TestColumn {
  col: string;
  displayname: string;
  breakpoint: string;
  css: string;
}
@Component({
  selector: 'app-customer-list-view',
  templateUrl: './customer-list-view.component.html',
  styleUrls: ['./customer-list-view.component.css']
})
export class CustomerListViewComponent implements OnInit {
  displayedColumns = ['custname', 'corpname', 'poc1name', 'poc1emailid', 'poc1mobileno'];
  @ViewChild(MatPaginator)
  paginator: MatPaginator;
  @ViewChild(MatSort)
  sort: MatSort;
  refreshTable: EventEmitter<null> = new EventEmitter();
  dataSource;
  filteredrecords = 0;
  isLoadingResults = true;
  customerFilterRequest: CustomerFilterRequest;
  constructor(private customerService: CustomerService) { }

  ngOnInit() {
    // set default load parameters
    this.paginator.pageSize = 10;
    this.sort.active = 'custname';
    this.sort.direction = 'asc';

    this.setDataListing();

    this.customerService.dataListSubject.subscribe(value => {
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
        // Flip flag to show that loading has finished.
        this.isLoadingResults = false;
        return data;
      }),
    ).subscribe(
      payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          console.log('Gazette  Response succeed');
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

}
