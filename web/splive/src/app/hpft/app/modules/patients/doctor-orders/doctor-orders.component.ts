import { animate, state, style, transition, trigger } from '@angular/animations';
import { Component, EventEmitter, OnInit, ViewChild } from '@angular/core';
import { MatPaginator, MatSort, MatTableDataSource } from '@angular/material';
import { CHECK_STATE } from 'app/app-constants';
import { DoctorOrderResponse } from 'app/models/api/doctor-order-response';
import { DoctorOrderRequest } from 'app/models/api/doctor-orders-request';
import { FileDownloadRequest } from 'app/models/api/file-download-request';
import { TransactionDetailsFilter } from 'app/models/api/transaction-details';
import { PatientService } from 'app/services/patient.service';
import { merge, Observable } from 'rxjs';
import { map, startWith, switchMap } from 'rxjs/operators';
import { DataListRequest, DataListResponse } from '../../../../../shared/models/api/data-list-models';
import { PayloadResponse } from '../../../../../shared/models/api/payload-models';
import { TranslatePipe } from '../../../../../shared/pipes/translate/translate.pipe';
import { AppLocalStorage } from '../../../../../shared/services/app-data-store/app-data-store';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';

@Component({
  selector: 'app-doctor-orders',
  templateUrl: './doctor-orders.component.html',
  styleUrls: ['./doctor-orders.component.css', '../transaction-details/transaction-details.component.css'],
  animations: [
    trigger('detailExpand', [
      state('collapsed', style({ height: '0px', minHeight: '0', display: 'none' })),
      state('expanded', style({ height: '*' })),
      transition('expanded <=> collapsed', animate('225ms cubic-bezier(0.4, 0.0, 0.2, 1)')),
    ]),
  ],
})
export class DoctorOrdersComponent implements OnInit {

  @ViewChild(MatPaginator)
  paginator: MatPaginator;
  @ViewChild(MatSort)
  sort: MatSort;
  refreshTable: EventEmitter<null> = new EventEmitter();
  dataSource;
  filteredrecords = 0;
  isLoadingResults: boolean;
  doctorOrderRequest: DoctorOrderRequest;
  expandedElement: DoctorOrderResponse | null;
  doctorOrderResponseArray: DoctorOrderResponse[] = [];
  dataListRequest: DataListRequest<DoctorOrderRequest>;
  isViewSchedule = false;
  CHECK_STATE:CHECK_STATE;
  constructor(
    private appNotificationService: AppNotificationService,
    private appLocalStorage: AppLocalStorage,
    private translatePipe: TranslatePipe,
    private patientService: PatientService) { }
  displayedColumns = ['ordertype', 'ordercreatedtime', 'by', 'view'];
  sortByColumns = [
    { text: 'Type', value: 'ordertype' },
    { text: 'On', value: 'ordercreatedtime' }
  ];
  // columnsToDisplay = ['fname', 'date'];
  ngOnInit() {
    this.paginator.pageSize = 10;
    this.sort.direction = 'asc';
    this.sort.active = 'admissionid';
    this.getDataListing();
  }
  getDataListing(): void {
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
            this.doctorOrderResponseArray = [];
            this.doctorOrderResponseArray = payloadResponse.data.records;
            this.dataSource = new MatTableDataSource<DoctorOrderResponse>(this.doctorOrderResponseArray);
          } else {
            this.dataSource = [];
          }
        }
      );
  }

  getDataList(): Observable<PayloadResponse<DataListResponse<DoctorOrderResponse>>> {
    const dataListRequest = new DataListRequest<DoctorOrderRequest>();
    dataListRequest.orderdirection = this.sort.direction;
    dataListRequest.limit = this.paginator.pageSize;
    dataListRequest.page = this.paginator.pageIndex + 1;
    dataListRequest.orderby = this.sort.active;
    dataListRequest.filter = new TransactionDetailsFilter();
    dataListRequest.filter.admissionid = this.patientService.admissionid;
    return this.patientService.getDoctorOrderDetails(dataListRequest);
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

   // code block for check status
   checkStatus(status: number) {
    if (status == CHECK_STATE.ACKNOWLEDGED) {
        return 'Acknowledged';
      } else {
        return 'New';
      }
  }// end of code block
}
