import { Component, OnInit, OnDestroy, ViewChild, EventEmitter } from '@angular/core';
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
import { CorporateDataListingItemResponse, CorporateFilterRequest } from '../../../../models/api/corporate-models';
import { CorporateService } from '../../../../services/corporate.service';

@Component({
  selector: 'app-corporate-list-view',
  templateUrl: './corporate-list-view.component.html',
  styleUrls: ['./corporate-list-view.component.css']
})
export class CorporateListViewComponent implements OnInit, OnDestroy {
  displayedColumns = ['corpname', 'corpmobileno', 'corpemailid', 'corplandlineno', 'action'];
  sortByColumns = [{ text: 'Corporate Name', value: 'corpname' },
  { text: 'Mobile No.', value: 'corpmobileno' },
  { text: 'Email', value: 'corpemailid' },
  { text: 'Landline No.', value: 'corplandlineno' }
  ];
  @ViewChild(MatPaginator)
  paginator: MatPaginator;
  @ViewChild(MatSort)
  sort: MatSort;
  refreshTable: EventEmitter<null> = new EventEmitter();
  dataSource;
  filteredrecords = 0;
  isLoadingResults = true;
  corporateFilterRequest: CorporateFilterRequest;
  dataListFilterChangedSubscription: Subscription;
  constructor(public corporateService: CorporateService,
    private router: Router) { }

  ngOnInit() {
    this.paginator.pageSize = 10;
    this.sort.active = 'corpname';
    this.sort.direction = 'asc';
    this.setDataListing();
    this.dataListFilterChangedSubscription = this.corporateService.dataListSubject.subscribe(value => {
      this.corporateFilterRequest = value;
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
  getDataList(): Observable<PayloadResponse<DataListResponse<CorporateDataListingItemResponse>>> {
    const dataListRequest = new DataListRequest<CorporateFilterRequest>();
    dataListRequest.filter = this.corporateFilterRequest;
    dataListRequest.page = this.paginator.pageIndex + 1;
    dataListRequest.limit = this.paginator.pageSize;
    dataListRequest.orderby = this.sort.active;
    dataListRequest.orderdirection = this.sort.direction;
    return this.corporateService.getDataList(dataListRequest);
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
    this.router.navigate(['corporates', 'add'], { queryParams: { id: id, callbackurl: 'corporates' }, skipLocationChange: true });
  }

  ngOnDestroy(): void {
    if (this.dataListFilterChangedSubscription) {
      this.dataListFilterChangedSubscription.unsubscribe();
    }
  }
}
