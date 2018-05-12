// import * as console from 'console';
import { Component, OnInit, EventEmitter, OnDestroy, ViewChild } from '@angular/core';
import { MatPaginator, MatSort } from '@angular/material';
import { Observable } from 'rxjs/Observable';
import { merge } from 'rxjs/observable/merge';
import { map } from 'rxjs/operators/map';
import { startWith } from 'rxjs/operators/startWith';
import { switchMap } from 'rxjs/operators/switchMap';
import { Subscription } from 'rxjs/Subscription';
import { Router } from '@angular/router';

import { DataListRequest, DataListResponse } from '../../../../../../shared/models/api/data-list-models';
import { PayloadResponse } from '../../../../../../shared/models/api/payload-models';
import { UserFilterRequest } from '../../../../../app/models/api/user-models';
import { UserService } from '../../../../services/user.service';
import { UserDataListResponse } from '../../../../../../shared/models/api/user-models';
@Component({
  selector: 'app-user-view',
  templateUrl: './user-view.component.html',
  styleUrls: ['./user-view.component.css']
})
export class UserViewComponent implements OnInit, OnDestroy {
  displayedColumns = ['usrname', 'usrcategory', 'urolename', 'usrstate', 'fname', 'lname', 'mobileno', 'action'];
  sortByColumns = [{ text: 'User Name', value: 'usrname' },
  { text: 'User Category', value: 'corpname' },
  { text: 'User State', value: 'usrcategory' },
  { text: 'User Role Name', value: 'urolename' },
  { text: 'User State Since', value: 'usrstate' }
  ];
  @ViewChild(MatPaginator)
  paginator: MatPaginator;
  @ViewChild(MatSort)
  sort: MatSort;
  refreshTable: EventEmitter<null> = new EventEmitter();
  dataSource;
  filteredrecords = 0;
  isLoadingResults = true;
  userFilterRequest: UserFilterRequest;
  dataListFilterChangedSubscription: Subscription;

  constructor(public userService: UserService,
    private router: Router) { }
  ngOnInit() {
    this.paginator.pageSize = 10;
    this.sort.active = 'usrname';
    this.sort.direction = 'asc';

    this.setDataListing();

    this.dataListFilterChangedSubscription = this.userService.dataListSubject.subscribe(value => {
      this.userFilterRequest = value;
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

  getDataList(): Observable<PayloadResponse<DataListResponse<UserDataListResponse>>> {
    const dataListRequest = new DataListRequest<UserFilterRequest>();
    dataListRequest.filter = this.userFilterRequest;
    dataListRequest.page = this.paginator.pageIndex + 1;
    dataListRequest.limit = this.paginator.pageSize;
    dataListRequest.orderby = this.sort.active;
    dataListRequest.orderdirection = this.sort.direction;
    return this.userService.getDataList(dataListRequest);
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
    this.router.navigate(['users', 'user-detail'], { queryParams: { id: id, callbackurl: 'users' }, skipLocationChange: true });
  }

  associateProduct(id: number) {
    this.router.navigate(['users', 'products'],
      { queryParams: { id: id, callbackurl: 'users' }, skipLocationChange: true });
  }

  ngOnDestroy(): void {
    if (this.dataListFilterChangedSubscription) {
      this.dataListFilterChangedSubscription.unsubscribe();
    }
  }
}
