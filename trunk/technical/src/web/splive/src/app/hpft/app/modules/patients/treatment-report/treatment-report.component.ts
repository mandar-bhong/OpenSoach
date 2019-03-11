import { Component, OnInit, ViewChild, EventEmitter } from '@angular/core';
import { trigger, state, style, transition, animate } from '@angular/animations';
import { MatPaginator, MatSort, MatTableDataSource } from '@angular/material';
import { TreatmentFilterRequest, TreatmentResponse } from 'app/models/api/patient-models';
import { Subscription, merge, Observable } from 'rxjs';
import { PatientService } from 'app/services/patient.service';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';
import { TranslatePipe } from '../../../../../shared/pipes/translate/translate.pipe';
import { startWith, switchMap, map } from 'rxjs/operators';
import { PayloadResponse } from '../../../../../shared/models/api/payload-models';
import { DataListResponse, DataListRequest } from '../../../../../shared/models/api/data-list-models';
import { TreatmentModel } from 'app/models/ui/patient-models';
import { FileDownloadRequest } from 'app/models/api/file-download-request';
import { AppLocalStorage } from '../../../../../shared/services/app-data-store/app-data-store';

@Component({
  selector: 'app-treatment-report',
  templateUrl: './treatment-report.component.html',
  styleUrls: ['./treatment-report.component.css'],
  animations: [
    trigger('detailExpand', [
      state('collapsed', style({ height: '0px', minHeight: '0', display: 'none' })),
      state('expanded', style({ height: '*' })),
      transition('expanded <=> collapsed', animate('225ms cubic-bezier(0.4, 0.0, 0.2, 1)')),
    ]),
  ],

})
export class TreatmentReportComponent implements OnInit {

  displayedColumns = ['treatmentdone', 'treatmentperformedtime', 'view'];
  sortByColumns = [
    { text: 'Treatment Done', value: 'treatmentdone' },
    { text: 'Performed On', value: 'treatmentperformedtime' }
  ];

  @ViewChild(MatPaginator)
  paginator: MatPaginator;
  @ViewChild(MatSort)
  sort: MatSort;
  isReportAdd: boolean;;
  refreshTable: EventEmitter<null> = new EventEmitter();
  dataSource;
  filteredrecords = 0;
  isLoadingResults: boolean;
  isViewSchedule = false;
  dataListFilterChangedSubscription: Subscription;
  dataModel = new TreatmentModel();
  admissionid: number;
  treatmentid: number;
  treatmentFilterRequest: TreatmentFilterRequest;
  expandedElement: TreatmentResponse | null;
  treatmentResponseArray: TreatmentResponse[] = [];
  dataListRequest: DataListRequest<TreatmentFilterRequest>;

  constructor(public patientService: PatientService,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe,
    private appLocalStorage: AppLocalStorage) {
    this.isReportAdd = false;
  }

  ngOnInit() {
    this.paginator.pageSize = 10;
    this.sort.direction = 'asc';
    this.paginator.pageIndex = 1;
    this.sort.active = 'admissionid';
    this.sort.direction = 'asc';
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
            this.treatmentResponseArray = [];
            this.treatmentResponseArray = payloadResponse.data.records;
            this.dataSource = new MatTableDataSource<TreatmentResponse>(this.treatmentResponseArray);
          } else {
            this.dataSource = [];
          }
        }
      );
  }

  getDataList(): Observable<PayloadResponse<DataListResponse<TreatmentResponse>>> {
    const dataListRequest = new DataListRequest<TreatmentFilterRequest>();
    dataListRequest.orderdirection = this.sort.direction;
    dataListRequest.limit = this.paginator.pageSize;
    dataListRequest.page = this.paginator.pageIndex;
    dataListRequest.orderby = this.sort.active;
    dataListRequest.page = this.paginator.pageIndex;
    dataListRequest.filter = new TreatmentFilterRequest();
    dataListRequest.filter.admissionid = this.patientService.admissionid;
    return this.patientService.getTreatmentList(dataListRequest);
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
  restFormData(value) {
    console.log('value', value);
    if (value == 1) {
      this.isReportAdd = !this.isReportAdd;
    } else {
      this.isReportAdd = !this.isReportAdd;
      this.setDataListing();
    }

  }
  addReport() {
    this.isReportAdd = !this.isReportAdd;
  }

}
