import { Component, EventEmitter, OnDestroy, OnInit, ViewChild, ChangeDetectorRef } from '@angular/core';
import { MatPaginator, MatSort, MatTableDataSource } from '@angular/material';
import { Router } from '@angular/router';
import { Observable, merge, Subscription } from 'rxjs';
import { map, startWith, switchMap } from 'rxjs/operators';
import {
 PatientDetaFilterRequest,
  PatientDataListResponse, PatientDetailAddRequest
} from '../../../../models/api/patient-models';
import { PayloadResponse } from '../../../../../../shared/models/api/payload-models';
import { DataListRequest, DataListResponse } from '../../../../../../shared/models/api/data-list-models';
import { AppNotificationService } from '../../../../../../shared/services/notification/app-notification.service';
import { TranslatePipe } from '../../../../../../shared/pipes/translate/translate.pipe';
import { PatientService } from '../../../../services/patient.service';
import { element } from 'protractor';
import { PatientDataModel } from '../../../../models/ui/patient-models';


@Component({
  selector: 'app-patient-view',
  templateUrl: './patient-view.component.html',
  styleUrls: ['./patient-view.component.css']
})
export class PatientViewComponent implements OnInit, OnDestroy {
  displayedColumns = ['patientname', 'emergencycontactno', 'ward', 'bedno', 'statusid'];
  sortByColumns = [{ text: 'Patient Name', value: 'patientname' },
  { text: 'Emergency Contact Number', value: 'emergencycontactno' },
  { text: 'Ward', value: 'ward' },
  { text: 'Bed/Room Number', value: 'bedno' }
  ];
  @ViewChild(MatPaginator)
  paginator: MatPaginator;
  @ViewChild(MatSort)
  sort: MatSort;
  refreshTable: EventEmitter<null> = new EventEmitter();
  // dataSource = [];
  patients = new PatientDataModel();
  // patients;
  patient = [];
  dataSource;
  stat;
  filteredrecords = 0;
  isLoadingResults = true;
  patientDetaFilterRequest: PatientDetaFilterRequest;
  dataListFilterChangedSubscription: Subscription;
  showEditForm = false;
  constructor(public patientService: PatientService,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe,
    private ch: ChangeDetectorRef) { }

  ngOnInit() {
    this.paginator.pageSize = 10;
    this.sort.active = 'patientname';
    this.sort.direction = 'asc';
    this.getDataList();

    this.patients = new PatientDataModel();
    // this.setDataListing();
    // this.dataListFilterChangedSubscription = this.patientService.dataListSubject.subscribe(value => {
    //   this.patientDetaFilterRequest = value;
    //   this.refreshTable.emit();
    // });
  }
  // setDataListing(): void {
  //   this.sort.sortChange.subscribe(() => this.paginator.pageIndex = 0);
  //   this.refreshTable.subscribe(() => this.paginator.pageIndex = 0);
  //   merge(this.sort.sortChange, this.paginator.page, this.refreshTable)
  //     .pipe(
  //       startWith({}),
  //       switchMap(() => {
  //         this.isLoadingResults = true;
  //         return this.getDataList();
  //       }),
  //       map(data => {
  //         this.isLoadingResults = false;
  //         return data;
  //       }),
  //   ).subscribe(
  //     payloadResponse => {
  //       if (payloadResponse && payloadResponse.issuccess) {
  //         this.filteredrecords = payloadResponse.data.filteredrecords;
  //         this.dataSource = payloadResponse.data.records;
  //         if (this.filteredrecords === 0) {
  //           this.appNotificationService.info(this.translatePipe.transform('INFO_NO_RECORDS_FOUND'));
  //         }
  //       } else {
  //         this.dataSource = [];
  //       }
  //     }
  //   );
  // }
  // getDataList(): Observable<PayloadResponse<PatientDataListResponse>> {
  //   const dataListRequest = new DataListRequest<PatientDetaFilterRequest>();
  //   dataListRequest.filter = this.patientDetaFilterRequest;
  //   dataListRequest.page = this.paginator.pageIndex + 1;
  //   dataListRequest.limit = this.paginator.pageSize;
  //   dataListRequest.orderby = this.sort.active;
  //   dataListRequest.orderdirection = this.sort.direction;
  //   return this.patientService.getDataList();
  // }
  getDataList() {
    this.patientService.getDataList().subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        // console.log('payloadResponse.data', payloadResponse.data);
        // console.log(JSON.parse(payloadResponse.data[0].patientdetails));

        payloadResponse.data.forEach(value => {
          // this.patient.push(JSON.parse(value.patientdetails));
          const a = new PatientDataModel();
          a.copyFrom(value);
          this.patient.push(a);
          console.log('this.dataSource', this.dataSource);
          this.dataSource = this.patient;
        });
      }
    });
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

  ngOnDestroy(): void {
    if (this.dataListFilterChangedSubscription) {
      this.dataListFilterChangedSubscription.unsubscribe();
    }
  }

}
