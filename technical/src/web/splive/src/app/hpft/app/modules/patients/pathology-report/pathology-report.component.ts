import { animate, state, style, transition, trigger } from '@angular/animations';
import { Component, EventEmitter, OnInit, ViewChild } from '@angular/core';
import { MatPaginator, MatSort, MatTableDataSource } from '@angular/material';
import { Router } from '@angular/router';
import { FileDownloadRequest } from 'app/models/api/file-download-request';
import { PathologyFilterRequest, PathologyResponse } from 'app/models/api/patient-data-models';
import { PatientService } from 'app/services/patient.service';
import { merge, Observable, Subscription } from 'rxjs';
import { map, startWith, switchMap } from 'rxjs/operators';

import { DataListRequest, DataListResponse } from '../../../../../shared/models/api/data-list-models';
import { PayloadResponse } from '../../../../../shared/models/api/payload-models';
import { TranslatePipe } from '../../../../../shared/pipes/translate/translate.pipe';
import { AppLocalStorage } from '../../../../../shared/services/app-data-store/app-data-store';
import { FloatingButtonMenuService } from '../../../../../shared/services/floating-button-menu.service';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';


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

  displayedColumns = ['testperformed', 'testperformedtime', 'view'];
  sortByColumns = [
    { text: 'Test Performed', value: 'testperformed' },
    { text: 'Performed On', value: 'testperformedtime' }
  ];
  @ViewChild(MatPaginator)
  paginator: MatPaginator;
  @ViewChild(MatSort)
  sort: MatSort;
  refreshTable: EventEmitter<null> = new EventEmitter();
  dataSource;
  filteredrecords = 0;
  isReportAdd: boolean;
  isLoadingResults: boolean;
  isViewSchedule = false;
  pathologyFilterRequest: PathologyFilterRequest;
  dataListFilterChangedSubscription: Subscription;
  admissionid: number;
  expandedElement: PathologyResponse | null;
  pathologyResponseArray: PathologyResponse[] = [];
  dataListRequest: DataListRequest<PathologyFilterRequest>;

  constructor(public patientService: PatientService,
    private floatingButtonMenuService: FloatingButtonMenuService,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private appLocalStorage: AppLocalStorage,
    private translatePipe: TranslatePipe) {
    this.isReportAdd = false
  }

  ngOnInit() {
    this.paginator.pageSize = 10;
    this.sort.direction = 'asc';
    this.sort.active = 'testperformed';
    this.sort.active = 'admissionid';
    this.setDataListing();
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
            this.pathologyResponseArray = [];
            this.pathologyResponseArray = payloadResponse.data.records;
            this.dataSource = new MatTableDataSource<PathologyResponse>(this.pathologyResponseArray);
          } else {
            this.dataSource = [];
          }
        }
      );
  }

  getDataList(): Observable<PayloadResponse<DataListResponse<PathologyResponse>>> {
    const dataListRequest = new DataListRequest<PathologyFilterRequest>();
    dataListRequest.orderdirection = this.sort.direction;
    dataListRequest.limit = this.paginator.pageSize;
    dataListRequest.page = this.paginator.pageIndex + 1;
    dataListRequest.orderby = this.sort.active;
    dataListRequest.filter = new PathologyFilterRequest();
    dataListRequest.filter.admissionid = this.patientService.admissionid;
    return this.patientService.getPathologyList(dataListRequest);

  }


  // code block checking obejct is empty.
  checkEmptyObjects(object): boolean {
    if (Object.keys(object).length > 0) {
      return true;
    } else {
      return false;
    }
  }

  // code block for view schedule detsils  of particular action 
  viewSchedule(element) {
    this.isViewSchedule = true;
  }

  setOpenCloseSchedule() {
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


  addReport() {
    this.isReportAdd = !this.isReportAdd;
  }
  restFormData(value) {
    if (value == 1) {
      this.isReportAdd = !this.isReportAdd;
    } else {
      this.isReportAdd = !this.isReportAdd;
      this.setDataListing();
    }

  }

  downloadFille(id, filename) {
    const fileDownloadRequest = new FileDownloadRequest();
    fileDownloadRequest.token = this.appLocalStorage.getObject('AUTH_TOKEN');
    fileDownloadRequest.uuid = id;
    this.patientService.downloadFile(fileDownloadRequest).subscribe((filePayloadResponse) => {
      if (filePayloadResponse) {
        this.patientService.saveFile(filePayloadResponse, filename);
      }
    });
  }


}
