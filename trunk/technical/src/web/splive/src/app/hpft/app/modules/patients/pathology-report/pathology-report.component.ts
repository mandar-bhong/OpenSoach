import { Component, OnInit, ViewChild, EventEmitter } from '@angular/core';
import { MatPaginator, MatSort, MatTableDataSource } from '@angular/material';
import { Subscription, Observable, merge } from 'rxjs';
import { trigger, state, style, transition, animate } from '@angular/animations';
import { PathologyResponse, PathologyFilterRequest, ActionPathologyDataValue } from 'app/models/api/patient-models';
import { PayloadResponse } from '../../../../../shared/models/api/payload-models';
import { DataListResponse, DataListRequest } from '../../../../../shared/models/api/data-list-models';
import { PatientService } from 'app/services/patient.service';
import { Router } from '@angular/router';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';
import { startWith, switchMap, map } from 'rxjs/operators';
import { TranslatePipe } from '../../../../../shared/pipes/translate/translate.pipe';
import { PathologyModel } from 'app/models/ui/patient-models';

@Component({
  selector: 'app-pathology-report',
  templateUrl: './pathology-report.component.html',
  styleUrls: ['./pathology-report.component.css'],
  animations: [
    trigger('detailExpand', [
      state('collapsed', style({ height: '0px', minHeight: '0', display: 'none' })),
      state('expanded', style({ height: '*' })),
      transition('expanded <=> collapsed', animate('225ms cubic-bezier(0.4, 0.0, 0.2, 1)')),
    ]),
  ],
})
export class PathologyReportComponent implements OnInit {

  displayedColumns = ['testperformed', 'txndate', 'view'];
  sortByColumns = [
    { text: 'Test Performed', value: 'testperformed' },
    { text: 'Performed On', value: 'txndate' }
  ];
  @ViewChild(MatPaginator)
  paginator: MatPaginator;
  @ViewChild(MatSort)
  sort: MatSort;
  refreshTable: EventEmitter<null> = new EventEmitter();
  dataSource;
  filteredrecords = 0;
  isLoadingResults: boolean;
  isViewSchedule = false;
  pathologyFilterRequest: PathologyFilterRequest;
  dataListFilterChangedSubscription: Subscription;
  dataModel = new PathologyModel();
  admissionid: number;
  pathoResponse: PathologyResponse<ActionPathologyDataValue>[] = []

  constructor(public patientService: PatientService,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe) { }

  ngOnInit() {
    this.paginator.pageSize = 10;
    this.sort.direction = 'asc';
    this.paginator.pageIndex = 1;
    this.sort.active = 'testperformed';
    // this.sort.active = this.admissionid;
    this.sort.direction = 'asc';
    // this.pathologyFilterRequest = new PathologyFilterRequest();
    // this.pathologyFilterRequest.admissionid = this.dataModel.admissionid;
    this.setDataListing();
    // this.dataListFilterChangedSubscription = this.patientService.dataListSubject.subscribe(value => {
    //   // this.pathologyFilterRequest = value;
    //   this.refreshTable.emit();
    // });
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
            payloadResponse.data.records.forEach((item: any) => {
              const pathologyResponse = new PathologyResponse<ActionPathologyDataValue>();
              Object.assign(pathologyResponse, item);
              this.pathoResponse.push(pathologyResponse);
            });

            this.dataSource = new MatTableDataSource<PathologyResponse<ActionPathologyDataValue>>(this.pathoResponse);

            if (this.filteredrecords === 0) {
              this.appNotificationService.info(this.translatePipe.transform('INFO_NO_RECORDS_FOUND'));
            }
          } else {
            this.dataSource = [];
          }
        }
      );
  }


  getDataList(): Observable<PayloadResponse<DataListResponse<PathologyResponse<string>[]>>> {
    const dataListRequest = new DataListRequest<PathologyFilterRequest>();
    dataListRequest.page = this.paginator.pageIndex;
    dataListRequest.limit = this.paginator.pageSize + 1;
    dataListRequest.orderby = this.sort.active;
    dataListRequest.filter = new PathologyFilterRequest();

    dataListRequest.filter.admissionid = this.patientService.admissionid;
    dataListRequest.orderdirection = this.sort.direction;
    return this.patientService.getPathologyList(dataListRequest);

  }


  // code block cehcking obejct is emppty.
  checkEmptyObjects(object): boolean {
    if (Object.keys(object).length > 0) {
      return true;
    } else {
      return false;
    }
  }// end of fucntion 

  // code bloxk for view schedule detsils  of particular action 
  viewSchedule(element) {
    console.log('view schedule clickd');
    this.isViewSchedule = true;
  }

  setOpenCloseSchedule() {
    console.log('view setOpenCloseSchedule clickd');
    this.isViewSchedule = false;
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

}
